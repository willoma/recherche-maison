package http

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/willoma/recherche-maison/core/city"
	"github.com/willoma/recherche-maison/web"
)

// modifyCitiesPage renders the page for managing cities
func (s *Server) modifyCitiesPage(w http.ResponseWriter, r *http.Request) {
	houses, err := s.houseService.ListHouses(r.Context())
	if err != nil {
		slog.Error("Failed to get houses", "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Get all cities
	cities, err := s.cityService.ListCities(r.Context())
	if err != nil {
		slog.Error("Failed to get cities", "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Render template
	component := web.CityManagementPage(cities, houses)
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render city management page", "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}

func (s *Server) modifyCities(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		slog.Error("Failed to parse form", "error", err)
		http.Error(w, "Erreur lors de la soumission du formulaire", http.StatusBadRequest)
		return
	}

	// Get action type
	action := r.FormValue("action")

	switch action {
	case "create":
		// Handle city creation
		cityName := r.FormValue("city_name")
		if cityName == "" {
			slog.Error("City name is required")
			http.Error(w, "Le nom de la ville est obligatoire", http.StatusBadRequest)
			return
		}

		// Create city
		err := s.cityService.CreateCity(r.Context(), cityName)
		if err != nil {
			slog.Error("Failed to create city", "name", cityName, "error", err)
			http.Error(w, "Erreur lors de la création de la ville", http.StatusInternalServerError)
			return
		}

	case "delete":
		// Handle city deletion
		cityIDStr := r.FormValue("city_id")
		if cityIDStr == "" {
			slog.Error("City ID is required")
			http.Error(w, "L'identifiant de la ville est obligatoire", http.StatusBadRequest)
			return
		}

		// Parse city ID
		cityID, err := strconv.ParseInt(cityIDStr, 10, 64)
		if err != nil {
			slog.Error("Invalid city ID", "id", cityIDStr, "error", err)
			http.Error(w, "Identifiant de ville invalide", http.StatusBadRequest)
			return
		}

		// Delete city
		err = s.cityService.DeleteCity(r.Context(), cityID)
		if err != nil {
			if errors.Is(err, city.ErrCityInUse) {
				slog.Error("Cannot delete city that is used by houses", "id", cityID)
				http.Error(w, "La ville est utilisée par des maisons et ne peut pas être supprimée", http.StatusBadRequest)
				return
			}
			slog.Error("Failed to delete city", "id", cityID, "error", err)
			http.Error(w, "Erreur lors de la suppression de la ville", http.StatusInternalServerError)
			return
		}

	default:
		slog.Error("Invalid action", "action", action)
		http.Error(w, "Action invalide", http.StatusBadRequest)
		return
	}

	// Redirect back to city management page
	http.Redirect(w, r, "/villes", http.StatusSeeOther)
}
