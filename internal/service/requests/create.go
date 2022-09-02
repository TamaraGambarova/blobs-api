package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/blobs/resources"
	"net/http"
)

func CreateNewBlobRequest(r *http.Request) (*resources.BlobsResponse, error) {
	var request resources.BlobsResponse

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to decode body")
	}

	return &request, ValidateBlobRequest(&request)
}

func ValidateBlobRequest(r *resources.BlobsResponse) error {
	return validation.Errors{
		"/data/attributes/owner": validation.Validate(&r.Data.Attributes.Owner, validation.Required),
		"/data/attributes/content": validation.Validate(&r.Data.Attributes.Content, validation.Required),
	}.Filter()
}
