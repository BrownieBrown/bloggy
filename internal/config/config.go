package config

import (
	"github.com/BrownieBrown/bloggy/internal/models"
	"os"
)

func LoadConfig() *models.Config {
	return &models.Config{
		ApiConfig: models.ApiConfig{
			Port: os.Getenv("PORT"),
		},
	}
}
