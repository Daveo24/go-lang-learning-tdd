package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Weather struct {
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LattLong     string `json:"latt_long"`
}

var locations = map[string]Weather{
	"London":  {Title: "London", LocationType: "Capital", Woeid: 44418, LattLong: "51.506321,-0.12714"},
	"Berlin":  {Title: "Berlin", LocationType: "City", Woeid: 2345496, LattLong: "52.5200,-13.4050"},
	"Florida": {Title: "Florida", LocationType: "State", Woeid: 11111, LattLong: "27.6648, -81.5158"},
}

func (w Weather) toJSON() []byte {
	ToJSON, err := json.Marshal(w)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

//FromJSON pass data as a byte
func FromJSON(data []byte) Weather {
	weather := Weather{}
	err := json.Unmarshal(data, &weather)
	if err != nil {
		panic(err)
	}
	return weather
}

func AllLocations() []Weather {
	values := make([]Weather, len(locations))
	i := 0
	for _, location := range locations {
		values[i] = location
		i++
	}
	return values
}

func GetWeatherLocation(title string) (Weather, bool) {
	location, found := locations[title]
	return location, found
}

func CreateWeatherLocation(location Weather) (string, bool) {
	_, exists := locations[location.Title]
	if exists {
		return "", false
	}
	locations[location.Title] = location
	return location.Title, true
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

func LocationsHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		locations := AllLocations()
		writeJSON(w, locations)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		location := FromJSON(body)
		title, created := CreateWeatherLocation(location)
		if created {
			w.Header().Add("Location", "/api/weather/locations"+title)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func LocationHandleFunc(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Path[len("/api/weather/locations/"):]
	switch method := r.Method; method {
	case http.MethodGet:
		location, found := GetWeatherLocation(location)
		if found {
			writeJSON(w, location)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		//implement update and delete
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}

}

func main() {
	http.HandleFunc("/api/weather/locations", LocationsHandleFunc)
	http.HandleFunc("/api/weather/locations/", LocationHandleFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
