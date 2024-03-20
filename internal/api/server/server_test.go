package server

import (
	"errors"
	"github.com/BrownieBrown/bloggy/internal/api/router"
	"github.com/BrownieBrown/bloggy/internal/models"
	"net/http"
	"testing"
)

const port = "8081"

func TestNewServer(t *testing.T) {
	server := createServer()
	if server == nil {
		t.Errorf("NewServer() failed, expected a server, got nil")
	}
	if server.server.Addr != ":"+port {
		t.Errorf("NewServer() failed, expected %v, got %v", ":"+port, server.server.Addr)
	}
}

func TestStart(t *testing.T) {
	server := createServer()
	go func() {
		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Errorf("Start() failed, expected nil or http.ErrServerClosed, got %v", err)
		}
	}()

	resp, err := http.Get("http://localhost:" + port)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Start() failed, expected status %v, got %v", http.StatusOK, resp.StatusCode)
	}
}

func createServer() *Server {
	cfg := loadConfig()
	r := router.NewRouter()
	r.Init()

	return NewServer(cfg, r)
}

func loadConfig() *models.Config {
	return &models.Config{
		ApiConfig: models.ApiConfig{
			Port: port,
		},
	}
}
