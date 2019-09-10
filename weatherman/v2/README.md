This is a rest api created using tdd. To test you can use the following `go test` or `go test -v`.

If you would like to check code coverage you can use the command `go test -cover`

To build the app at the command line use `go build && ./v2` if using Mac OS and `go build && ./v2.exe` for Windows OS

Once built you can either curl or use postman to call the simple rest api.

The following will return results providing your default port is 8080, if you have set up an export on your port it will be found there eg 9090

GET `http://localhost:8080/api/weather/locations`
POST `http://localhost:8080/api/weather/locations` with json in the body eg `{
                                                                             	"title":"Newcastle",
                                                                             	"location_type":"City",
                                                                             	"woeid":44418,
                                                                             	"latt_long":"51.506321,-0.12714"
                                                                             }`
GET `http://localhost:8080/api/weather/locations/Sunderland`                                                                          