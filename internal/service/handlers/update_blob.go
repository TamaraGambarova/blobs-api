package handlers

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/blobs/internal/data"
	"gitlab.com/tokend/blobs/internal/service/requests"
	"gitlab.com/tokend/blobs/resources"
	"net/http"
	"strconv"
)

func UpdateBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.UpdateBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	id, convErr := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if convErr != nil {
		ape.RenderErr(w, problems.BadRequest(convErr)...)
		return
	}
	blob, _ := BlobsQ(r).GetByID(id).Get()
	if blob == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	unMarshaledContent, err := request.Data.Attributes.Content.MarshalJSON()
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	err = BlobsQ(r).Transaction(func(q data.Blobs) error {
		err = BlobsQ(r).Update(
			id,
			&data.Blob{
				Content: string(unMarshaledContent),
			})

		if err != nil {
			return errors.Wrap(err, "failed to update blob")
		}

		return nil
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to execute db update transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.BlobsResponse{
		Data: createBlobInfo(chi.URLParam(r, "id"), request),
	})
}
