package main

import (
	"log/slog"
	"os"

	_ "modernc.org/sqlite"

	"github.com/willoma/recherche-maison/core/city"
	"github.com/willoma/recherche-maison/core/file"
	"github.com/willoma/recherche-maison/core/house"
	"github.com/willoma/recherche-maison/core/http"
	"github.com/willoma/recherche-maison/db"
)

const (
	dbPath     = "recherche-maison.db"
	dbOptions  = "_pragma=foreign_keys(1)&_time_format=sqlite"
	uploadsDir = "uploads"
	port       = 8910
)

func main() {
	// Ensure uploads directory exists
	if err := os.MkdirAll(uploadsDir, 0o755); err != nil {
		slog.Error("Failed to create uploads directory", "error", err)
		os.Exit(1)
	}

	// Initialize database
	dbConn, err := db.Init(dbPath, dbOptions)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	// Initialize services
	queries := db.New(dbConn)
	fileService := file.NewService(uploadsDir)
	houseService := house.NewService(queries, uploadsDir)
	cityService := city.NewService(queries)

	http.Run(fileService, houseService, cityService, uploadsDir, port)
}
