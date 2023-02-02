package main

import (
	"fmt"
	"log"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Welcome to Homepage")
	if err != nil {
		log.Println("Error writing response", err)
	}
	log.Println("Wrote", n, "bytes")
}

func About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Welcome to About Page and Total is %d", addValue(10, 20))
	if err != nil {
		log.Println("Error writing response", err)
	}
	log.Println("Wrote", n, "bytes")
}

func addValue(a int, b int) int {
	return a + b
}

func main() {
	const portNumber = ":8080"
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/about", About)

	log.Println("Starting server on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
