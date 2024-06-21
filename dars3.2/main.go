package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// ! GET ALL
	resp, err := http.Get("http://localhost:8080/user")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(" ")
	
	// ! GET
	resp, err = http.Get("http://localhost:8080/user/8832c8fc-a020-4b95-9b55-abc66a57c538")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body1, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body1))

	// ! POST
	_, err = http.Post("http://localhost:8080/user",
		"json",
		bytes.NewBuffer([]byte(`{"name":"sdaf","email":"asdfdas","brithday":"2015-03-04","password":"pass"}`)))
	if err != nil {
		panic(err)
	}

	// ! PUT
	client := http.Client{}
	req, err := http.NewRequest("PUT",
		"http://localhost:8080/user/0cf71016-13a9-4ce1-9a88-5bc385308f57",
		bytes.NewBuffer([]byte(`{"name":"sdaf","email":"asdfdas"}`)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/text")
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	// ! DELETE
	client = http.Client{}
	req, err = http.NewRequest("DELETE",
		"http://localhost:8080/user/0cf71016-13a9-4ce1-9a88-5bc385308f57", nil)
	if err != nil {
		panic(err)
	}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}
