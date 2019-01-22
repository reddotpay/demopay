package proccess

import (
	"io/ioutil"
	"net/http"
)

type Card struct {
	Merchant string  `json:"merchant"`
	Number   string  `json:"number"`
	Expiry   *Expiry `json:"expiry"`
	CCV      string  `json:"ccv"`
	Name     string  `json:"name"`
}
type Expiry struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}
type Address struct {
	City string `json:"city"`
}

func CreateCard(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

}
