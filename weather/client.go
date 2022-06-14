package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.tomorrow.io",
	Path:   "/v4/",
}

type Client struct {
	c      *http.Client
	apiKey string
}

func New(apiKey string) *Client {
	c := &http.Client{Timeout: time.Minute}

	return &Client{
		c:      c,
		apiKey: apiKey,
	}
}

// separating LatLon into it's own struct makes it always not null
// And for a LatLon struct, the zero value is still a valid pair
// of coordinates; LatLon{0,0} is the coordinates of where
// the equator and prime meridian meet.
type LatLon struct{ Lat, Lon float64 }

type ForecastArgs struct {
	LatLon *LatLon
	Fields []string
	Units  string
}

func (args ForecastArgs) QueryParams() url.Values {
	q := make(url.Values)
	if args.LatLon != nil {
		latLon := strings.Join([]string{strconv.FormatFloat(args.LatLon.Lat, 'f', -1, 64), ",", strconv.FormatFloat(args.LatLon.Lon, 'f', -1, 64)}, "")
		q.Add("location", latLon)
	}

	if len(args.Fields) > 0 {
		q.Add("fields", strings.Join(args.Fields, ","))
	}

	if args.Units != "" {
		q.Add("units", args.Units)
	}

	return q
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
}

func (err *ErrorResponse) Error() string {
	if err.ErrorCode == "" {
		return fmt.Sprintf("%d API error: %s", err.StatusCode, err.Message)
	}
	return fmt.Sprintf("%d (%s) API error: %s", err.StatusCode, err.ErrorCode, err.Message)
}

func (c *Client) HourlyForecast(args ForecastArgs) (Weather, error) {
	var nilWeather Weather
	endpt := baseURL.ResolveReference((&url.URL{Path: "timelines"}))
	req, err := http.NewRequest("GET", endpt.String(), nil)
	if err != nil {
		return nilWeather, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("apikey", c.apiKey)
	req.URL.RawQuery = args.QueryParams().Encode()

	res, err := c.c.Do(req)
	if err != nil {
		return nilWeather, err
	}

	/*deserialize success or error response and return its data*/
	defer res.Body.Close()
	switch res.StatusCode {
	case 200:
    var weatherSamples Weather
    if err := json.NewDecoder(res.Body).Decode(&weatherSamples); err != nil {
        log.Fatalf("error deserializing weather data: %v",err)
    }
		// if err := json.NewDecoder(res.Body).Decode(&weatherSamples); err != nil {
		// 	return nil, err
		// }
		return weatherSamples, nil
	case 400, 401, 403, 500:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nilWeather, err
		}

		if errRes.StatusCode == 0 {
			errRes.StatusCode = res.StatusCode
		}
		return nilWeather, &errRes
	default:
		return nilWeather, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

}
