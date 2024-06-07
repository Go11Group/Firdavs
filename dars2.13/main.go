package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("GET /hello", hello)
	http.HandleFunc("GET /hi", hi)
	http.HandleFunc("GET /go", goland)
	http.HandleFunc("GET /python", python)

}




func hello(w http.ResponseWriter, r *http.Request) {
	//if r.Method != http.MethodGet {
	//	w.Write([]byte("This method is not GET"))
	//	return
	//}
	fmt.Println(r.URL)
	fmt.Println(r.Host)
	fmt.Println(r.Method)

	n, err := w.Write([]byte("OZOOOOOD ......"))
	if err != nil {
		fmt.Println(err, n)
	}
}


func hi(w http.ResponseWriter, r *http.Request) {
	//if r.Method != http.MethodGet {
	//	w.Write([]byte("This method is not GET"))
	//	return
	//}
	fmt.Println(r.URL)
	fmt.Println(r.Host)
	fmt.Println(r.Method)

	n, err := w.Write([]byte("Hi good night"))
	if err != nil {
		fmt.Println(err, n)
	}
}

func goland(w http.ResponseWriter, r *http.Request) {
	//if r.Method != http.MethodGet {
	//	w.Write([]byte("This method is not GET"))
	//	return
	//}
	fmt.Println(r.URL)
	fmt.Println(r.Host)
	fmt.Println(r.Method)

	n, err := w.Write([]byte("Welcome to goland"))
	if err != nil {
		fmt.Println(err, n)
	}
}

func python(w http.ResponseWriter, r *http.Request) {
	//if r.Method != http.MethodGet {
	//	w.Write([]byte("This method is not GET"))
	//	return
	//}
	fmt.Println(r.URL)
	fmt.Println(r.Host)
	fmt.Println(r.Method)

	n, err := w.Write([]byte("Welcome to pythone"))
	if err != nil {
		fmt.Println(err, n)
	}
}