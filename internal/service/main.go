package service

import (
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/tokend/blobs/internal/config"
	"gitlab.com/tokend/blobs/internal/data"
	"gitlab.com/tokend/blobs/internal/data/pg"

	"net"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	blobs    data.Blobs
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)

}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		blobs:    pg.NewBlobs(cfg.DB().Clone()),
	}
}
func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
