package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Welcome to new server")
	})

	http.HandleFunc("/2", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Welcome to new server2")
	})

	http.HandleFunc("/3", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Welcome to new server3")
	})

	address := ":8080"
	fmt.Println("server running ", address)
	http.ListenAndServe(address, nil)

}
