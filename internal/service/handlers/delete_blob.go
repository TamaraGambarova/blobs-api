package handlers

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/blobs/internal/data"
	"net/http"
	"strconv"
)

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	blob, _ := BlobsQ(r).GetByID(id).Get()
	if blob == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	err = BlobsQ(r).Transaction(func(q data.Blobs) error {
		err = BlobsQ(r).Delete(id)

		if err != nil {
			return errors.Wrap(err, "failed to delete blob")
		}

		return nil
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to execute db delete transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(204)
}
