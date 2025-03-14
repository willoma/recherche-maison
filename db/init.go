package db

import (
	"database/sql"
	_ "embed"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schema string

// Init initializes the database connection and creates tables if they don't exist
func Init(dbPath, dbOptions string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", "file:"+dbPath+"?"+dbOptions)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}
