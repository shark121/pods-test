package podUtils

import (
	"fmt"
	"math"
	"sort"

	"github.com/shark121/pods-test/calc"
	testing "github.com/shark121/pods-test/testingUtils"
	types "github.com/shark121/pods-test/types"
)

type Path struct {
	Origin      types.Location
	Destination types.Location
	myInt       interface {
		Area() int
	}
}

func RankRidesByProximityToPod(ridesArray []types.RideObject, pod types.Pod) []types.RideObject {
	podMidpoint := calc.GetMidpoint(pod.Origin, pod.Destination)

	fmt.Println(podMidpoint)

	rankedRides := make([]types.RideObject, len(ridesArray))
	for i, ride := range ridesArray {
		rideMidpoint := calc.GetMidpoint(ride.Origin, ride.Destination)
		distance := math.Sqrt(math.Pow(rideMidpoint["x"]-podMidpoint["x"], 2) + math.Pow(rideMidpoint["y"]-podMidpoint["y"], 2))
		bearing := calc.CalculateAngleBetweenRides(ride.Origin, ride.Destination, pod.Origin, pod.Destination)

		rideCopy := ride
		rideCopy.RideDistance = distance
		rideCopy.RideBearing = &bearing
		rankedRides[i] = rideCopy
	}

	sort.Slice(rankedRides, func(i, j int) bool {
		return rankedRides[i].RideDistance < rankedRides[j].RideDistance
	})

	return rankedRides
}

func MatchRide(ride types.RideObject) types.Pod {

	pod := types.CreatePod(ride)

	randomRides := testing.GenerateRandomRides(3, ride.Origin)

	rankedRides := RankRidesByProximityToPod(randomRides, pod)

	for i := range 2 {
		pod.PodRides[rankedRides[i].RideID] = rankedRides[i]
	}

	return pod
}
