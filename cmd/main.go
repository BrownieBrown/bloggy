package main

import (
	"database/sql"
	"github.com/BrownieBrown/bloggy/internal/api/handler"
	"github.com/BrownieBrown/bloggy/internal/api/middleware"
	"github.com/BrownieBrown/bloggy/internal/api/router"
	"github.com/BrownieBrown/bloggy/internal/api/server"
	"github.com/BrownieBrown/bloggy/internal/config"
	"github.com/BrownieBrown/bloggy/internal/database/postgres"
	"github.com/joho/godotenv"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")

	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := postgres.Init(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}(db)

	hh := handler.NewHealthHandler(cfg)
	eh := handler.NewErrorHandler(cfg)
	uh := handler.NewUserHandler(cfg)
	r := router.NewRouter()
	r.Init(hh, eh, uh)
	corsRouter := middleware.Cors(r)

	srv := server.NewServer(cfg, corsRouter)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
