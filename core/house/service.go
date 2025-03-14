package house

import (
	"context"
	"io"
	"time"

	"github.com/willoma/recherche-maison/db"
	"github.com/willoma/recherche-maison/models"
)

// Service provides methods for managing houses
type Service struct {
	queries *db.Queries
}

// NewService creates a new house service
func NewService(queries *db.Queries) *Service {
	return &Service{
		queries: queries,
	}
}

// GetHouse retrieves a house by ID
func (s *Service) GetHouse(ctx context.Context, id int64) (models.House, error) {
	// To be implemented
	dbHouse := db.House{}
	return models.FromDBHouse(dbHouse), nil
}

// ListHouses retrieves all houses
func (s *Service) ListHouses(ctx context.Context) ([]models.House, error) {
	// To be implemented
	var houses []models.House
	return houses, nil
}

// CreateHouse creates a new house
func (s *Service) CreateHouse(ctx context.Context, house models.House) (int64, error) {
	// Convert to db model and create
	dbParams := house.ToCreateHouseParams()
	// To be implemented - using dbParams in the actual implementation
	_ = dbParams // Temporary to avoid unused variable warning
	return 0, nil
}

// UpdateHouse updates an existing house
func (s *Service) UpdateHouse(ctx context.Context, id int64, house models.House) error {
	// Convert to db model and update
	house.ID = id
	dbParams := house.ToUpdateHouseParams()
	// To be implemented - using dbParams in the actual implementation
	_ = dbParams // Temporary to avoid unused variable warning
	return nil
}

// DeleteHouse deletes a house and its associated files
func (s *Service) DeleteHouse(ctx context.Context, id int64) error {
	// To be implemented
	return nil
}

// GetPublicationURLs retrieves all publication URLs for a house
func (s *Service) GetPublicationURLs(ctx context.Context, houseID int64) ([]models.PublicationURL, error) {
	// To be implemented
	var publicationURLs []models.PublicationURL
	return publicationURLs, nil
}

// AddPublicationURL adds a new publication URL for a house
func (s *Service) AddPublicationURL(ctx context.Context, houseID int64, url string, publicationDate time.Time) error {
	// To be implemented
	return nil
}

// UpdatePublicationURL updates an existing publication URL
func (s *Service) UpdatePublicationURL(ctx context.Context, params models.UpdatePublicationURLParams) error {
	// Convert to db model and update
	dbParams := params.ToUpdatePublicationURLParams()
	// To be implemented - using dbParams in the actual implementation
	_ = dbParams // Temporary to avoid unused variable warning
	return nil
}

// DeletePublicationURL deletes a publication URL
func (s *Service) DeletePublicationURL(ctx context.Context, id int64) error {
	// To be implemented
	return nil
}

// SavePhoto saves a photo file for a house
func (s *Service) SavePhoto(houseID int64, filename string, file io.Reader) (string, error) {
	// To be implemented
	return "", nil
}

// GetPhotos retrieves all photos for a house
func (s *Service) GetPhotos(houseID int64) ([]string, error) {
	// To be implemented
	return nil, nil
}

// DeletePhoto deletes a photo file
func (s *Service) DeletePhoto(houseID int64, filename string) error {
	// To be implemented
	return nil
}

// SaveAttachment saves an attachment file for a house
func (s *Service) SaveAttachment(houseID int64, originalFilename string, file io.Reader) (string, error) {
	// To be implemented
	return "", nil
}

// GetAttachments retrieves all attachments for a house
func (s *Service) GetAttachments(houseID int64) ([]string, error) {
	// To be implemented
	return nil, nil
}

// DeleteAttachment deletes an attachment file
func (s *Service) DeleteAttachment(houseID int64, filename string) error {
	// To be implemented
	return nil
}

// ParsePublicationDate parses a date string in the format "YYYY-MM-DD"
func ParsePublicationDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}
