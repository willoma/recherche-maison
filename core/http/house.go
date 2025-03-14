package http

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/willoma/recherche-maison/config"
	"github.com/willoma/recherche-maison/core/house"
	"github.com/willoma/recherche-maison/db"
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
	houseParams, err, errMsg := parseHouseForm(r)
	if err != nil {
		slog.Error("Failed to parse house form", "error", err)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	// Set MainPhoto to empty as it will be set later
	houseParams.MainPhoto = sql.NullString{}

	// Create the house and get its ID
	houseID, err := s.houseService.CreateHouse(r.Context(), houseParams)
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
	photos, err := s.houseService.GetPhotos(id)
	if err != nil {
		slog.Error("Failed to get photos", "house_id", id, "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Get attachments
	attachments, err := s.houseService.GetAttachments(id)
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
	// Get house ID from URL
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		slog.Error("Invalid house ID", "id", idStr, "error", err)
		http.Error(w, "Identifiant de maison invalide", http.StatusBadRequest)
		return
	}

	// Parse form data
	houseParams, err, errMsg := parseHouseForm(r)
	if err != nil {
		slog.Error("Failed to parse house form", "error", err)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	// Get the current main photo
	houseData, err := s.houseService.GetHouse(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get house", "id", id, "error", err)
		http.Error(w, "Maison introuvable", http.StatusNotFound)
		return
	}
	mainPhoto := houseData.MainPhoto.String

	// Create house update parameters
	updateParams := db.UpdateHouseParams{
		ID:                   id,
		Title:                houseParams.Title,
		CityID:               houseParams.CityID,
		Address:              houseParams.Address,
		Price:                houseParams.Price,
		Surface:              houseParams.Surface,
		Rooms:                houseParams.Rooms,
		Bedrooms:             houseParams.Bedrooms,
		Bathrooms:            houseParams.Bathrooms,
		Floors:               houseParams.Floors,
		ConstructionYear:     houseParams.ConstructionYear,
		HouseType:            houseParams.HouseType,
		LandSurface:          houseParams.LandSurface,
		HasGarage:            houseParams.HasGarage,
		OutdoorParkingSpaces: houseParams.OutdoorParkingSpaces,
		MainPhoto:            sql.NullString{String: mainPhoto, Valid: mainPhoto != ""},
		Notes:                houseParams.Notes,
	}

	// Update the house in the database
	err = s.houseService.UpdateHouse(r.Context(), updateParams)
	if err != nil {
		slog.Error("Failed to update house", "id", id, "error", err)
		http.Error(w, "Erreur lors de la mise à jour de la maison", http.StatusInternalServerError)
		return
	}

	// Handle publication URLs
	urls := r.Form["pub_url[]"]
	dates := r.Form["pub_date[]"]

	if len(urls) == 0 {
		slog.Error("At least one publication URL is required")
		http.Error(w, "Au moins une URL de publication est obligatoire", http.StatusBadRequest)
		return
	}

	// First, get existing publication URLs
	existingURLs, err := s.houseService.GetPublicationURLs(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get existing publication URLs", "house_id", id, "error", err)
		// Continue anyway, as we want to add the new URLs
	}

	// Delete existing publication URLs
	for _, url := range existingURLs {
		err = s.houseService.DeletePublicationURL(r.Context(), url.ID)
		if err != nil {
			slog.Error("Failed to delete publication URL", "id", url.ID, "error", err)
			// Continue anyway
		}
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
		err = s.houseService.AddPublicationURL(r.Context(), id, url, date)
		if err != nil {
			slog.Error("Failed to add publication URL", "house_id", id, "url", url, "error", err)
		}
	}

	// Handle photo uploads
	// TODO: Implement photo uploads

	// Redirect to house details page
	http.Redirect(w, r, "/maison/"+strconv.FormatInt(id, 10), http.StatusSeeOther)
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
	photos, err := s.houseService.GetPhotos(id)
	if err != nil {
		slog.Error("Failed to get photos", "house_id", id, "error", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Get attachments
	attachments, err := s.houseService.GetAttachments(id)
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
func parseHouseForm(r *http.Request) (db.CreateHouseParams, error, string) {
	var houseParams db.CreateHouseParams
	var err error

	// Parse multipart form data (for file uploads)
	if err := r.ParseMultipartForm(config.MaxUploadSize); err != nil {
		return houseParams, err, "Erreur lors de la soumission du formulaire"
	}

	// Get required fields
	houseParams.Title = r.FormValue("title")
	if houseParams.Title == "" {
		return houseParams, fmt.Errorf("title is required"), "Le titre est obligatoire"
	}

	cityIDStr := r.FormValue("city_id")
	if cityIDStr == "" {
		return houseParams, fmt.Errorf("city ID is required"), "La ville est obligatoire"
	}

	houseParams.CityID, err = strconv.ParseInt(cityIDStr, 10, 64)
	if err != nil {
		return houseParams, fmt.Errorf("invalid city ID: %w", err), "Identifiant de ville invalide"
	}

	// Parse address (optional)
	if addressStr := r.FormValue("address"); addressStr != "" {
		houseParams.Address = sql.NullString{String: addressStr, Valid: true}
	}

	// Parse price
	priceStr := r.FormValue("price")
	if priceStr != "" {
		houseParams.Price, err = strconv.ParseInt(priceStr, 10, 64)
		if err != nil {
			return houseParams, fmt.Errorf("invalid price: %w", err), "Prix invalide"
		}
	}

	// Parse surface
	surfaceStr := r.FormValue("surface")
	if surfaceStr != "" {
		houseParams.Surface, err = strconv.ParseInt(surfaceStr, 10, 64)
		if err != nil {
			return houseParams, fmt.Errorf("invalid surface: %w", err), "Surface invalide"
		}
	}

	// Parse rooms
	roomsStr := r.FormValue("rooms")
	if roomsStr != "" {
		houseParams.Rooms, err = strconv.ParseInt(roomsStr, 10, 64)
		if err != nil {
			return houseParams, fmt.Errorf("invalid rooms: %w", err), "Nombre de pièces invalide"
		}
	}

	// Parse bedrooms
	bedroomsStr := r.FormValue("bedrooms")
	if bedroomsStr != "" {
		houseParams.Bedrooms, err = strconv.ParseInt(bedroomsStr, 10, 64)
		if err != nil {
			return houseParams, fmt.Errorf("invalid bedrooms: %w", err), "Nombre de chambres invalide"
		}
	}

	// Parse bathrooms
	bathroomsStr := r.FormValue("bathrooms")
	if bathroomsStr != "" {
		houseParams.Bathrooms, err = strconv.ParseInt(bathroomsStr, 10, 64)
		if err != nil {
			return houseParams, fmt.Errorf("invalid bathrooms: %w", err), "Nombre de salles de bain invalide"
		}
	}

	// Parse floors
	floorsStr := r.FormValue("floors")
	if floorsStr != "" {
		houseParams.Floors, err = strconv.ParseInt(floorsStr, 10, 64)
		if err != nil {
			return houseParams, fmt.Errorf("invalid floors: %w", err), "Nombre d'étages invalide"
		}
	}

	// Parse construction year (optional)
	if yearStr := r.FormValue("construction_year"); yearStr != "" {
		year, err := strconv.ParseInt(yearStr, 10, 64)
		if err != nil {
			return houseParams, fmt.Errorf("invalid construction year: %w", err), "Année de construction invalide"
		}
		houseParams.ConstructionYear = sql.NullInt64{Int64: year, Valid: true}
	}

	// Parse house type
	houseParams.HouseType = r.FormValue("house_type")

	// Parse land surface (optional)
	if landSurfaceStr := r.FormValue("land_surface"); landSurfaceStr != "" {
		ls, err := strconv.ParseInt(landSurfaceStr, 10, 64)
		if err != nil {
			return houseParams, fmt.Errorf("invalid land surface: %w", err), "Surface du terrain invalide"
		}
		houseParams.LandSurface = sql.NullInt64{Int64: ls, Valid: true}
	}

	// Parse has garage (optional)
	if hasGarageStr := r.FormValue("has_garage"); hasGarageStr != "" {
		hasGarageBool, err := strconv.ParseBool(hasGarageStr)
		if err != nil {
			return houseParams, fmt.Errorf("invalid has garage: %w", err), "Valeur invalide pour le garage"
		}
		houseParams.HasGarage = sql.NullBool{Bool: hasGarageBool, Valid: true}
	}

	// Parse outdoor parking spaces (optional)
	if outdoorParkingStr := r.FormValue("outdoor_parking_spaces"); outdoorParkingStr != "" {
		ops, err := strconv.ParseInt(outdoorParkingStr, 10, 64)
		if err != nil {
			return houseParams, fmt.Errorf("invalid outdoor parking spaces: %w", err), "Nombre de places de parking extérieures invalide"
		}
		houseParams.OutdoorParkingSpaces = sql.NullInt64{Int64: ops, Valid: true}
	}

	// Parse notes (optional)
	if notesStr := r.FormValue("notes"); notesStr != "" {
		houseParams.Notes = sql.NullString{String: notesStr, Valid: true}
	}

	return houseParams, nil, ""
}
