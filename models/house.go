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

// ToCreateHouseParams converts a HouseForm to db.CreateHouseParams
func (h *House) ToCreateHouseParams() db.CreateHouseParams {
	return db.CreateHouseParams{
		Title:                h.Title,
		CityID:               h.CityID,
		Address:              h.Address,
		Price:                h.Price,
		Surface:              h.Surface,
		Rooms:                h.Rooms,
		Bedrooms:             h.Bedrooms,
		Bathrooms:            h.Bathrooms,
		Floors:               h.Floors,
		ConstructionYear:     h.ConstructionYear,
		HouseType:            h.HouseType,
		LandSurface:          h.LandSurface,
		HasGarage:            h.HasGarage,
		OutdoorParkingSpaces: h.OutdoorParkingSpaces,
		Notes:                h.Notes,
		MainPhoto:            h.MainPhoto,
	}
}

// ToUpdateHouseParams converts a HouseForm to db.UpdateHouseParams
func (h *House) ToUpdateHouseParams() db.UpdateHouseParams {
	return db.UpdateHouseParams{
		ID:                   h.ID,
		Title:                h.Title,
		CityID:               h.CityID,
		Address:              h.Address,
		Price:                h.Price,
		Surface:              h.Surface,
		Rooms:                h.Rooms,
		Bedrooms:             h.Bedrooms,
		Bathrooms:            h.Bathrooms,
		Floors:               h.Floors,
		ConstructionYear:     h.ConstructionYear,
		HouseType:            h.HouseType,
		LandSurface:          h.LandSurface,
		HasGarage:            h.HasGarage,
		OutdoorParkingSpaces: h.OutdoorParkingSpaces,
		MainPhoto:            h.MainPhoto,
		Notes:                h.Notes,
	}
}
