package models

import "github.com/BrownieBrown/bloggy/internal/database"

type Config struct {
	ApiConfig      *ApiConfig
	PostgresConfig *PostgresConfig
}

type ApiConfig struct {
	Port string
	DB   *database.Queries
}

type PostgresConfig struct {
	Host       string
	Port       string
	User       string
	Password   string
	Dbname     string
	TestDbname string
	Url        string
	SSLMode    string
}
