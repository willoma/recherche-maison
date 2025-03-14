package http

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/willoma/recherche-maison/core/fichier"
	"github.com/willoma/recherche-maison/core/maison"
	"github.com/willoma/recherche-maison/core/ville"
	"github.com/willoma/recherche-maison/static"
)

// Server handles HTTP requests for the application
type Server struct {
	fileService   *fichier.Service
	maisonService *maison.Service
	villeService  *ville.Service
	uploadsDir    string
	port          int
}

// NewServer creates a new HTTP server
func NewServer(fileService *fichier.Service, maisonService *maison.Service, villeService *ville.Service, uploadsDir string, port int) *Server {
	return &Server{
		fileService:   fileService,
		maisonService: maisonService,
		villeService:  villeService,
		uploadsDir:    uploadsDir,
		port:          port,
	}
}

// Run starts the HTTP server
func Run(fileService *fichier.Service, maisonService *maison.Service, villeService *ville.Service, uploadsDir string, port int) {
	server := NewServer(fileService, maisonService, villeService, uploadsDir, port)
	server.Start()
}

// Start configures and starts the HTTP server
func (s *Server) Start() {
	mux := http.NewServeMux()

	// Static files
	mux.HandleFunc("GET /script.js", static.ServeScript)
	mux.HandleFunc("GET /style.css", static.ServeStyle)

	// Main page
	mux.HandleFunc("GET /{$}", s.handleMainPage)

	// House routes
	mux.HandleFunc("GET /maisons/nouvelle", s.createHousePage)
	mux.HandleFunc("POST /maisons/nouvelle", s.createHouse)
	mux.HandleFunc("GET /maisons/{id}/modifier", s.modifyHousePage)
	mux.HandleFunc("POST /maisons/{id}/modifier", s.modifyHouse)
	mux.HandleFunc("GET /maisons/{id}/supprimer", s.deleteHousePage)
	mux.HandleFunc("POST /maisons/{id}/supprimer", s.deleteHouse)
	mux.HandleFunc("GET /maisons/{id}", s.housePage)
	mux.HandleFunc("GET /maisons/{id}/uploads/{filename}", s.houseFile)

	// City routes
	mux.HandleFunc("GET /villes", s.modifyCitiesPage)
	mux.HandleFunc("POST /villes", s.modifyCities)

	// Start server
	addr := fmt.Sprintf(":%d", s.port)
	slog.Info("Starting server", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Server error", "error", err)
		os.Exit(1)
	}
}
