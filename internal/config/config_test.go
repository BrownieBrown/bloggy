package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_USER", "user")
	os.Setenv("POSTGRES_PASSWORD", "password")
	os.Setenv("POSTGRES_DB", "dbname")
	os.Setenv("POSTGRES_URL", "url")
	os.Setenv("POSTGRES_SSL_MODE", "disable")

	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("POSTGRES_HOST")
		os.Unsetenv("POSTGRES_PORT")
		os.Unsetenv("POSTGRES_USER")
		os.Unsetenv("POSTGRES_PASSWORD")
		os.Unsetenv("POSTGRES_DB")
		os.Unsetenv("POSTGRES_URL")
		os.Unsetenv("POSTGRES_SSL_MODE")
	}()

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cfg.ApiConfig.Port != "8080" {
		t.Errorf("Expected port to be 8080, got %s", cfg.ApiConfig.Port)
	}

	if cfg.PostgresConfig.Host != "localhost" {
		t.Errorf("Expected host to be localhost, got %s", cfg.PostgresConfig.Host)
	}

}

func TestLoadConfig_MissingEnv(t *testing.T) {
	_, err := LoadConfig()
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
