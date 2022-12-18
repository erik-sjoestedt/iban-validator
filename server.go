package main

import (
	"fmt"
	"iban-validator/iban"
	"log"
	"net/http"
)

func main() {

	// API routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world from GfG")
	})
	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		iban, ok := r.URL.Query()["iban"]

		if !ok || len(iban) == 0 {
			http.Error(w, "IBAN is missing", 400)
			return
		}
		if len(iban) != 1 {
			http.Error(w, "Can only validate one IBAN per request", 400)
			return
		}
		fmt.Println("IBAN: " + iban[0])
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)
	fmt.Println("countries are " + iban.Countries)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}
