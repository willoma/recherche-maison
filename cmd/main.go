package main

import (
	"log/slog"
	"os"

	_ "modernc.org/sqlite"

	"github.com/willoma/recherche-maison/config"
	"github.com/willoma/recherche-maison/core/city"
	"github.com/willoma/recherche-maison/core/file"
	"github.com/willoma/recherche-maison/core/house"
	"github.com/willoma/recherche-maison/core/http"
	"github.com/willoma/recherche-maison/db"
)

func main() {
	// Ensure uploads directory exists
	if err := os.MkdirAll(config.UploadsDir, 0o755); err != nil {
		slog.Error("Failed to create uploads directory", "error", err)
		os.Exit(1)
	}

	// Initialize database
	dbConn, err := db.Init()
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	// Initialize services
	queries := db.New(dbConn)
	fileService := file.NewService()

	houseService := house.NewService(queries, dbConn)
	cityService := city.NewService(queries)

	http.Run(fileService, houseService, cityService)
}
