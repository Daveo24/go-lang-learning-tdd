package v1

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestGetRequest(t *testing.T) {
	expected := `[{"title":"London","location_type":"City","woeid":44418,"latt_long":"51.506321,-0.12714"}]`

	actual := GetRequest()

	if actual != expected {
		t.Errorf("expected %s actual %s", expected, actual)
	}

}

func TestPostRequest(t *testing.T) {
	expected := 533
	jsonData := map[string]string{"title":"London","location_type":"City","latt_long":"51.506321,-0.12714"}
	jsonValue, _ := json.Marshal(jsonData)

	actual := len(PostRequest(jsonValue))

	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", strconv.FormatInt(int64(expected), 10), strconv.FormatInt(int64(actual), 10))
	}
}
