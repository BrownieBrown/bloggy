package server

import (
	"github.com/BrownieBrown/bloggy/internal/models"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(cfg *models.Config, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:    ":" + cfg.ApiConfig.Port,
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
