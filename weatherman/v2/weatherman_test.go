package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWeatherToJSON(t *testing.T) {
	expected := `{"title":"London","location_type":"City","woeid":44418,"latt_long":"51.506321,-0.12714"}`

	weather := Weather{Title: "London", LocationType: "City", Woeid: 44418, LattLong: "51.506321,-0.12714"}

	json := weather.toJSON()

	actual := string(json)

	if actual != expected {
		t.Errorf("expected %s actual %s", expected, actual)
	}
}

func TestWeatherFromJSON(t *testing.T) {
	expected := Weather{Title: "London", LocationType: "City", Woeid: 44418, LattLong: "51.506321,-0.12714"}
	json := []byte(`{"title":"London","location_type":"City","woeid":44418,"latt_long":"51.506321,-0.12714"}`)
	weather := FromJSON(json)
	actual := weather

	if actual != expected {
		t.Errorf("expected %q actual %q", expected, actual)
	}
}

func TestAllLocations(t *testing.T) {
	expected := 3

	locations := AllLocations()

	actual := len(locations)

	if actual != expected {
		t.Errorf("expected %d actual %d", expected, actual)
	}
}

func TestGetWeatherLocation(t *testing.T) {

	t.Run("Assert expected and actual the good old fashioned way", func(t *testing.T) {
		expected := "London"

		weather := Weather{Title: "London", LocationType: "City", Woeid: 44418, LattLong: "51.506321,-0.12714"}

		weather, _ = GetWeatherLocation("London")

		actual := weather.Title

		if actual != expected {
			t.Errorf("expected %s actual %s", expected, actual)
		}

	})
	t.Run("Assert expected and actual", func(t *testing.T) {
		weather := Weather{Title: "London", LocationType: "City", Woeid: 44418, LattLong: "51.506321,-0.12714"}

		weather, _ = GetWeatherLocation("London")

		assert.Equal(t, "London", weather.Title, "Title not London")
	})
}

func TestCreateWeatherLocation(t *testing.T) {
	expectedTitle := "Sunderland"
	expectedCreation := true

	location := Weather{Title: "Sunderland", LocationType: "City", Woeid: 7862, LattLong: "77, -122"}
	title, created := CreateWeatherLocation(location)

	actualTitle := title
	actualCreation := created

	if actualTitle != expectedTitle {
		t.Errorf("expected %s actual %s", expectedTitle, actualTitle)
	}

	if actualCreation != expectedCreation {
		t.Errorf("expected %t actual %t", expectedCreation, actualCreation)
	}
}
