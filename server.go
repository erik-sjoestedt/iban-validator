package main

import (
	"fmt"
	iban_validator "iban-validator/iban"
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
		is_valid := iban_validator.Validate(iban[0])
		if is_valid {
			fmt.Fprintf(w, "IBAN is valid")
		} else {
			http.Error(w, "IBAN is invalid", 400)
		}
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)
	fmt.Println("countries are " + iban_validator.Countries)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}
