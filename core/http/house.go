package http

import (
	"log/slog"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/willoma/recherche-maison/web"
)

func (s *Server) createHousePage(w http.ResponseWriter, r *http.Request) {
	// Get cities for the dropdown
	cities, err := s.villeService.ListCities(r.Context())
	if err != nil {
		slog.Error("Failed to get cities", "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Render template
	component := web.CreateHousePage(cities)
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render create house page", "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}

func (s *Server) createHouse(w http.ResponseWriter, r *http.Request) {
	// New house form submission handler will be implemented later
}

func (s *Server) modifyHousePage(w http.ResponseWriter, r *http.Request) {
	// Get house ID from URL
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		slog.Error("Invalid house ID", "id", idStr, "error", err)
		http.Error(w, "Identifiant de maison invalide", http.StatusBadRequest)
		return
	}

	// Get house from database
	house, err := s.maisonService.GetHouse(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get house", "id", id, "error", err)
		http.Error(w, "Maison introuvable", http.StatusNotFound)
		return
	}

	// Get publication URLs
	publicationURLs, err := s.maisonService.GetPublicationURLs(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get publication URLs", "house_id", id, "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Get photos
	photos, err := s.maisonService.GetPhotos(id)
	if err != nil {
		slog.Error("Failed to get photos", "house_id", id, "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Get attachments
	attachments, err := s.maisonService.GetAttachments(id)
	if err != nil {
		slog.Error("Failed to get attachments", "house_id", id, "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Get cities for the dropdown
	cities, err := s.villeService.ListCities(r.Context())
	if err != nil {
		slog.Error("Failed to get cities", "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Render template
	component := web.ModifyHousePage(house, publicationURLs, photos, attachments, cities)
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render modify house page", "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}

func (s *Server) modifyHouse(w http.ResponseWriter, r *http.Request) {
	// Edit house form submission handler will be implemented later
}

func (s *Server) deleteHousePage(w http.ResponseWriter, r *http.Request) {
	// Get house ID from URL
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		slog.Error("Invalid house ID", "id", idStr, "error", err)
		http.Error(w, "Identifiant de maison invalide", http.StatusBadRequest)
		return
	}

	// Get house from database
	house, err := s.maisonService.GetHouse(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get house", "id", id, "error", err)
		http.Error(w, "Maison introuvable", http.StatusNotFound)
		return
	}

	// Render template
	component := web.DeleteHousePage(house)
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render delete house page", "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}

func (s *Server) deleteHouse(w http.ResponseWriter, r *http.Request) {
	// Delete house submission handler will be implemented later
}

func (s *Server) housePage(w http.ResponseWriter, r *http.Request) {
	// Get house ID from URL
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		slog.Error("Invalid house ID", "id", idStr, "error", err)
		http.Error(w, "Identifiant de maison invalide", http.StatusBadRequest)
		return
	}

	// Get house from database
	house, err := s.maisonService.GetHouse(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get house", "id", id, "error", err)
		http.Error(w, "Maison introuvable", http.StatusNotFound)
		return
	}

	// Get publication URLs
	publicationURLs, err := s.maisonService.GetPublicationURLs(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get publication URLs", "house_id", id, "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Get photos
	photos, err := s.maisonService.GetPhotos(id)
	if err != nil {
		slog.Error("Failed to get photos", "house_id", id, "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Get attachments
	attachments, err := s.maisonService.GetAttachments(id)
	if err != nil {
		slog.Error("Failed to get attachments", "house_id", id, "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Render template
	component := web.HousePage(house, publicationURLs, photos, attachments)
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render house page", "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}

func (s *Server) housePhoto(w http.ResponseWriter, r *http.Request) {
	houseID := r.PathValue("id")
	filename := r.PathValue("filename")

	filePath := filepath.Join(s.uploadsDir, houseID, "photos", filename)
	http.ServeFile(w, r, filePath)
}

func (s *Server) houseAttachment(w http.ResponseWriter, r *http.Request) {
	houseID := r.PathValue("id")
	filename := r.PathValue("filename")

	filePath := filepath.Join(s.uploadsDir, houseID, "attachments", filename)
	http.ServeFile(w, r, filePath)
}
