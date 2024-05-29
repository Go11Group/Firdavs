package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "postgres"
	password = "123"
)

type User struct {
	ID     string
	Name   string
	Age    int
	Gender string
}

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Hello")
		panic(err)
	}

	// INSERT
	// _, err = db.Exec(`INSERT INTO customers (name, age) VALUES ($1, $2)`, "Alice", 30)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = db.Exec(`INSERT INTO orders (customer_id, product, amount) VALUES ($1, $2, $3)`, 1, "Laptop", 1200.50)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Data inserted successfully!")

	// // UPDATE
	// _, err = db.Exec(`UPDATE customers SET age = $1 WHERE name = $2`, 31, "Alice")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Data updated successfully!")

	// DELETE
	// _, err = db.Exec(`DELETE FROM customers WHERE name = $1`, "Alice")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Data deleted successfully!")

	// Query
	rows, err := db.Query(`SELECT customers.name, orders.product, orders.amount 
	FROM customers 
	JOIN orders ON customers.id = orders.customer_id 
	WHERE customers.id = $1`, 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, product string
		var amount float64
		err := rows.Scan(&name, &product, &amount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Customer: %s, Product: %s, Amount: %.2f\n", name, product, amount)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Jadval birlashmasi bilan QueryRow
	// var product string
	// var amount float64
	// err = db.QueryRow(`SELECT orders.product, orders.amount 
	// FROM customers 
	// JOIN orders ON customers.id = orders.customer_id 
	// WHERE customers.name = $1`, "Alice").Scan(&product, &amount)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Product: %s, Amount: %.2f\n", product, amount)
}
