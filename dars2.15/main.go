package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	//db, err := postgres.ConnectDB()
	//if err != nil {
	//	panic(err)
	//}
	//
	//pl := postgres.PlayerRepo{db}
	//
	//server := handler.NewHandler(pl)
	//server.ListenAndServe()
	r := mux.NewRouter()

	r.HandleFunc("/", just).Methods(http.MethodGet)

	r2 := r.PathPrefix("/user").Subrouter()

	// /user/me
	r2.HandleFunc("/me", just2)

	http.ListenAndServe(":8080", r)
}

// REST FULL API

func just(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func just2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Dunyo"))
}
