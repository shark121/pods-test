package directions

import (
	"context"
	"fmt"

	util "github.com/pod-server-test/utils"
)

var ctx context.Context = context.Background()

func GetMapDirections(rides any) {
	// client, err := m.NewClient(m.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))

	// if err != nil {
	// 	fmt.Println("error loading maps", err)
	// }

	cfg := util.ReadConfig("C:/Users/HP/Desktop/nuclear-launch-codes/pods-test/config.json")

	fmt.Println(cfg.Maps_key, "from directions")
	// client.Directions(ctx)
}
