package router

import (
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
	router.Init()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
