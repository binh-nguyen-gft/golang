package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var inventory []Item

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	_ = json.NewDecoder(r.Body).Decode(&newItem)
	inventory = append(inventory, newItem)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
}

func main() {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/items", AddItem).Methods("POST")

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

// curl -X POST -H "Content-Type: application/json" -d '{"id":"1", "name":"Book", "price":10}' http://localhost:8000/items