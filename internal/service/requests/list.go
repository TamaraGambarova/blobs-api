package requests

import (
  "gitlab.com/tokend/blobs/internal/data"
  "net/http"

)

type ListRequest struct {
	*base
	PageParams *data.PageParams
}

func NewListRequest(r *http.Request) (*ListRequest, error) {
	b, err := newBase(r, baseOpts{})
	if err != nil {
		return nil, err
	}

	params, err := data.GetPageParams(r)

	request := ListRequest{
		base:       b,
		PageParams: params,
	}

	return &request, nil
}
