package main

import (
	"database/sql"
	"fmt"

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

	// user := User{}
	// err = db.QueryRow(`select * from  user_car`).
	// 	Scan(&user.ID, &user.Name, &user.Age, &user.Gender)
	// if err != nil {
	// 	panic(err)
	// }
	//INSERT
	// _, err = db.Exec("insert into user_car(id, name, age, gender) values ($1, $2, $4, $3)",
	// 	uuid.NewString(), "Ibrohim", "m", 17)
	// if err != nil {
	// 	panic(err)
	// }
	// UPDATE
	// _, err = db.Exec("update user_car set name=$1  where  age=$2 ", "ali", 17)
	// if err != nil {
	// 	panic(err)
	// }
	// DELETE
	_, err = db.Exec("delete from user_car  where  age=$1 ",  17)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	fmt.Println("Successfully connected!")
}
