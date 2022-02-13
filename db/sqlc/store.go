package db

import (
	"database/sql"
)

// Provides all functions to execute database queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// Creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
