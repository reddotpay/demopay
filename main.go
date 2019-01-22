package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/card", proccess.CreateCard).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090", router))
}
