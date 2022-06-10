package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//curl --request GET --url \
// > 'https://api.tomorrow.io/v4/timelines?location=-73.98529171943665,40.75872069597532&fields=temperature&timesteps=1h&units=metric&apikey=Xw1fiUbiRuepIb6RFFqTjaoQWhQOE1PR'

func main() {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://api.tomorrow.io/v4/timelines?location=-73.98529171943665,40.75872069597532&fields=temperature",
		nil,
	)
	if err != nil {
		log.Fatalf("error creating HTTP request: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("apikey", os.Getenv("CLIMACELL_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading HTTP Response Body: %v", err)
	}
	log.Println("We got response:", string(responseBytes))
}
