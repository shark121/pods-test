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

func GetMapDirections(rides []types.RideObject) {
	cfg := util.ReadConfig("C:/Users/HP/Desktop/nuclear-launch-codes/pods-test/config.json")

	client, err := m.NewClient(m.WithAPIKey(cfg.Maps_key))

	if err != nil {
		fmt.Println("error loading maps", err)
	}

	client.Directions(ctx, x)

	// fmt.Println(cfg.Maps_key, "from directions")
	// client.Directions(ctx)
}
