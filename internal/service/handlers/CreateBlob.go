package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/blobs/internal/data"
	"gitlab.com/tokend/blobs/internal/service/requests"
	"gitlab.com/tokend/blobs/resources"
	"net/http"
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

	ape.Render(w, resources.BlobsListResponse{
		Data: createBlobInfo(id, request),
	})

}

func createBlobInfo(blobId string, request *resources.BlobsListResponse) resources.BlobsList {
	return resources.BlobsList{
		Key: resources.NewKey(blobId, resources.BLOBS),
		Attributes: resources.BlobsListAttributes{
			Id:      request.Data.Attributes.Id,
			Owner:   request.Data.Attributes.Owner,
			Content: request.Data.Attributes.Content,
		},
	}
}
