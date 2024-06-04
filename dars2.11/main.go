package main

import (
	"database/sql"
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:123@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for i := 0; i < 900000; i++ {
		_, err = db.Exec("insert into product (id ,name, category, cost) values ($1, $2, $3,$4)",
			uuid.NewString(), faker.FirstName(), faker.LastName(), 4234)
		if err != nil {
			fmt.Println(err)
		}
		if i%1000 == 0 {
			fmt.Println(i)
		}
	}
}
