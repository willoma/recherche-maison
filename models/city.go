package models

import (
	"github.com/willoma/recherche-maison/db"
)

// City represents a city
type City struct {
	ID     int64
	Name   string
	IsUsed bool
}

// FromDBCity converts a db.City to a models.City
func FromDBCity(dbCity db.City) City {
	return City{
		ID:     dbCity.ID,
		Name:   dbCity.Name,
		IsUsed: dbCity.IsUsed,
	}
}

// FromDBCities converts a slice of db.City to a slice of models.City
func FromDBCities(dbCities []db.City) []City {
	cities := make([]City, len(dbCities))
	for i, dbCity := range dbCities {
		cities[i] = FromDBCity(dbCity)
	}
	return cities
}
