package house

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/willoma/recherche-maison/config"
	"github.com/willoma/recherche-maison/db"
	"github.com/willoma/recherche-maison/models"
)

// Service provides house-related functionality
type Service struct {
	queries *db.Queries
	db      *sql.DB // Direct access to the database for transactions
}

// NewService creates a new house service
func NewService(queries *db.Queries, dbConn *sql.DB) *Service {
	return &Service{
		queries: queries,
		db:      dbConn,
	}
}

// GetHouse retrieves a house by ID
func (s *Service) GetHouse(ctx context.Context, id int64) (models.House, error) {
	dbHouse, err := s.queries.GetHouse(ctx, id)
	if err != nil {
		return models.House{}, fmt.Errorf("failed to get house: %w", err)
	}
	return models.FromDBHouse(dbHouse), nil
}

// ListHouses retrieves all houses
func (s *Service) ListHouses(ctx context.Context) ([]models.House, error) {
	dbHouses, err := s.queries.ListHouses(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list houses: %w", err)
	}

	houses := make([]models.House, len(dbHouses))
	for i, dbHouse := range dbHouses {
		houses[i] = models.FromDBHouse(dbHouse)
	}
	return houses, nil
}

// CreateHouse creates a new house
func (s *Service) CreateHouse(ctx context.Context, house models.House) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback()

	queries := s.queries.WithTx(tx)

	id, err := queries.CreateHouse(ctx, db.CreateHouseParams{
		Title:                house.Title,
		CityID:               house.CityID,
		Address:              house.Address,
		Price:                house.Price,
		Surface:              house.Surface,
		Rooms:                house.Rooms,
		Bedrooms:             house.Bedrooms,
		Bathrooms:            house.Bathrooms,
		Floors:               house.Floors,
		ConstructionYear:     house.ConstructionYear,
		HouseType:            house.HouseType,
		LandSurface:          house.LandSurface,
		HasGarage:            house.HasGarage,
		OutdoorParkingSpaces: house.OutdoorParkingSpaces,
		Notes:                house.Notes,
		MainPhoto:            house.MainPhoto,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to create house: %w", err)
	}

	// Create the uploads directories for this house
	photosDir := filepath.Join(config.UploadsDir, strconv.FormatInt(id, 10), "photos")
	if err := os.MkdirAll(photosDir, 0o755); err != nil {
		return 0, fmt.Errorf("failed to create photos directory: %w", err)
	}
	attachmentsDir := filepath.Join(config.UploadsDir, strconv.FormatInt(id, 10), "attachments")
	if err := os.MkdirAll(attachmentsDir, 0o755); err != nil {
		return 0, fmt.Errorf("failed to create attachments directory: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return id, nil
}

// UpdateHouse updates an existing house
func (s *Service) UpdateHouse(ctx context.Context, id int64, house models.House) error {
	// Update the house in the database
	if err := s.queries.UpdateHouse(ctx, db.UpdateHouseParams{
		ID:                   id,
		Title:                house.Title,
		CityID:               house.CityID,
		Address:              house.Address,
		Price:                house.Price,
		Surface:              house.Surface,
		Rooms:                house.Rooms,
		Bedrooms:             house.Bedrooms,
		Bathrooms:            house.Bathrooms,
		Floors:               house.Floors,
		ConstructionYear:     house.ConstructionYear,
		HouseType:            house.HouseType,
		LandSurface:          house.LandSurface,
		HasGarage:            house.HasGarage,
		OutdoorParkingSpaces: house.OutdoorParkingSpaces,
		MainPhoto:            house.MainPhoto,
		Notes:                house.Notes,
	}); err != nil {
		return fmt.Errorf("failed to update house: %w", err)
	}

	return nil
}

// DeleteHouse deletes a house and its associated files
func (s *Service) DeleteHouse(ctx context.Context, id int64) error {
	// Begin a transaction
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Create a new queries instance with the transaction
	qtx := s.queries.WithTx(tx)

	// Delete all publication URLs for this house
	if err := qtx.DeleteAllPublicationURLs(ctx, id); err != nil {
		return fmt.Errorf("failed to delete publication URLs: %w", err)
	}

	// Delete the house from the database
	if err := qtx.DeleteHouse(ctx, id); err != nil {
		return fmt.Errorf("failed to delete house: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Delete the uploads directory for this house
	uploadsDir := filepath.Join(config.UploadsDir, fmt.Sprintf("%d", id))
	if err := os.RemoveAll(uploadsDir); err != nil {
		slog.Error("Failed to delete uploads directory", "error", err, "path", uploadsDir)
		// Continue even if directory deletion fails, as the database records are already deleted
	}

	return nil
}

// GetPublicationURLs retrieves all publication URLs for a house
func (s *Service) GetPublicationURLs(ctx context.Context, houseID int64) ([]models.PublicationURL, error) {
	dbPublicationURLs, err := s.queries.GetPublicationURLs(ctx, houseID)
	if err != nil {
		return nil, fmt.Errorf("failed to get publication URLs: %w", err)
	}

	publicationURLs := make([]models.PublicationURL, len(dbPublicationURLs))
	for i, dbPub := range dbPublicationURLs {
		publicationURLs[i] = models.PublicationURL{
			ID:              dbPub.ID,
			HouseID:         dbPub.HouseID,
			URL:             dbPub.URL,
			PublicationDate: dbPub.PublicationDate,
		}
	}

	return publicationURLs, nil
}

// AddPublicationURL adds a new publication URL for a house
func (s *Service) AddPublicationURL(ctx context.Context, houseID int64, url string, publicationDate time.Time) error {
	params := db.CreatePublicationURLParams{
		HouseID:         houseID,
		URL:             url,
		PublicationDate: publicationDate,
	}

	if err := s.queries.CreatePublicationURL(ctx, params); err != nil {
		return fmt.Errorf("failed to add publication URL: %w", err)
	}

	return nil
}

// UpdatePublicationURL updates an existing publication URL
func (s *Service) UpdatePublicationURL(ctx context.Context, id int64, url string, publicationDate time.Time) error {
	params := db.UpdatePublicationURLParams{
		ID:              id,
		URL:             url,
		PublicationDate: publicationDate,
	}

	if err := s.queries.UpdatePublicationURL(ctx, params); err != nil {
		return fmt.Errorf("failed to update publication URL: %w", err)
	}

	return nil
}

// DeletePublicationURL deletes a publication URL
func (s *Service) DeletePublicationURL(ctx context.Context, id int64) error {
	if err := s.queries.DeletePublicationURL(ctx, id); err != nil {
		return fmt.Errorf("failed to delete publication URL: %w", err)
	}

	return nil
}

// GetPhotos retrieves all photos for a house
func (s *Service) GetPhotos(ctx context.Context, houseID int64) ([]string, error) {
	uploadsDir := filepath.Join(config.UploadsDir, fmt.Sprintf("%d", houseID))

	// Check if directory exists
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		return nil, nil
	}

	// Read directory entries
	entries, err := os.ReadDir(uploadsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read photos directory: %w", err)
	}

	// Filter for image files
	var photos []string
	for _, entry := range entries {
		if !entry.IsDir() && isImageFile(entry.Name()) {
			photos = append(photos, entry.Name())
		}
	}

	return photos, nil
}

// GetAttachments retrieves all attachments for a house
func (s *Service) GetAttachments(ctx context.Context, houseID int64) ([]string, error) {
	uploadsDir := filepath.Join(config.UploadsDir, fmt.Sprintf("%d", houseID))

	// Check if directory exists
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		return nil, nil
	}

	// Read directory entries
	entries, err := os.ReadDir(uploadsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read attachments directory: %w", err)
	}

	// Filter for non-image files
	var attachments []string
	for _, entry := range entries {
		if !entry.IsDir() && !isImageFile(entry.Name()) {
			attachments = append(attachments, entry.Name())
		}
	}

	return attachments, nil
}

// Helper function to check if a file is an image based on its extension
func isImageFile(filename string) bool {
	ext := filepath.Ext(filename)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp":
		return true
	default:
		return false
	}
}

// ParsePublicationDate parses a date string in the format "YYYY-MM-DD"
func ParsePublicationDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}
