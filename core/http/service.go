package http

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/willoma/recherche-maison/config"
	"github.com/willoma/recherche-maison/core/city"
	"github.com/willoma/recherche-maison/core/file"
	"github.com/willoma/recherche-maison/core/house"
	"github.com/willoma/recherche-maison/static"
)

// Server handles HTTP requests for the application
type Server struct {
	fileService  *file.Service
	houseService *house.Service
	cityService  *city.Service
}

// NewServer creates a new HTTP server
func NewServer(fileService *file.Service, houseService *house.Service, cityService *city.Service) *Server {
	return &Server{
		fileService:  fileService,
		houseService: houseService,
		cityService:  cityService,
	}
}

// Run starts the HTTP server
func Run(fileService *file.Service, houseService *house.Service, cityService *city.Service) {
	server := NewServer(fileService, houseService, cityService)
	server.Start()
}

// Start starts the HTTP server
func (s *Server) Start() {
	mux := http.NewServeMux()
	s.registerRoutes(mux)
	s.startServer(mux)
}

// registerRoutes registers all HTTP routes
func (s *Server) registerRoutes(mux *http.ServeMux) {
	// Static files
	mux.HandleFunc("GET /script.js", static.ServeScript)
	mux.HandleFunc("GET /style.css", static.ServeStyle)

	// Main page
	mux.HandleFunc("GET /{$}", s.mainPage)

	// House routes
	mux.HandleFunc("GET /maison/creer", s.createHousePage)
	mux.HandleFunc("POST /maison/creer", s.createHouse)
	mux.HandleFunc("GET /maison/{id}", s.housePage)
	mux.HandleFunc("GET /maison/{id}/modifier", s.modifyHousePage)
	mux.HandleFunc("POST /maison/{id}/modifier", s.modifyHouse)
	mux.HandleFunc("POST /maison/{id}/supprimer", s.deleteHouse)
	mux.HandleFunc("GET /maison/{id}/photos/{filename}", s.housePhoto)
	mux.HandleFunc("GET /maison/{id}/piecesjointes/{filename}", s.houseAttachment)

	// City routes
	mux.HandleFunc("GET /villes", s.modifyCitiesPage)
	mux.HandleFunc("POST /villes", s.modifyCities)
}

// startServer starts the HTTP server
func (s *Server) startServer(mux *http.ServeMux) {
	slog.Info("Starting server", "addr", fmt.Sprintf(":%d", config.Port))
	if err := http.ListenAndServe(":"+fmt.Sprintf("%d", config.Port), mux); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Server error", "error", err)
		os.Exit(1)
	}
}
