package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/reddotpay/demopay/process"
)

func init() {
	// set the variables

	process.CardPayAPI = os.Getenv("CARDPAY_API")
	process.CardAPI = os.Getenv("CARD_API")
	process.AltPayAPI = os.Getenv("ALTPAY_API")
	process.SecureAPI = os.Getenv("SECURE_API")

	process.OAuthURL = os.Getenv("OAUTH_URL")
	process.APIClientSecret = os.Getenv("API_CLIENT_SECRET")
	process.APIClientID = os.Getenv("API_CLIENT_ID")
	process.APIKey = os.Getenv("API_KEY")

	process.UIDomain = os.Getenv("UI_DOMAIN")
	process.DemoPayEndpoint = os.Getenv("DEMOPAY_ENDPOINT")
}

// main function to boot up everything
func main() {
	log.Println("Starting...")
	router := mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(process.Processor)
	log.Fatal(http.ListenAndServe(":9090", router))
}
