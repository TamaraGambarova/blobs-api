package requests

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"gitlab.com/distributed_lab/figure"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	pageParamLimit = "page[limit]"
)

//base - base struct describing params provided by user for the request handler
type base struct {
	filter map[string]string

	queryValues *url.Values
	request     *http.Request
}

type baseOpts struct {
	supportedFilters map[string]struct{}
}

// newRequest - creates new instance of request
func newBase(r *http.Request, opts baseOpts) (*base, error) {
	request := base{
		request: r,
	}

	err := request.unmarshalQuery(opts)
	if err != nil {
		return nil, err
	}

	return &request, nil
}

func (r *base) unmarshalQuery(opts baseOpts) error {
	r.queryValues = new(url.Values)
	*r.queryValues = r.request.URL.Query()

	var err error
	r.filter, err = r.getFilters(opts.supportedFilters)
	if err != nil {
		return err
	}

	return nil
}

func (r *base) URL() *url.URL {
	return r.request.URL
}

func (r *base) populateFilters(target interface{}) error {
	filter := make(map[string]interface{})
	for k, v := range r.filter {
		filter[k] = v
	}

	err := figure.Out(target).From(filter).Please()
	if err != nil {
		f := errors.GetFields(err)
		return validation.Errors{
			toSnakeCase(cast.ToString(f["field"])): err,
		}
	}

	return nil
}

func (r *base) marshalQuery() string {
	var builder strings.Builder

	for key, values := range *r.queryValues {
		if !strings.Contains(key, "page[") {
			builder.WriteString(key + "=" + strings.Join(values, ",") + "&")
		}
	}

	return strings.TrimSuffix(builder.String(), "&")
}

func mkJsonTag(fieldName string) string {
	return fmt.Sprintf("json:\"%s\"", fieldName)
}

func toSnakeCase(str string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

func (r *base) ShouldFilter(name string) bool {
	_, ok := r.filter[name]
	return ok
}

// getString - tries to get string from URL param, if empty gets from query values
func (r *base) getString(name string) string {
	result := chi.URLParam(r.request, name)
	if result != "" {
		return strings.TrimSpace(result)
	}

	return strings.TrimSpace(r.queryValues.Get(name))
}

func (r *base) getUint64(name string) (uint64, error) {
	strVal := r.getString(name)
	if strVal == "" {
		return 0, nil
	}

	return strconv.ParseUint(strVal, 0, 64)
}

func (r *base) getUint64ID() (uint64, error) {
	result, err := r.getUint64("id")
	if err != nil {
		return 0, validation.Errors{
			"id": err,
		}
	}

	return result, nil
}

func (r *base) getInt64(name string) (int64, error) {
	strVal := r.getString(name)
	if strVal == "" {
		return 0, nil
	}

	return strconv.ParseInt(strVal, 0, 64)
}

func (r *base) getInt32(name string) (int32, error) {
	strVal := r.getString(name)
	if strVal == "" {
		return 0, nil
	}

	raw, err := strconv.ParseInt(strVal, 0, 32)
	if err != nil {
		return 0, errors.Wrap(err, "overflow during int32 parsing")
	}

	return int32(raw), nil
}

func (r *base) getFilters(supportedFilters map[string]struct{}) (map[string]string, error) {
	filters := make(map[string]string)
	for queryParam, values := range *r.queryValues {
		if strings.Contains(queryParam, "filter") {
			filterKey := strings.TrimPrefix(queryParam, "filter[")
			filterKey = strings.TrimSuffix(filterKey, "]")
			if _, supported := supportedFilters[filterKey]; !supported {
				return nil, validation.Errors{
					queryParam: errors.New(
						fmt.Sprintf("filter is not supported; supported values: %v",
							getSliceOfSupportedIncludes(supportedFilters)),
					),
				}
			}

			if len(values) == 0 {
				continue
			}

			if len(values) > 1 {
				return nil, validation.Errors{
					queryParam: errors.New("multiple values per one filter are not supported"),
				}
			}

			if values[0] == "" {
				continue
			}

			filters[filterKey] = values[0]
		}
	}

	return filters, nil
}

func (r *base) getIncludes(supportedIncludes map[string]struct{}) (map[string]struct{}, error) {
	const fieldName = "include"
	rawIncludes := r.getString(fieldName)
	if rawIncludes == "" {
		return nil, nil
	}
	includes := strings.Split(rawIncludes, ",")
	requestIncludes := map[string]struct{}{}
	for _, include := range includes {
		if _, supported := supportedIncludes[include]; !supported {
			return nil, validation.Errors{
				fieldName: errors.New(fmt.Sprintf("`%s` is not supported; supported values`: %v", include,
					getSliceOfSupportedIncludes(supportedIncludes))),
			}
		}

		// Note: Because compound documents require full linkage (except when relationship linkage is excluded by sparse fieldsets),
		// intermediate resources in a multi-part path must be returned along with the leaf nodes. For example,
		// a response to a request for comments.author should include comments as well as the author of each of those comments.
		subIncludes := strings.Split(include, ".")
		parentInclude := ""
		for i := range subIncludes {
			if parentInclude == "" {
				parentInclude = subIncludes[i]
			} else {
				parentInclude += "." + subIncludes[i]
			}
			requestIncludes[parentInclude] = struct{}{}
		}
	}

	return requestIncludes, nil
}

func getSliceOfSupportedIncludes(includes map[string]struct{}) []string {
	result := make([]string, 0, len(includes))
	for include := range includes {
		result = append(result, include)
	}

	return result
}

func (r *base) getLimit(defaultLimit, maxLimit uint64) (uint64, error) {
	result, err := r.getUint64(pageParamLimit)
	if err != nil {
		return 0, validation.Errors{
			pageParamLimit: errors.New("Must be a valid uint64 value"),
		}
	}

	if result == 0 {
		return defaultLimit, nil
	}

	if result > maxLimit {
		return 0, validation.Errors{
			pageParamLimit: fmt.Errorf("limit must not exceed %d", maxLimit),
		}
	}

	return result, nil
}
