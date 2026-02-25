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
	const precision = 5
	const maxWaitTime = 30 * time.Minute
	const defaultMaxAngle = 60.0
	const defaultMaxKm = 50.0

	// We calculate distance and direction dynamically
	req.Direction = calc.CalculateBearing(req.Origin, req.Destination)
	req.RideDistance = calc.DistanceBetweenTwoPoints(req.Origin, req.Destination)

	geo := geohash.EncodeWithPrecision(req.Origin.Lat, req.Origin.Lng, precision)
	podsRef := client.Collection("pods")

	nearby := append([]string{geo}, geohash.Neighbors(geo)...)
	var matchedPod *firestore.DocumentSnapshot
	var podToMatch t.Pod

	for _, h := range nearby {
		iter := podsRef.Where("geohash", "==", h).Documents(ctx)
		for {
			doc, err := iter.Next()
			if err != nil {
				break
			}
			var pod t.Pod
			if err := doc.DataTo(&pod); err != nil {
				continue
			}

			// Capacity check
			currentOccupancy := 0
			for _, r := range pod.PodRides {
				currentOccupancy += int(r.RideCapacity)
			}
			if currentOccupancy+int(req.RideCapacity) > int(pod.PodCapacity) {
				continue
			}

			// Expiration check
			if time.Since(pod.CreatedAt) > maxWaitTime {
				continue
			}

			// Destination and trajectory constraint check
			if !isCloseEnough(pod, req, defaultMaxKm, defaultMaxAngle) {
				continue
			}

			matchedPod = doc
			podToMatch = pod
			break
		}

		if matchedPod != nil {
			break
		}
	}

	if matchedPod != nil {
		podToMatch.PodRides[req.RideID] = req
		_, err := matchedPod.Ref.Set(ctx, podToMatch)
		return err
	}

	newPod := t.CreatePod(req)
	newPod.Geohash = geo
	_, err := podsRef.Doc(newPod.PodID).Set(ctx, newPod)
	return err
}

func isCloseEnough(pod t.Pod, req t.RideObject, maxKm float64, maxAngle float64) bool {
	angleDifference := calc.CalculateAngleBetweenRides(pod.PodOrigin, pod.PodDestination, req.Origin, req.Destination)

	// Since they both have Lat/Lng coordinates, calculate distance between their midpoints
	podMid := calc.GetMidpoint(pod.PodOrigin, pod.PodDestination)
	reqMid := calc.GetMidpoint(req.Origin, req.Destination)

	podMidLoc := t.Location{Lat: podMid["y"], Lng: podMid["x"]}
	reqMidLoc := t.Location{Lat: reqMid["y"], Lng: reqMid["x"]}

	distanceBetweenMidPoints := calc.DistanceBetweenTwoPoints(podMidLoc, reqMidLoc)

	return (angleDifference <= maxAngle) && (distanceBetweenMidPoints <= maxKm)
}
