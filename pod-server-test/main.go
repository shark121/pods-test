package main

import (
	"net/http"

	firebaseutils "github.com/shark121/pods-test/firebaseUtils"
	"github.com/shark121/pods-test/helpers"
	"github.com/shark121/pods-test/utils"
)

func main() {

	// podAndRides := map[string]any{"randomRides": randomRides, "pod": pod, "ranked": rankRidesByProximityToPod(randomRides, pod)}
	appInstance, err := firebaseutils.CreateAppInstance()

	utils.HasErr(err)

	helpers.UseHandler(appInstance)

	appInstance.Add(map[string]any{"hello": "world"})

	print("server started running")

	http.ListenAndServe(":5000", nil)

}
