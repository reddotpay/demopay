package process

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/oauth2/clientcredentials"
)

// Processor process the request from the frontend to the API3
func Processor(w http.ResponseWriter, r *http.Request) {
	var endpoint string

	re := regexp.MustCompile(DemoPayEndpoint + `/(\w+)/(.*)`)
	match := re.FindStringSubmatch(r.URL.Path)

	// toggle between the different endpoint
	switch match[1] {
	case "card":
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
	fmt.Println("Endpoint", endpoint)
	fmt.Println("Method", r.Method)

	// Set the OAuth client
	config := clientcredentials.Config{
		ClientID:     APIClientID,
		ClientSecret: APIClientSecret,
		TokenURL:     OAuthURL,
	}

	oauth, err := config.Token(context.Background())
	if err != nil {

	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", oauth.AccessToken)
	r.Header.Set("x-api-key", APIKey)

	// make request to card pay
	body, _ := ioutil.ReadAll(r.Body)
	data := bytes.NewBuffer(body)
	req, err := http.NewRequest(r.Method, endpoint, data)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}

	w.Header().Set("Access-Control-Allow-Origin", UIDomain)
	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}
