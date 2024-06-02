package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	CreateUser(db, "username", "email@example.com", "password")
}

func CreateUser(db *sql.DB, username, email, password string) {
	query := `
    INSERT INTO users (username, email, password)
    VALUES ($1, $2, $3)
    RETURNING id`

	var userID int
	err := db.QueryRow(query, username, email, password).Scan(&userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New user ID is %d\n", userID)
}

func CreateProduct(db *sql.DB, name, description string, price float64, stockQuantity int) {
	query := `
    INSERT INTO products (name, description, price, stock_quantity)
    VALUES ($1, $2, $3, $4)
    RETURNING id`

	var productID int
	err := db.QueryRow(query, name, description, price, stockQuantity).Scan(&productID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New product ID is %d\n", productID)
}

func GetUser(db *sql.DB, userID int) {
	query := `
    SELECT id, username, email, password
    FROM users
    WHERE id = $1`

	var id int
	var username, email, password string
	err := db.QueryRow(query, userID).Scan(&id, &username, &email, &password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User: %d, %s, %s, %s\n", id, username, email, password)
}

func GetProduct(db *sql.DB, productID int) {
	query := `
    SELECT id, name, description, price, stock_quantity
    FROM products
    WHERE id = $1`

	var id int
	var name, description string
	var price float64
	var stockQuantity int
	err := db.QueryRow(query, productID).Scan(&id, &name, &description, &price, &stockQuantity)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Product: %d, %s, %s, %.2f, %d\n", id, name, description, price, stockQuantity)
}

func UpdateUser(db *sql.DB, userID int, username, email, password string) {
	query := `
    UPDATE users
    SET username = $1, email = $2, password = $3
    WHERE id = $4`

	_, err := db.Exec(query, username, email, password, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User updated successfully")
}

func UpdateProduct(db *sql.DB, productID int, name, description string, price float64, stockQuantity int) {
	query := `
    UPDATE products
    SET name = $1, description = $2, price = $3, stock_quantity = $4
    WHERE id = $5`

	_, err := db.Exec(query, name, description, price, stockQuantity, productID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Product updated successfully")
}

func DeleteUser(db *sql.DB, userID int) {
	query := `
    DELETE FROM users
    WHERE id = $1`

	_, err := db.Exec(query, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User deleted successfully")
}

func DeleteProduct(db *sql.DB, productID int) {
	query := `
    DELETE FROM products
    WHERE id = $1`

	_, err := db.Exec(query, productID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Product deleted successfully")
}

func CreateUserWithProducts(db *sql.DB, username, email, password string, products []Product) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	queryUser := `
    INSERT INTO users (username, email, password)
    VALUES ($1, $2, $3)
    RETURNING id`

	var userID int
	err = tx.QueryRow(queryUser, username, email, password).Scan(&userID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	for _, product := range products {
		queryProduct := `
        INSERT INTO products (name, description, price, stock_quantity)
        VALUES ($1, $2, $3, $4)`

		_, err = tx.Exec(queryProduct, product.Name, product.Description, product.Price, product.StockQuantity)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User and products created successfully")
}

type Product struct {
	Name          string
	Description   string
	Price         float64
	StockQuantity int
}
