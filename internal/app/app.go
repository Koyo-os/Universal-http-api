package app

import (
	"context"
	"net/http"

	"github.com/koyo-os/universal-http-api/internal/config"
	"github.com/koyo-os/universal-http-api/internal/handler"
	"github.com/koyo-os/universal-http-api/internal/server"
	"github.com/koyo-os/universal-http-api/pkg/loger"
)

func App(ctx context.Context) error {
	logger := loger.New()

	logger.Info().Msg("starting app...")

	cfg, err := config.New("config.toml")
	if err != nil{
		logger.Error().Err(err)
		return err
	}

	s := server.New(cfg)

	go func ()  {
		<-ctx.Done()
		s.Stop(ctx)	
	}()
	
	handle := handler.New(cfg)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handle.MainHandler)

	s.SetHandler(mux)

	return s.Run()
}