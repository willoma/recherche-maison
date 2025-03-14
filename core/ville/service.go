package ville

import (
	"context"

	"github.com/willoma/recherche-maison/db"
)

// Service provides methods for managing cities
type Service struct {
	queries *db.Queries
}

// NewService creates a new city service
func NewService(queries *db.Queries) *Service {
	return &Service{
		queries: queries,
	}
}

// GetCity retrieves a city by ID
func (s *Service) GetCity(ctx context.Context, id int64) (db.City, error) {
	// To be implemented
	return db.City{}, nil
}

// ListCities retrieves all cities
func (s *Service) ListCities(ctx context.Context) ([]db.City, error) {
	// To be implemented
	return nil, nil
}

// CreateCity creates a new city
func (s *Service) CreateCity(ctx context.Context, name string) (db.City, error) {
	// To be implemented
	return db.City{}, nil
}

// UpdateCity updates an existing city
func (s *Service) UpdateCity(ctx context.Context, id int64, name string) (db.City, error) {
	// To be implemented
	return db.City{}, nil
}

// DeleteCity deletes a city if it's not used by any house
func (s *Service) DeleteCity(ctx context.Context, id int64) error {
	// To be implemented
	return nil
}

// IsCityUsedByHouses checks if a city is used by any house
func (s *Service) IsCityUsedByHouses(ctx context.Context, id int64) (bool, error) {
	// To be implemented
	return false, nil
}
