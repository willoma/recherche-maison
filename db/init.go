package db

import (
	"database/sql"
	_ "embed"

	"github.com/willoma/recherche-maison/config"
	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schema string

// Init initializes the database connection and creates tables if they don't exist
func Init() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "file:"+config.DBPath+"?"+config.DBOptions)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}
