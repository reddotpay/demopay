package process

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"golang.org/x/oauth2/clientcredentials"
)

// Processor process the request from the frontend to the API3
func Processor(w http.ResponseWriter, r *http.Request) {
	var endpoint string

	log.Println("request endpoint", r.URL.Path)
	re := regexp.MustCompile(`/(\w+)/(.*)`)
	match := re.FindStringSubmatch(r.URL.Path)

	// toggle between the different endpoint
	switch match[1] {
	case "card":
		fmt.Println("CardAPI", CardAPI)
		endpoint = CardAPI + match[2]
		break
	case "cardpay":
		endpoint = CardPayAPI + match[2]
		break
	case "altpay":
		endpoint = AltPayAPI + match[2]
		break
	case "secure":
		endpoint = SecureAPI + match[2]
		break
	}

	// Set the OAuth client
	config := clientcredentials.Config{
		ClientID:     APIClientID,
		ClientSecret: APIClientSecret,
		TokenURL:     OAuthURL,
	}

	oauth, err := config.Token(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// make request to card pay
	body, _ := ioutil.ReadAll(r.Body)
	data := bytes.NewBuffer(body)
	log.Println("request body", string(body))

	req, err := http.NewRequest(r.Method, endpoint, data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", oauth.AccessToken)
	req.Header.Set("x-api-key", APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("response body", string(respBody))
	w.Header().Set("Access-Control-Allow-Origin", UIDomain)
	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}
