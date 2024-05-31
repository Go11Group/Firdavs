package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("postgres://postgres:pass@localhost:5432/go11?sslmode=disable"))

	if err != nil {
		return nil, err
	}

	return db, nil
}
