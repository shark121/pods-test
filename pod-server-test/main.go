package main

import (
	"net/http"

	"github.com/shark121/pods-test/helpers"
)

func main() {

	// podAndRides := map[string]any{"randomRides": randomRides, "pod": pod, "ranked": rankRidesByProximityToPod(randomRides, pod)}

	helpers.UseHandler()

	print("server started running")

	http.ListenAndServe(":5000", nil)

}
