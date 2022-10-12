package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/es", handlerSpanish)

	http.HandleFunc("/en", handlerEnglish)

	http.HandleFunc("/pt", handlerPortuguese)

	address := ":8080"
	fmt.Println("server running ", address)
	http.ListenAndServe(address, nil)

}

func handlerSpanish(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "hola")
}

func handlerEnglish(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "hello")
}

func handlerPortuguese(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "ola")
}
