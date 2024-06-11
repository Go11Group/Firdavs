package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func connect() (*sql.DB, error) {
	dsn := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}


