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

func FormatToString(loc types.Location) string {

	return fmt.Sprintf("%f,%f", loc.Lat, loc.Lng)
}

func GetMapDirections(pod types.Pod) {
	cfg := util.ReadConfig("C:/Users/HP/Desktop/nuclear-launch-codes/pods-test/config.json")

	client, err := m.NewClient(m.WithAPIKey(cfg.Maps_key))

	if err != nil {
		fmt.Println("error loading maps", err)
	}

	waypoints := []string{}

	for _, value := range pod.PodRides {
		waypoints = append(waypoints, FormatToString(value.Origin))
		waypoints = append(waypoints, FormatToString(value.Destination))
	}

	direction :=
		&m.DirectionsRequest{
			Origin:      fmt.Sprintf("%f,%f", pod.PodOrigin.Lat, pod.PodOrigin.Lng),
			Destination: fmt.Sprintf("%f,%f", pod.PodDestination.Lat, pod.PodDestination.Lng),
			Waypoints:   waypoints,
		}

	route, stops, err := client.Directions(ctx, direction)

	if err != nil {
		fmt.Println("directions request error", err)
	}

	fmt.Println(route, stops)
}
