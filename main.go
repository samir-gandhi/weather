package main

import (
	// "encoding/json"
	// "fmt"
	"github.com/samir-gandhi/weather/weather"
	"log"
	// "net/http"
	"os"
)

//curl --request GET --url \
// > 'https://api.tomorrow.io/v4/timelines?location=-73.98529171943665,40.75872069597532&fields=temperature&timesteps=1h&units=metric&apikey=Xw1fiUbiRuepIb6RFFqTjaoQWhQOE1PR'

// func main() {
// 	req, err := http.NewRequest(
// 		http.MethodGet,
// 		"https://api.tomorrow.io/v4/timelines?location=39.542534,-104.965448&fields=temperature&units=imperial",
// 		nil,
// 	)
// 	if err != nil {
// 		log.Fatalf("error creating HTTP request: %v", err)
// 	}
// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("apikey", os.Getenv("weather_API_KEY"))

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		log.Fatalf("error sending HTTP request: %v", err)
// 	}
// 	if res.StatusCode != http.StatusOK {
// 		fmt.Println("Unexpect HTTP Status: ", res.Status)
// 	}
// 	var weatherSamples weather.Weather
// 	d := json.NewDecoder(res.Body)
// 	if err := d.Decode(&weatherSamples); err != nil {
// 		log.Fatalf("error deserializing weather data")
// 	}
// 	var nilValues weather.Values
// 	for _, w := range weatherSamples.Data.Timelines[0].Intervals {
// 		if w.Values != nilValues {
// 			log.Printf("The temperature at %s is %f degrees",
// 				w.StartTime, w.Values.Temperature)
// 		} else {
// 			log.Printf("No temperature data available at %s\n",
// 				w.StartTime)
// 		}
// 	}

// 	log.Println("We got response:", weatherSamples)
// }

func main() {
	c := weather.New(os.Getenv("CLIMACELL_API_KEY"))
	weatherSamples, err := c.HourlyForecast(weather.ForecastArgs{
		LatLon: &weather.LatLon{Lat: 42.3826, Lon: -71.146},
		Fields: []string{"temperature"},
		Units: "imperial",
	})
	if err != nil {
		log.Fatalf("error getting forecast data: %v", err)
	}

	var tempAtFive *weather.Values
	for i, w := range weatherSamples.Data.Timelines[0].Intervals {
			if w.StartTime.Hour() == 21 {
					tempAtFive = &weatherSamples.Data.Timelines[0].Intervals[i].Values
					break
			}
	}
	if tempAtFive == nil {
		log.Print("just printing this ", tempAtFive.Temperature)
	} else if t := *&tempAtFive.Temperature; t < 60 {
			log.Printf("It'll be %f out. Better make some hot tea! 🌺🍵\n", t)
	} else {
			log.Printf("It'll be %f out. Iced tea it is! 🌺🍹\n", t)
	}
}
