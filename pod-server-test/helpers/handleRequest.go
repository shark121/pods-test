package helpers

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/mmcloughlin/geohash"
	"github.com/pod-server-test/calc"
	t "github.com/pod-server-test/types"
)

func HandleRideRequest(ctx context.Context, client *firestore.Client, req t.RideObject) error {
	const precision = 6
	const maxDistance = 5.0
	const maxWaitTime = 300
	const maxCapacity = 4
	const defaultMaxAngle = 60
	const defaultMaxKm = 50

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
			// if !isCloseEnough(pod.PodOrigin, req.Origin, maxDistance) ||
			// 	!isCloseEnough(pod.PodDestination, req.Destination, maxDistance) {
			// 	continue
			// }

			if !isCloseEnough(pod, req, defaultMaxKm, defaultMaxAngle) {
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

func isCloseEnough(pod t.Pod, req t.RideObject, maxKm float64, maxAngle float64) bool {

	angleDifference :=
		calc.CalculateAngleBetweenRides(pod.PodOrigin, pod.PodDestination, req.Origin, req.Destination)

	podMid := calc.GetMidpoint(pod.PodOrigin, pod.PodDestination)
	reqMid := calc.GetMidpoint(req.Origin, req.Destination)

	podMidToLoc := t.Location{Lat: podMid["x"], Lng: podMid["y"]}
	reqMidToLoc := t.Location{Lat: reqMid["x"], Lng: reqMid["y"]}

	distanceBetweenMidPoints := calc.DistanceBetweenTwoPoints(podMidToLoc, reqMidToLoc)

	return (angleDifference < maxAngle) && (distanceBetweenMidPoints < maxKm)
}
