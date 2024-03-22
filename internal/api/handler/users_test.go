package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/BrownieBrown/bloggy/internal/config"
	"github.com/BrownieBrown/bloggy/internal/database"
	"github.com/BrownieBrown/bloggy/internal/models"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

func TestCreateUserIntegration(t *testing.T) {
	if err := godotenv.Load("/Users/marcobraun/dev/bloggy/.env"); err != nil {
		log.Fatal("Error loading .env file")

	}
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatal(err)

	}
	db := createTestDatabase(t, cfg)
	defer db.Close()

	uh := NewUserHandler(cfg)

	req, err := http.NewRequest("POST", "/v1/users", strings.NewReader(`{"name":"test user"}`))
	if err != nil {
		t.Fatalf("Failed to make a request: %v", err)
	}
	rr := httptest.NewRecorder()
	uh.CreateUser(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Decode the response body
	var resp UserResponse
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Check the name field in the response
	if resp.Name != "test user" {
		t.Errorf("Handler returned unexpected body: got name %v want name %v", resp.Name, "test user")
	}
}

func TestCreateUserIntegrationBadRequest(t *testing.T) {
	if err := godotenv.Load("/Users/marcobraun/dev/bloggy/.env"); err != nil {
		log.Fatal("Error loading .env file")

	}
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatal(err)

	}
	db := createTestDatabase(t, cfg)
	defer db.Close()

	cfg.ApiConfig.DB = database.New(db)
	uh := NewUserHandler(cfg)

	req, err := http.NewRequest("POST", "/v1/users", strings.NewReader(`{"name":}`))
	if err != nil {
		t.Fatalf("Failed to make a request: %v", err)
	}
	rr := httptest.NewRecorder()
	uh.CreateUser(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func createTestDatabase(t *testing.T, cfg *models.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.PostgresConfig.Host, cfg.PostgresConfig.Port, cfg.PostgresConfig.User, cfg.PostgresConfig.Password, cfg.PostgresConfig.TestDbname, cfg.PostgresConfig.SSLMode)
	dn := "postgres"
	db, err := sql.Open(dn, dsn)
	if err != nil {
		t.Fatal(err)
	}

	cfg.ApiConfig.DB = database.New(db)

	return db
}
