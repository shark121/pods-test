package directions

import (
	"context"
	"fmt"

	"github.com/pod-server-test/types"

	util "github.com/pod-server-test/utils"

	m "googlemaps.github.io/maps"
)

var ctx context.Context = context.Background()

type TravelMode string

const (
	Driving TravelMode = "driving"
	Walking TravelMode = "walking"
)

type DirectionsRequest struct {
	Origin string

	Destination string

	Mode TravelMode
}

func GetMapDirections(pod types.Pod) {
	cfg := util.ReadConfig("C:/Users/HP/Desktop/nuclear-launch-codes/pods-test/config.json")

	client, err := m.NewClient(m.WithAPIKey(cfg.Maps_key))

	if err != nil {
		fmt.Println("error loading maps", err)
	}

	direction :=
		// &m.DirectionsRequest{Origin: "San Fransisco", Destination: "New Jersey"}
		&m.DirectionsRequest{Origin: fmt.Sprintf("%f,%f", pod.PodOrigin.Lat, pod.PodOrigin.Lng), Destination: fmt.Sprintf("%f,%f", pod.PodDestination.Lat, pod.PodDestination.Lng)}

	route, waypoints, err := client.Directions(ctx, direction)

	if err != nil {
		fmt.Println("directions request error", err)
	}

	fmt.Println(route, waypoints)
}
