package main

import (
	"n11/Firdavs/dars2.14/handler"
	"n11/Firdavs/dars2.14/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	book := postgres.NewBookRepo(db)

	server := handler.NewHandler(handler.Handler{Book: book})

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
