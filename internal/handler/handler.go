package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/koyo-os/universal-http-api/internal/config"
	"github.com/koyo-os/universal-http-api/pkg/loger"
)

type Handler struct{
	logger loger.Logger
	config *config.Config
}

func New(cfg *config.Config) *Handler {
	return &Handler{
		logger: loger.New(),
		config: cfg,
	}
}

func (h *Handler) createReverseProxy(target string) *httputil.ReverseProxy {
	targetURL, err := url.Parse(target)
	if err != nil {
		h.logger.Error().Err(err)
		return nil
	}
	return httputil.NewSingleHostReverseProxy(targetURL)
}

func (h *Handler) MainHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.Split(path, "/")[0]

	for _, v := range h.config.Urls {
		if path == v.UrlPrefix {
			proxy := h.createReverseProxy(v.ServiceAddr)
			proxy.ServeHTTP(w, r)
		}
	}
}