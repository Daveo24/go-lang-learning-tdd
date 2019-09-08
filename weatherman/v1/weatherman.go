package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//GetRequest example
func GetRequest() string {
	response, err := http.Get("https://www.metaweather.com/api/location/search/?query=london")
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	json := fmt.Sprintln(string(data))
	s := strings.TrimSpace(json)
	return s
}

//PostRequest Example
func PostRequest(jsonValue []byte) string {
	response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	postJson := fmt.Sprintln(string(data))

	return postJson
}

func main() {
	fmt.Println("Start...")
	fmt.Printf(GetRequest() + "\n")

	jsonData := map[string]string{"title": "London", "location_type": "City", "latt_long":"51.506321,-0.12714"}
	jsonValue, _ := json.Marshal(jsonData)
	fmt.Println(PostRequest(jsonValue))

	fmt.Println("Finish!")
}
