package handler

import (
	"fmt"
	"net/http"
	"strings"
)

func concatination(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Query()["name"][0] + r.URL.Query()["lastname"][0]))
	fmt.Println(r.URL.Query()["age"][0] + r.URL.Query()["age"][1])
}

func juice(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.Write([]byte(fruites[strings.TrimPrefix(r.URL.Path, "/fruit/")]))
}
