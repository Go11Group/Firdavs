package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectDB PostgreSQL ma'lumotlar bazasiga bog'lanish uchun funksiya.
// Bog'lanish parametrlari o'zgaruvchilar orqali aniqlanadi va bog'lanish amalga oshiriladi.
func ConnectDB() (*sql.DB, error) {
	// Bog'lanish parametrlari
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		dbname   = "imtihon_2"
		password = "123"
	)

	// Bog'lanish so'rovi tuzilishi
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)

	// PostgreSQL ga bog'lanish ochiladi
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	

	// Bog'lanishni sinovlaydi
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Ma'lumotlar bazasiga muvaffaqiyatli bog'lanish qaytariladi
	return db, nil
}
