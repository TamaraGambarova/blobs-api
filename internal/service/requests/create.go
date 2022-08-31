package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/blobs/resources"
	"net/http"
)

func CreateNewBlobRequest(r *http.Request) (*resources.BlobsListResponse, error) {
	var request resources.BlobsListResponse

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to decode body")
	}

	return &request, ValidateCreateBlobRequest(&request)
}

func ValidateCreateBlobRequest(r *resources.BlobsListResponse) error {
	ownerLen := len(r.Data.Attributes.Owner)
	contentLen := len(r.Data.Attributes.Content)

	if ownerLen > 0 && contentLen > 0 {
		return nil
	} else {
		return problems.InternalError()
	}

}
