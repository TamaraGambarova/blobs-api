package handlers

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/blobs/resources"
	"net/http"
	"strconv"
)

func GetBlobById(w http.ResponseWriter, r *http.Request) {
	id, convErr := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if convErr != nil {
		ape.RenderErr(w, problems.BadRequest(convErr)...)
		return
	}

	blobById, err := BlobsQ(r).GetByID(id).Get()

	if err != nil || blobById == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, resources.BlobsResponse{
		Data: resources.Blobs{
			Key: resources.Key{
				ID:   chi.URLParam(r, "id"),
				Type: resources.BLOBS,
			},
			Attributes: resources.BlobsAttributes{
				Content: blobById.Content,
			},
			Relationships: &resources.BlobsRelationships{
				Owner: resources.Relation{
					Data: &resources.Key{
						ID:   blobById.Owner,
						Type: resources.OWNER},
				}},
		},
	})
}
