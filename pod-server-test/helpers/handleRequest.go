package helpers

import (
	"context"
	"math"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/mmcloughlin/geohash"
	t "github.com/pod-server-test/types"
)

func HandleRideRequest(ctx context.Context, client *firestore.Client, req t.RideObject) error {
	const precision = 6
	const maxDistance = 5.0
	const maxWaitTime = 300
	const maxCapacity = 4

	geo := geohash.EncodeWithPrecision(req.Origin.Lat, req.Origin.Lng, precision)
	podsRef := client.Collection("pods")

	nearby := append([]string{geo}, geohash.Neighbors(geo)...)
	var matchedPod *firestore.DocumentSnapshot

	for _, h := range nearby {
		iter := podsRef.Where("geohash", "==", h).Documents(ctx)
		for {
			doc, err := iter.Next()
			if err != nil {
				break
			}
			var pod t.Pod
			doc.DataTo(&pod)

			if len(pod.PodRides) >= int(pod.PodCapacity) {
				continue
			}
			if time.Now().Unix()-pod.CreatedAt.Unix() > maxWaitTime {
				continue
			}
			if !isCloseEnough(pod.PodOrigin, req.Origin, maxDistance) ||
				!isCloseEnough(pod.PodDestination, req.Destination, maxDistance) {
				continue
			}

			matchedPod = doc
			break
		}

		if matchedPod != nil {
			break
		}
	}

	if matchedPod != nil {

		_, err := matchedPod.Ref.Update(ctx, []firestore.Update{
			{Path: "occupants", Value: firestore.Increment(1)},
		})
		return err
	}

	newPod := t.CreatePod(req)

	_, err := podsRef.Doc(newPod.PodID).Set(ctx, newPod)
	return err
}

//TODO: work on this function

func isCloseEnough(loc1, loc2 t.Location, maxKm float64) bool {
	const R = 6371
	dLat := toRad(loc2.Lat - loc1.Lat)
	dLng := toRad(loc2.Lng - loc1.Lng)
	lat1 := toRad(loc1.Lat)
	lat2 := toRad(loc2.Lat)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dLng/2)*math.Sin(dLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R*c <= maxKm
}

func toRad(deg float64) float64 {
	return deg * math.Pi / 180
}
