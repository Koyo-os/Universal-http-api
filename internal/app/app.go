package app

import (
	"context"

	"github.com/koyo-os/universal-http-api/internal/config"
	"github.com/koyo-os/universal-http-api/internal/server"
	"github.com/koyo-os/universal-http-api/pkg/loger"
)

func App(ctx context.Context) error {
	logger := loger.New()
	cfg, err := config.New("config.toml")
	if err != nil{
		logger.Error().Err(err)
		return err
	}

	s := server.New(cfg)
	
}