package main

import (
	"fmt"
	"net/http"
	"strings"
)


var pc = map[string]string{"brend":"Asus", "model":"Vivobook", "sys": "ubuntu"} 

func main() {
	http.HandleFunc("GET /api/mypc", info),
	http.HandleFunc("GET /api/yourpc", params),
	http.HandleFunc("GET /api/mypc", body)
	
	err := http.ListenAndServe(":8080",nil )
	if err != nil {
		panic(err)
	}
}

func info(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.URL.Path)
	w.Write([]byte(pc[strings.TrimPrefix((r.URL.Path, "api/mypc"))]))
}