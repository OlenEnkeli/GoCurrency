package api

import (
	"context"
	"net"
	"net/http"

	"github.com/OlenEnkeli/GoCurrency/internal/settings"
)

type Server struct {
	baseServer *http.Server
}

func NewServer() *Server {
	router := NewRouter()
	router.Use(CORSMiddleware())

	return &Server{
		baseServer: &http.Server{
			Addr:           net.JoinHostPort(settings.Settings.API.Host, settings.Settings.API.Port),
			Handler:        router,
			MaxHeaderBytes: settings.Settings.API.MaxHeaderBytes,
			ReadTimeout:    settings.Settings.API.ReadTimeout,
			WriteTimeout:   settings.Settings.API.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.baseServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.baseServer.Shutdown(ctx)
}
