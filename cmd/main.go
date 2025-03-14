package main

import (
	"log/slog"
	"os"

	_ "modernc.org/sqlite"

	"github.com/willoma/recherche-maison/core/fichier"
	"github.com/willoma/recherche-maison/core/http"
	"github.com/willoma/recherche-maison/core/maison"
	"github.com/willoma/recherche-maison/core/ville"
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
	fileService := fichier.NewService(uploadsDir)
	maisonService := maison.NewService(queries, uploadsDir)
	villeService := ville.NewService(queries)

	http.Run(fileService, maisonService, villeService, uploadsDir, port)
}
