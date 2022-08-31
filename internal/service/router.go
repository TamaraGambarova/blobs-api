package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/blobs/internal/service/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxBlobsQ(s.blobs),
		),
	)
	r.Route("/integrations", func(r chi.Router) {
		r.Get("/blobs", handlers.ListOfBlobs)
		r.Post("/blobs", handlers.CreateBlob)
		r.Patch("/blobs/{id}", handlers.UpdateBlob)
		r.Delete("/blobs/{id}", handlers.DeleteBlob)
	})

	return r
}
