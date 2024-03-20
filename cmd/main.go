package main

import (
	"github.com/BrownieBrown/bloggy/internal/api/middleware"
	"github.com/BrownieBrown/bloggy/internal/api/router"
	"github.com/BrownieBrown/bloggy/internal/api/server"
	"github.com/BrownieBrown/bloggy/internal/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")

	}

	cfg := config.LoadConfig()
	r := router.NewRouter()
	r.Init()
	corsRouter := middleware.Cors(r)

	srv := server.NewServer(cfg, corsRouter)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
