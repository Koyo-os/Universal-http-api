package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/koyo-os/universal-http-api/internal/config"
	"github.com/koyo-os/universal-http-api/pkg/loger"
)

type Server struct{
	*http.Server
	logger loger.Logger
}

func New(cfg *config.Config) *Server {
	return &Server{
		&http.Server{
			Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		},
		loger.New(),
	}
}

func (s *Server) SetHandler(mux *http.ServeMux) {
	s.Handler = mux
}

func (s *Server) Run() error {
	s.logger.Info().Msg("starting server...")

	return s.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info().Msg("stopping server...")

	return s.Shutdown(ctx)
}

