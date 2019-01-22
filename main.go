package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reddotpay/demopay/process"
)

func init() {
	// set the variables

}

// main function to boot up everything
func main() {
	fmt.Println("Starting...")
	router := mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(process.Processor)
	log.Fatal(http.ListenAndServe(":9090", router))
}
