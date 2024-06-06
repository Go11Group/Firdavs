package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var fruites = map[string]string{"1": "apple", "54": "banana", "23": "lemon"}

func main() {
	http.HandleFunc("GET /fruit/", juice)
	http.HandleFunc("GET /concatinate-full-name", concatination)
	http.HandleFunc("POST /book", book)

	err := http.ListenAndServe(":8070", nil)
	if err != nil {
		panic(err)
	}
}

func juice(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.Write([]byte(fruites[strings.TrimPrefix(r.URL.Path, "/fruit/")]))
}

func concatination(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Query()["name"][0] + r.URL.Query()["lastname"][0]))
	fmt.Println(r.URL.Query()["name"][0] + r.URL.Query()["lastname"][0])
}

func book(w http.ResponseWriter, r *http.Request) {
	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	panic(err)
	//}
	var b Book
	err := json.NewDecoder(r.Body).Decode(&b)
	//err = json.Unmarshal(body, &b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while Decode, err: %s", err.Error())))
		return
	}

	fmt.Println(b)
	w.Write([]byte("SUCCESS"))
}

type Book struct {
	category, sys string
}
