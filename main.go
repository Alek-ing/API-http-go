package main

import (
	"fmt"
	"net/http"

	h "github.com/Alek-ing/API-http-go/handlers"
)

func main() {
	http.HandleFunc("/es", h.HandlerSpanish)

	http.HandleFunc("/en", h.HandlerEnglish)

	http.HandleFunc("/pt", h.HandlerPortuguese)

	address := ":8080"
	fmt.Println("server running ", address)
	http.ListenAndServe(address, nil)

}
