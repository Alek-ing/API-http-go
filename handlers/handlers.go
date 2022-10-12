package handlers

import (
	"fmt"
	"net/http"
)

func HandlerSpanish(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "hola")
}

func HandlerEnglish(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "hello")
}

func HandlerPortuguese(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "ola")
}
