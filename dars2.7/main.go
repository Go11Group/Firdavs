package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "go11"
	password = "pass"
)

type User struct {
	ID     string
	Name   string
	Age    int
	Gender string
	Course string
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

	user := User{}
	err = db.QueryRow(`select s.id, s.name, age, gender, c.name from student s
    	left join course c on c.id = s.course_id offset 2`).
		Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
	if err != nil {
		panic(err)
	}
	//
	_, err = db.Exec("insert into student(id, name, age, gender, course_id) values ($1, $2, $4, $3, $2)",
		uuid.NewString(), "Ibrohim", "m", 17, "0a3164d4-1511-444a-8b58-e722474fbffb")
	if err != nil {
		panic(err)
	}
	db.Exec("update student set name=$1, age=$2 ", "", 34)

	fmt.Println(user)

	fmt.Println("Successfully connected!")
}
