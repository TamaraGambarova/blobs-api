package requests

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type ListRequest struct {
	pgdb.OffsetPageParams
	FilterOwner *string `filter:"owner"`
}

func NewListRequest(r *http.Request) (*ListRequest, error) {
	request := ListRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return &request, err
	}

	return &request, nil
}
