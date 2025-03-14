package models

import (
	"time"

	"github.com/willoma/recherche-maison/db"
)

// House represents a house with all its details
type House struct {
	ID                   int64
	Title                string
	CityID               int64
	CityName             string
	Address              string
	Price                int64
	Surface              int64
	Rooms                int64
	Bedrooms             int64
	Bathrooms            int64
	Floors               int64
	ConstructionYear     int64
	HouseType            string
	LandSurface          int64
	HasGarage            bool
	OutdoorParkingSpaces int64
	MainPhoto            string
	Notes                string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

// FromDBHouse converts a db.House to a models.House
func FromDBHouse(dbHouse db.House) House {
	return House{
		ID:                   dbHouse.ID,
		Title:                dbHouse.Title,
		CityID:               dbHouse.CityID,
		CityName:             dbHouse.CityName,
		Address:              dbHouse.Address,
		Price:                dbHouse.Price,
		Surface:              dbHouse.Surface,
		Rooms:                dbHouse.Rooms,
		Bedrooms:             dbHouse.Bedrooms,
		Bathrooms:            dbHouse.Bathrooms,
		Floors:               dbHouse.Floors,
		ConstructionYear:     dbHouse.ConstructionYear,
		HouseType:            dbHouse.HouseType,
		LandSurface:          dbHouse.LandSurface,
		HasGarage:            dbHouse.HasGarage,
		OutdoorParkingSpaces: dbHouse.OutdoorParkingSpaces,
		MainPhoto:            dbHouse.MainPhoto,
		Notes:                dbHouse.Notes,
		CreatedAt:            dbHouse.CreatedAt,
		UpdatedAt:            dbHouse.UpdatedAt,
	}
}
