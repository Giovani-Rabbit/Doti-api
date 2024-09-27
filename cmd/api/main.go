package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("MINHAUUU")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/items", getItems).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
