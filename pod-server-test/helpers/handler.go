package helpers

import (
	"encoding/json"
	"net/http"
)

type location struct {
	lat     float64
	long    float64
	placeId string
}

type RideObject struct {
	rideId       string
	rideTime     string
	rideStatus   string
	origin       location
	destination  location
	rideCapacity int16
	direction    float64
	rideDistance float64
}

func CreateServer(data any) func(res http.ResponseWriter, req *http.Request) {

	return func(res http.ResponseWriter, req *http.Request) {

		res.Header().Set("Content-Type", "application/json")                     // Set Content-Type
		res.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // Example Cache-Control
		res.Header().Set("Pragma", "no-cache")                                   // Example Pragma
		res.Header().Set("Expires", "0")
		res.Header().Set("Access-Control-Allow-Origin", "*")

		if req.Method != "GET" {
			print("wrong method")
		}

		response, err := json.Marshal(data)

		if err != nil {
			print(err)
		}

		res.Write(response)
	}

}

func UseHandler(data any) {
	getHandler := CreateServer(data)
	http.HandleFunc("/", getHandler)
}
