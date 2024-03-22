package config

import (
	"errors"
	"github.com/BrownieBrown/bloggy/internal/models"
	"os"
)

func LoadConfig() (*models.Config, error) {
	apiConfig, err := loadApiConfig()
	if err != nil {
		return nil, err
	}

	postgresConfig, err := loadPostgresConfig()
	if err != nil {
		return nil, err
	}

	return &models.Config{
		ApiConfig:      &apiConfig,
		PostgresConfig: &postgresConfig,
	}, nil
}

func loadApiConfig() (models.ApiConfig, error) {
	config := models.ApiConfig{
		Port: os.Getenv("PORT"),
	}

	if config.Port == "" {
		return config, errors.New("PORT is required")
	}

	return config, nil
}

func loadPostgresConfig() (models.PostgresConfig, error) {
	config := models.PostgresConfig{
		Host:       os.Getenv("POSTGRES_HOST"),
		Port:       os.Getenv("POSTGRES_PORT"),
		User:       os.Getenv("POSTGRES_USER"),
		Password:   os.Getenv("POSTGRES_PASSWORD"),
		Dbname:     os.Getenv("POSTGRES_DB"),
		TestDbname: os.Getenv("POSTGRES_TEST_DB"),
		Url:        os.Getenv("POSTGRES_URL"),
		SSLMode:    os.Getenv("POSTGRES_SSL_MODE"),
	}

	if config.Host == "" {
		return config, errors.New("POSTGRES_HOST is required")
	}

	if config.Port == "" {
		return config, errors.New("POSTGRES_PORT is required")
	}

	if config.User == "" {
		return config, errors.New("POSTGRES_USER is required")
	}

	if config.Password == "" {
		return config, errors.New("POSTGRES_PASSWORD is required")
	}

	if config.Dbname == "" {
		return config, errors.New("POSTGRES_DB is required")
	}

	if config.Url == "" {
		return config, errors.New("POSTGRES_URL is required")
	}

	if config.SSLMode == "" {
		return config, errors.New("POSTGRES_SSL_MODE is required")
	}

	return config, nil
}
