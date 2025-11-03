package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	podUtils "github.com/shark121/pods-test/podUtils"
	t "github.com/shark121/pods-test/types"
	check "github.com/shark121/pods-test/utils"
)

// type location struct {
// 	lat     float64
// 	long    float64
// 	placeId string
// }

// type RideObject struct {
// 	rideId       string
// 	rideTime     string
// 	rideStatus   string
// 	origin       location
// 	destination  location
// 	rideCapacity int16
// 	direction    float64
// 	rideDistance float64
// }

func CreateHandler() func(res http.ResponseWriter, req *http.Request) {

	return func(res http.ResponseWriter, req *http.Request) {

		res.Header().Set("Content-Type", "application/json")                     // Set Content-Type
		res.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // Example Cache-Control
		res.Header().Set("Pragma", "no-cache")                                   // Example Pragma
		res.Header().Set("Expires", "0")
		res.Header().Set("Access-Control-Allow-Origin", "*")

		if req.Method != "POST" {
			print("wrong method")
		}

		requestBody := req.Body

		reqBytes, err := io.ReadAll(requestBody)

		check.HasErr(err)

		fmt.Println(reqBytes)

		reqBodyJson := t.RideObject{}

		json.Unmarshal(reqBytes, &reqBodyJson)

		fmt.Println(reqBodyJson)

		pod := podUtils.MatchRide(reqBodyJson)

		// response, err := json.Marshal(data)

		// if err != nil {
		// 	print(err)
		// }
		json.NewEncoder(res).Encode(pod)

	}

}

func UseHandler() {
	getHandler := CreateHandler()
	http.HandleFunc("/", getHandler)
}
