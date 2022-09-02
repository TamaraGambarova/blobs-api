package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/blobs/internal/data"
	"gitlab.com/tokend/blobs/internal/service/requests"
	"gitlab.com/tokend/blobs/resources"
	"net/http"
	"strconv"
)

func CreateBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.CreateNewBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	id, err := BlobsQ(r).Create(&data.Blob{
		Owner:   request.Data.Attributes.Owner,
		Content: request.Data.Attributes.Content,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to create blob")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.BlobsResponse{
		Data: createBlobInfo(id, request),
	})

}

func createBlobInfo(blobId string, request *resources.BlobsResponse) resources.Blobs {
	convertedId, _ := strconv.ParseInt(blobId, 10, 64)

	return resources.Blobs{
		Key: resources.NewKeyInt64(convertedId, resources.BLOBS),
		Attributes: resources.BlobsAttributes{
			Owner:   request.Data.Attributes.Owner,
			Content: request.Data.Attributes.Content,
		},
	}
}