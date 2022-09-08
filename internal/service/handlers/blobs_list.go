package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/blobs/internal/data"
	"gitlab.com/tokend/blobs/internal/service/requests"
	"gitlab.com/tokend/blobs/resources"
	"net/http"
)

func ListOfBlobs(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blobQ := BlobsQ(r)
	applyFilters(blobQ, *request)
	blobs, err := blobQ.BlobsPage(&request.OffsetPageParams).Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to select blobs")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.BlobsListResponse{
		Data:  createBlobsData(blobs),
		Links: data.GetOffsetLinks(r, request.OffsetPageParams),
	})
}

func applyFilters(q data.Blobs, request requests.ListRequest) {
	if request.FilterOwner != nil && len(*request.FilterOwner) > 0 {
		q.GetByOwner(*request.FilterOwner)
	}
}

func createBlobsData(blobs []data.Blob) []resources.Blobs {
	result := make([]resources.Blobs, 0, len(blobs))
	for _, blob := range blobs {
		result = append(result, resources.Blobs{
			Key: resources.NewKeyInt64(blob.Id, resources.BLOBS),
			Attributes: resources.BlobsAttributes{
				Content: blob.Content,
			},
			Relationships: &resources.BlobsRelationships{
				Owner: resources.Relation{
					Data: &resources.Key{
						ID:   blob.Owner,
						Type: resources.OWNER,
					},
				},
			}})
	}
	return result
}
