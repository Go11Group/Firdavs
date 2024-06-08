package handler

import (
	"n11/Firdavs/dars2.14/storage/postgres"
	"net/http"
)

var fruites = map[string]string{"1": "apple", "54": "banana", "23": "lemon"}

type Handler struct {
	Book *postgres.BookRepo
}

func NewHandler(handler Handler) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /fruit/", juice)
	mux.HandleFunc("/concatinate-full-name", concatination)
	mux.HandleFunc("/book/", handler.book)

	return &http.Server{Handler: mux}
}

type Book struct {
	Name, Author, Publisher string
}
