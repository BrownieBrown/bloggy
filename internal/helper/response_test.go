package helper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithJSON(t *testing.T) {
	w := httptest.NewRecorder()
	payload := map[string]string{"message": "test"}

	RespondWithJSON(w, http.StatusOK, payload)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, resp.StatusCode)
	}

	var data map[string]string
	json.NewDecoder(resp.Body).Decode(&data)

	if data["message"] != "test" {
		t.Errorf("expected message %v, got %v", "test", data["message"])
	}
}

func TestRespondWithError(t *testing.T) {
	w := httptest.NewRecorder()

	RespondWithError(w, http.StatusBadRequest, "error")

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %v, got %v", http.StatusBadRequest, resp.StatusCode)
	}

	var data map[string]string
	json.NewDecoder(resp.Body).Decode(&data)

	if data["error"] != "error" {
		t.Errorf("expected error %v, got %v", "error", data["error"])
	}
}
