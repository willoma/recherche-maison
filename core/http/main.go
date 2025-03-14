package http

import (
	"log/slog"
	"net/http"

	"github.com/willoma/recherche-maison/web"
)

// handleMainPage renders the main page with the list of houses
func (s *Server) handleMainPage(w http.ResponseWriter, r *http.Request) {
	// Get houses from database
	houses, err := s.maisonService.ListHouses(r.Context())
	if err != nil {
		slog.Error("Failed to get houses", "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Render template
	component := web.MainPage(houses)
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render main page", "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}
