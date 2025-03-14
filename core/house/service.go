package house

import (
	"context"
	"io"

	"github.com/willoma/recherche-maison/db"
)

// Service provides methods for managing houses
type Service struct {
	queries    *db.Queries
	uploadsDir string
}

// NewService creates a new house service
func NewService(queries *db.Queries, uploadsDir string) *Service {
	return &Service{
		queries:    queries,
		uploadsDir: uploadsDir,
	}
}

// GetHouse retrieves a house by ID
func (s *Service) GetHouse(ctx context.Context, id int64) (db.House, error) {
	// To be implemented
	return db.House{}, nil
}

// ListHouses retrieves all houses
func (s *Service) ListHouses(ctx context.Context) ([]db.House, error) {
	// To be implemented
	return nil, nil
}

// CreateHouse creates a new house
func (s *Service) CreateHouse(ctx context.Context, params db.CreateHouseParams) (db.House, error) {
	// To be implemented
	return db.House{}, nil
}

// UpdateHouse updates an existing house
func (s *Service) UpdateHouse(ctx context.Context, params db.UpdateHouseParams) (db.House, error) {
	// To be implemented
	return db.House{}, nil
}

// DeleteHouse deletes a house and its associated files
func (s *Service) DeleteHouse(ctx context.Context, id int64) error {
	// To be implemented
	return nil
}

// GetPublicationURLs retrieves all publication URLs for a house
func (s *Service) GetPublicationURLs(ctx context.Context, houseID int64) ([]db.PublicationURL, error) {
	// To be implemented
	return nil, nil
}

// CreatePublicationURL creates a new publication URL
func (s *Service) CreatePublicationURL(ctx context.Context, params db.CreatePublicationURLParams) (db.PublicationURL, error) {
	// To be implemented
	return db.PublicationURL{}, nil
}

// UpdatePublicationURL updates an existing publication URL
func (s *Service) UpdatePublicationURL(ctx context.Context, params db.UpdatePublicationURLParams) (db.PublicationURL, error) {
	// To be implemented
	return db.PublicationURL{}, nil
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
