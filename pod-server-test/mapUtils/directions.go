package directions

import (
	"context"
	"fmt"

	"github.com/shark121/pods-test/types"

	util "github.com/shark121/pods-test/utils"

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
	cfg := util.LoadConfig()

	client, err := m.NewClient(m.WithAPIKey(cfg.MapsKey))

	if err != nil {
		fmt.Println("error loading maps", err)
		return
	}

	waypoints := []string{}

	for _, value := range pod.PodRides {
		waypoints = append(waypoints, FormatToString(value.Origin))
		waypoints = append(waypoints, FormatToString(value.Destination))
	}

	direction :=
		&m.DirectionsRequest{
			Origin:      FormatToString(pod.Origin),
			Destination: FormatToString(pod.Destination),
			Waypoints:   waypoints,
			Optimize:    true,
		}

	route, stops, err := client.Directions(ctx, direction)

	if err != nil {
		fmt.Println("directions request error", err)
	}

	fmt.Println(route, "\n", stops)
}
