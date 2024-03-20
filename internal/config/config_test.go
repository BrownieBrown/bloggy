package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("PORT", "8080")
	defer os.Unsetenv("PORT")

	cfg := LoadConfig()

	if cfg.ApiConfig.Port != "8080" {
		t.Errorf("Expected port to be 8080, got %s", cfg.ApiConfig.Port)
	}
}
