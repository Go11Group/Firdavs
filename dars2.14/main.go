package main

import (
	"n11/Firdavs/dars2.14/storage/postgres"
	"n11/Firdavs/dars2.14/handler"

)

func main() {
	db, err := postgres.User_ConnectDB()
	if err != nil {
		panic(err)
	}
	// usere := postgres.User_NewRepo(db)
	product := postgres.Product_NewRepo(db)

	server := handler.NewHandler(handler.Handler{Productr:  product})

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}