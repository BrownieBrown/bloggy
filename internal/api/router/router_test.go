package router

import (
	"github.com/BrownieBrown/bloggy/internal/api/handler"
	"github.com/BrownieBrown/bloggy/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouter(t *testing.T) {
	router := NewRouter()
	if router == nil {
		t.Errorf("NewRouter() failed, expected a router, got nil")
	}
	if router == nil {
		t.Errorf("NewRouter() failed, expected a ServeMux, got nil")
	}
}

func TestInit(t *testing.T) {
	router := NewRouter()
	hh := &handler.HealthHandler{
		Config: &models.Config{},
	}
	eh := &handler.ErrorHandler{
		Config: &models.Config{},
	}

	router.Init(hh, eh)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	req, err = http.NewRequest("GET", "/v1/readiness", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	req, err = http.NewRequest("GET", "/v1/err", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}
