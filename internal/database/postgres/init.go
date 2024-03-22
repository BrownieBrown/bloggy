package postgres

import (
	"database/sql"
	"fmt"
	"github.com/BrownieBrown/bloggy/internal/database"
	"github.com/BrownieBrown/bloggy/internal/models"
)

func Init(cfg *models.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.PostgresConfig.Host,
		cfg.PostgresConfig.User,
		cfg.PostgresConfig.Password,
		cfg.PostgresConfig.Dbname,
		cfg.PostgresConfig.Port,
		cfg.PostgresConfig.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the PostgreSQL database.")

	cfg.ApiConfig.DB = database.New(db)

	return db, nil
}
