package ville

import (
	"context"
	"log/slog"

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
	city, err := s.queries.GetCity(ctx, id)
	if err != nil {
		slog.Error("Failed to get city", "id", id, "error", err)
		return db.City{}, err
	}
	return city, nil
}

// ListCities retrieves all cities
func (s *Service) ListCities(ctx context.Context) ([]db.City, error) {
	cities, err := s.queries.ListCities(ctx)
	if err != nil {
		slog.Error("Failed to list cities", "error", err)
		return nil, err
	}
	return cities, nil
}

// CreateCity creates a new city
func (s *Service) CreateCity(ctx context.Context, name string) error {
	err := s.queries.CreateCity(ctx, name)
	if err != nil {
		slog.Error("Failed to create city", "name", name, "error", err)
		return err
	}
	return nil
}

// UpdateCity updates an existing city
func (s *Service) UpdateCity(ctx context.Context, id int64, name string) error {
	err := s.queries.UpdateCity(ctx, name, id)
	if err != nil {
		slog.Error("Failed to update city", "id", id, "name", name, "error", err)
		return err
	}
	return nil
}

// DeleteCity deletes a city if it's not used by any house
func (s *Service) DeleteCity(ctx context.Context, id int64) error {
	if err := s.queries.DeleteCity(ctx, id); err != nil {
		slog.Error("Failed to delete city", "id", id, "error", err)
		return err
	}

	return nil
}
