package podUtils

import (
	"context"
	"fmt"
	"time"

	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/mmcloughlin/geohash"
	"github.com/shark121/pods-test/calc"
	t "github.com/shark121/pods-test/types"
)

<<<<<<< HEAD:pod-server-test/podUtils/handleRideRequest.go
func HandleRideRequest(ctx context.Context, client *firestore.Client, req t.RideObject) error {
	const precision = 6
	const maxDistance = 5.0
	const maxWaitTime = 300
	const defaultMaxAngle = 60
	const defaultMaxKm = 50

	geo := geohash.EncodeWithPrecision(req.Origin.Lat, req.Origin.Lng, precision)
	fmt.Println("hash :", geo, geohash.Neighbors(geo))

	podsRef := client.Collection("pods")
=======
func HandleRideRequest(ctx context.Context, client *firestore.Client, req t.RideObject) (t.Pod, error) {
	const precision = 5
	const maxWaitTime = 30 * time.Minute
	const defaultMaxAngle = 60.0
	const defaultMaxKm = 50.0

	req.Direction = calc.CalculateBearing(req.Origin, req.Destination)
	req.RideDistance = calc.DistanceBetweenTwoPoints(req.Origin, req.Destination)

	geo := geohash.EncodeWithPrecision(req.Origin.Lat, req.Origin.Lng, precision)

	tierDoc := fmt.Sprintf("tier_%d", req.RideCapacity)
	podsRef := client.Collection("pods").Doc(tierDoc).Collection("activePods")
>>>>>>> d39e2fc54b334054a6b30ba699952213009a9465:pod-server-test/helpers/handleRequest.go

	nearby := append([]string{geo}, geohash.Neighbors(geo)...)
	var matchedPod *firestore.DocumentSnapshot
	var podToMatch t.Pod

	for _, h := range nearby {
		iter := podsRef.Where(geo, "==", h).Documents(ctx)
		for {
			doc, err := iter.Next()
			if err != nil {
				break
			}
			var pod t.Pod
			if err := doc.DataTo(&pod); err != nil {
				fmt.Println("Error converting pod data: ", err)
				continue
			}

<<<<<<< HEAD:pod-server-test/podUtils/handleRideRequest.go
			if len(pod.PodRides) == int(pod.PodCapacity) {
				continue
			}

			if time.Now().Unix()-pod.CreatedAt.Unix() > maxWaitTime {
				continue
			}

			if isCloseEnough(pod, req, defaultMaxKm, defaultMaxAngle) {
				matchedPod = doc
				break
			}

=======
			if len(pod.PodRides) >= int(pod.PodCapacity) {
				fmt.Println("Pod is full")
				continue
			}

			if time.Since(pod.CreatedAt) > maxWaitTime {
				fmt.Println("Pod is too old")
				continue
			}

			if pod.PodStatus == "dispatched" || pod.PodStatus == "completed" {
				fmt.Println("Pod is dispatched or completed")
				continue
			}

			podMid := calc.GetMidpoint(pod.PodOrigin, pod.PodDestination)
			reqMid := calc.GetMidpoint(req.Origin, req.Destination)
			podMidLoc := t.Location{Lat: podMid["y"], Lng: podMid["x"]}
			reqMidLoc := t.Location{Lat: reqMid["y"], Lng: reqMid["x"]}

			distanceBetweenMidPoints := calc.DistanceBetweenTwoPoints(podMidLoc, reqMidLoc)
			angleDifference := calc.CalculateAngleBetweenRides(pod.PodOrigin, pod.PodDestination, req.Origin, req.Destination)

			detourPenalty := (distanceBetweenMidPoints / defaultMaxKm) * 100.0
			anglePenalty := (angleDifference / defaultMaxAngle) * 100.0

			fmt.Println(detourPenalty, anglePenalty)
			if detourPenalty > 100 || anglePenalty > 100 {
				fmt.Println("Detour penalty or angle penalty is too high")
				continue
			}

			// First pod that passes constraints is deemed successful (Greedy Match)
			matchedPod = doc
			podToMatch = pod
			break
>>>>>>> d39e2fc54b334054a6b30ba699952213009a9465:pod-server-test/helpers/handleRequest.go
		}

		if matchedPod != nil {
			break
		}
	}

	if matchedPod != nil {
		podToMatch.PodRides[req.RideID] = req
		_, err := matchedPod.Ref.Set(ctx, podToMatch)
		return podToMatch, err
	}

	newPod := t.CreatePod(req)
	newPod.Geohash = geo
	_, err := podsRef.Doc(newPod.PodID).Set(ctx, newPod)
<<<<<<< HEAD:pod-server-test/podUtils/handleRideRequest.go
	return err
}

//TODO: work on this function

func isCloseEnough(pod t.Pod, req t.RideObject, maxKm float64, maxAngle float64) bool {

	angleDifference :=
		calc.CalculateAngleBetweenRides(pod.Origin, pod.Destination, req.Origin, req.Destination)

	podMid := calc.GetMidpoint(pod.Origin, pod.Destination)
	reqMid := calc.GetMidpoint(req.Origin, req.Destination)

	podMidToLoc := t.Location{Lat: podMid["y"], Lng: podMid["x"]}
	reqMidToLoc := t.Location{Lat: reqMid["y"], Lng: reqMid["x"]}

	distanceBetweenMidPoints := calc.DistanceBetweenTwoPoints(podMidToLoc, reqMidToLoc)

	return (angleDifference < maxAngle) && (distanceBetweenMidPoints < maxKm)
=======
	return newPod, err
>>>>>>> d39e2fc54b334054a6b30ba699952213009a9465:pod-server-test/helpers/handleRequest.go
}
