package models

import (
	"time"

	"github.com/willoma/recherche-maison/db"
)

// PublicationURL represents a publication URL for a house
type PublicationURL struct {
	ID              int64
	HouseID         int64
	URL             string
	PublicationDate time.Time
}

// FromDBPublicationURL converts a db.PublicationURL to a models.PublicationURL
func FromDBPublicationURL(dbPub db.PublicationURL) PublicationURL {
	return PublicationURL{
		ID:              dbPub.ID,
		HouseID:         dbPub.HouseID,
		URL:             dbPub.URL,
		PublicationDate: dbPub.PublicationDate,
	}
}

// ToDBPublicationURL converts a models.PublicationURL to a db.PublicationURL
func (p *PublicationURL) ToDBPublicationURL() db.PublicationURL {
	return db.PublicationURL{
		ID:              p.ID,
		HouseID:         p.HouseID,
		URL:             p.URL,
		PublicationDate: p.PublicationDate,
	}
}

// FromDBPublicationURLs converts a slice of db.PublicationURL to a slice of models.PublicationURL
func FromDBPublicationURLs(dbPubs []db.PublicationURL) []PublicationURL {
	pubs := make([]PublicationURL, len(dbPubs))
	for i, dbPub := range dbPubs {
		pubs[i] = FromDBPublicationURL(dbPub)
	}
	return pubs
}

// ToDBPublicationURLs converts a slice of models.PublicationURL to a slice of db.PublicationURL
func ToDBPublicationURLs(pubs []PublicationURL) []db.PublicationURL {
	dbPubs := make([]db.PublicationURL, len(pubs))
	for i, pub := range pubs {
		dbPubs[i] = pub.ToDBPublicationURL()
	}
	return dbPubs
}

// UpdatePublicationURLParams represents parameters for updating a publication URL
type UpdatePublicationURLParams struct {
	ID              int64
	URL             string
	PublicationDate time.Time
}

// ToUpdatePublicationURLParams converts models.UpdatePublicationURLParams to db.UpdatePublicationURLParams
func (p *UpdatePublicationURLParams) ToUpdatePublicationURLParams() db.UpdatePublicationURLParams {
	return db.UpdatePublicationURLParams{
		ID:              p.ID,
		URL:             p.URL,
		PublicationDate: p.PublicationDate,
	}
}
