package http

import (
	"log/slog"
	"net/http"

	"github.com/willoma/recherche-maison/web"
)

// modifyCitiesPage renders the page for managing cities
func (s *Server) modifyCitiesPage(w http.ResponseWriter, r *http.Request) {
	// Get all cities
	cities, err := s.villeService.ListCities(r.Context())
	if err != nil {
		slog.Error("Failed to get cities", "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// For each city, check if it's used by any house
	usedCities := make(map[int64]bool)
	for _, city := range cities {
		used, err := s.villeService.IsCityUsedByHouses(r.Context(), city.ID)
		if err != nil {
			slog.Error("Failed to check if city is used", "city_id", city.ID, "error", err)
			continue
		}
		usedCities[city.ID] = used
	}

	// Render template
	component := web.CityManagementPage(cities, usedCities)
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render city management page", "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}

func (s *Server) modifyCities(w http.ResponseWriter, r *http.Request) {
	// City management form submission handler will be implemented later
}
