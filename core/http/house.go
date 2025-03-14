package http

import (
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/willoma/recherche-maison/config"
	"github.com/willoma/recherche-maison/core/house"
	"github.com/willoma/recherche-maison/models"
	"github.com/willoma/recherche-maison/web"
)

func (s *Server) createHousePage(w http.ResponseWriter, r *http.Request) {
	// Get cities for the dropdown
	cities, err := s.cityService.ListCities(r.Context())
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
	// Parse form data
	houseForm, err, errMsg := parseHouseForm(r)
	if err != nil {
		slog.Error("Failed to parse house form", "error", err)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	// Create the house and get its ID
	houseID, err := s.houseService.CreateHouse(r.Context(), houseForm)
	if err != nil {
		slog.Error("Failed to create house", "error", err)
		http.Error(w, "Erreur lors de la création de la maison", http.StatusInternalServerError)
		return
	}

	// Handle publication URLs
	urls := r.Form["pub_url[]"]
	dates := r.Form["pub_date[]"]

	slog.Info("FORM", "form", r.Form)

	if len(urls) == 0 {
		slog.Error("At least one publication URL is required")
		http.Error(w, "Au moins une URL de publication est obligatoire", http.StatusBadRequest)
		return
	}

	// Add publication URLs
	for i, url := range urls {
		if url == "" {
			continue
		}

		// Skip if we don't have a matching date
		if i >= len(dates) || dates[i] == "" {
			continue
		}

		// Parse the date
		date, err := house.ParsePublicationDate(dates[i])
		if err != nil {
			slog.Error("Invalid publication date", "date", dates[i], "error", err)
			continue
		}

		// Create publication URL
		err = s.houseService.AddPublicationURL(r.Context(), houseID, url, date)
		if err != nil {
			slog.Error("Failed to add publication URL", "house_id", houseID, "url", url, "error", err)
		}
	}

	// Handle photo uploads
	// TODO: Implement photo uploads

	// Redirect to house details page
	http.Redirect(w, r, "/maison/"+strconv.FormatInt(houseID, 10), http.StatusSeeOther)
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
	house, err := s.houseService.GetHouse(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get house", "id", id, "error", err)
		http.Error(w, "Maison introuvable", http.StatusNotFound)
		return
	}

	// Get publication URLs
	publicationURLs, err := s.houseService.GetPublicationURLs(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get publication URLs", "house_id", id, "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Get photos
	photos, err := s.houseService.GetPhotos(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get photos", "house_id", id, "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Get attachments
	attachments, err := s.houseService.GetAttachments(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get attachments", "house_id", id, "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Get cities for the dropdown
	cities, err := s.cityService.ListCities(r.Context())
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
	// Get house ID from URL path
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		slog.Error("Invalid house ID", "error", err)
		http.Error(w, "Identifiant de maison invalide", http.StatusBadRequest)
		return
	}

	// Parse form data
	houseForm, err, errMsg := parseHouseForm(r)
	if err != nil {
		slog.Error("Failed to parse house form", "error", err)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	// Update the house in the database
	err = s.houseService.UpdateHouse(r.Context(), id, houseForm)
	if err != nil {
		slog.Error("Failed to update house", "error", err)
		http.Error(w, "Erreur lors de la mise à jour de la maison", http.StatusInternalServerError)
		return
	}

	// Redirect to house page
	http.Redirect(w, r, fmt.Sprintf("/maison/%d", id), http.StatusSeeOther)
	return
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
	house, err := s.houseService.GetHouse(r.Context(), id)
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
	house, err := s.houseService.GetHouse(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get house", "id", id, "error", err)
		http.Error(w, "Maison introuvable", http.StatusNotFound)
		return
	}

	// Get publication URLs
	publicationURLs, err := s.houseService.GetPublicationURLs(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get publication URLs", "house_id", id, "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Get photos
	photos, err := s.houseService.GetPhotos(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get photos", "house_id", id, "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Get attachments
	attachments, err := s.houseService.GetAttachments(r.Context(), id)
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

	filePath := filepath.Join(config.UploadsDir, houseID, "photos", filename)
	http.ServeFile(w, r, filePath)
}

func (s *Server) houseAttachment(w http.ResponseWriter, r *http.Request) {
	houseID := r.PathValue("id")
	filename := r.PathValue("filename")

	filePath := filepath.Join(config.UploadsDir, houseID, "attachments", filename)
	http.ServeFile(w, r, filePath)
}

// parseHouseForm parses the form data for house creation and modification
// Returns the parsed house data, an error if parsing fails, and a translated error message
func parseHouseForm(r *http.Request) (models.House, error, string) {
	var houseForm models.House
	var err error

	// Parse multipart form data (for file uploads)
	if err := r.ParseMultipartForm(config.MaxUploadSize); err != nil {
		return houseForm, err, "Erreur lors de la soumission du formulaire"
	}

	// Parse title
	houseForm.Title = r.FormValue("title")
	if houseForm.Title == "" {
		return houseForm, fmt.Errorf("title is required"), "Le titre est obligatoire"
	}

	// Parse city ID
	cityIDStr := r.FormValue("city_id")
	if cityIDStr == "" {
		return houseForm, fmt.Errorf("city ID is required"), "L'identifiant de ville est obligatoire"
	}

	houseForm.CityID, err = strconv.ParseInt(cityIDStr, 10, 64)
	if err != nil {
		return houseForm, fmt.Errorf("invalid city ID: %w", err), "Identifiant de ville invalide"
	}

	// Parse address (optional)
	houseForm.Address = r.FormValue("address")

	// Parse price
	priceStr := r.FormValue("price")
	if priceStr != "" {
		houseForm.Price, err = strconv.ParseInt(priceStr, 10, 64)
		if err != nil {
			return houseForm, fmt.Errorf("invalid price: %w", err), "Prix invalide"
		}
	}

	// Parse surface
	surfaceStr := r.FormValue("surface")
	if surfaceStr != "" {
		houseForm.Surface, err = strconv.ParseInt(surfaceStr, 10, 64)
		if err != nil {
			return houseForm, fmt.Errorf("invalid surface: %w", err), "Surface invalide"
		}
	}

	// Parse rooms
	roomsStr := r.FormValue("rooms")
	if roomsStr != "" {
		houseForm.Rooms, err = strconv.ParseInt(roomsStr, 10, 64)
		if err != nil {
			return houseForm, fmt.Errorf("invalid rooms: %w", err), "Nombre de pièces invalide"
		}
	}

	// Parse bedrooms
	bedroomsStr := r.FormValue("bedrooms")
	if bedroomsStr != "" {
		houseForm.Bedrooms, err = strconv.ParseInt(bedroomsStr, 10, 64)
		if err != nil {
			return houseForm, fmt.Errorf("invalid bedrooms: %w", err), "Nombre de chambres invalide"
		}
	}

	// Parse bathrooms
	bathroomsStr := r.FormValue("bathrooms")
	if bathroomsStr != "" {
		houseForm.Bathrooms, err = strconv.ParseInt(bathroomsStr, 10, 64)
		if err != nil {
			return houseForm, fmt.Errorf("invalid bathrooms: %w", err), "Nombre de salles de bain invalide"
		}
	}

	// Parse floors
	floorsStr := r.FormValue("floors")
	if floorsStr != "" {
		houseForm.Floors, err = strconv.ParseInt(floorsStr, 10, 64)
		if err != nil {
			return houseForm, fmt.Errorf("invalid floors: %w", err), "Nombre d'étages invalide"
		}
	}

	// Parse construction year (optional)
	if yearStr := r.FormValue("construction_year"); yearStr != "" {
		houseForm.ConstructionYear, err = strconv.ParseInt(yearStr, 10, 64)
		if err != nil {
			return houseForm, fmt.Errorf("invalid construction year: %w", err), "Année de construction invalide"
		}
	}

	// Parse house type
	houseForm.HouseType = r.FormValue("house_type")

	// Parse land surface (optional)
	if landSurfaceStr := r.FormValue("land_surface"); landSurfaceStr != "" {
		houseForm.LandSurface, err = strconv.ParseInt(landSurfaceStr, 10, 64)
		if err != nil {
			return houseForm, fmt.Errorf("invalid land surface: %w", err), "Surface du terrain invalide"
		}
	}

	// Parse has garage (optional)
	if hasGarageStr := r.FormValue("has_garage"); hasGarageStr != "" {
		houseForm.HasGarage, err = strconv.ParseBool(hasGarageStr)
		if err != nil {
			return houseForm, fmt.Errorf("invalid has garage: %w", err), "Valeur invalide pour le garage"
		}
	}

	// Parse outdoor parking spaces (optional)
	if outdoorParkingStr := r.FormValue("outdoor_parking_spaces"); outdoorParkingStr != "" {
		houseForm.OutdoorParkingSpaces, err = strconv.ParseInt(outdoorParkingStr, 10, 64)
		if err != nil {
			return houseForm, fmt.Errorf("invalid outdoor parking spaces: %w", err), "Nombre de places de parking extérieures invalide"
		}
	}

	// Parse notes (optional)
	houseForm.Notes = r.FormValue("notes")

	// Get main photo
	houseForm.MainPhoto = r.FormValue("main_photo")

	return houseForm, nil, ""
}
