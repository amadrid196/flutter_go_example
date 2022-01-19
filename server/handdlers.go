package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}
	fmt.Fprintf(w, "Hello there %s", "vistor")
}

func getContries(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(countrie)
}

func addCountries(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(countrie)
	country := &Country{}

	err := json.NewDecoder(r.Body).Decode(country)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error to add")
		return
	}

	countrie = append(countrie, country)
	fmt.Fprintf(w, "Country was add")
}
