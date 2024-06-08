package handler

import (
	"n11/Firdavs/dars2.14/storage/postgres"
	"net/http"
)

type Handler struct {
	Userr    *postgres.UserRepo
	Productr *postgres.ProductRepo
}

func NewHandler(handler Handler) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/", handler.User)
	mux.HandleFunc("/userdelete/", handler.Userdelete)
	mux.HandleFunc("/usercreate/", handler.UserCreate)
	mux.HandleFunc("/userupdate/", handler.UserUpdate)
	mux.HandleFunc("/products/", handler.Product)
	mux.HandleFunc("/productcreate/", handler.ProductCreate)
	mux.HandleFunc("/productdelete/", handler.Productdelete)
	mux.HandleFunc("/productupdate/", handler.ProductUpdate)
	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

type Userr struct {
	Username, Email, Password string
}

type Productr struct {
	Name           string
	Description    string
	Price          float32
	Stock_quantity int
}
