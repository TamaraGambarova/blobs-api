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
	id := chi.URLParam(r, "id")

	a, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = BlobsQ(r).Transaction(func(q data.Blobs) error {
		err = BlobsQ(r).Delete(a)

		if err != nil {
			return errors.Wrap(err, "failed to update blob")
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
