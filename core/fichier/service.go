package fichier

import (
	"io"
)

// Service provides methods for managing files
type Service struct {
	uploadsDir string
}

// NewService creates a new file service
func NewService(uploadsDir string) *Service {
	return &Service{
		uploadsDir: uploadsDir,
	}
}

// EnsureHouseDir ensures that the directory for a house exists
func (s *Service) EnsureHouseDir(houseID int64) error {
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

// DeleteHouseFiles deletes all files for a house
func (s *Service) DeleteHouseFiles(houseID int64) error {
	// To be implemented
	return nil
}
