package main

import (
	"database/sql"
	"fmt"
	"github.com/go-faker/faker/v4"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:pass@localhost:5432/go11?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for i := 0; i < 100000; i++ {
		_, err = db.Exec("insert into product (name, category, cost) values ($1, $2, $3)",
			faker.FirstName(), faker.LastName(), 4234)
		if err != nil {
			fmt.Println(err)
		}
		if i%1000 == 0 {
			fmt.Println(i)
		}
	}
}
