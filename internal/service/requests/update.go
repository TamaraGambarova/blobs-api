package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/blobs/resources"
	"net/http"
)

func UpdateBlobRequest(r *http.Request) (*resources.BlobsResponse, error) {
	var request resources.BlobsResponse

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to decode body")
	}

	return &request, ValidateBlobRequest(&request)
}
