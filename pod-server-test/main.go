package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"sort"

	"github.com/google/uuid"
	"github.com/pod-server-test/calc"
	"github.com/pod-server-test/helpers"
	"github.com/pod-server-test/types"

	directions "github.com/pod-server-test/mapUtils"
)

type User struct {
	Name            string         `json:"name"`
	Age             int16          `json:"age"`
	DefaultLocation types.Location `json:"defaultLocation"`
	Dob             string         `json:"dob"`
}

func createUser(name string, age int16, defaultLocation types.Location, dob string) User {
	return User{Name: name, Age: age, DefaultLocation: defaultLocation, Dob: dob}
}

func getDirection(origin types.Location, destination types.Location) float64 {
	// fmt.Println(origin, destination)
	return 0.000
}

func getDistance(origin types.Location, destination types.Location) float64 {
	// fmt.Println(origin, destination)
	return 0.000
}

func createRide(rideTime string, origin types.Location, destination types.Location, rideCapacity int8) types.RideObject {
	rideStatus := "pending"
	rideID := uuid.New().String()
	direction := getDirection(origin, destination)
	distance := getDistance(origin, destination)

	return types.RideObject{
		RideID:       rideID,
		RideTime:     rideTime,
		RideStatus:   rideStatus,
		Origin:       origin,
		Destination:  destination,
		RideCapacity: rideCapacity,
		Direction:    direction,
		RideDistance: distance,
	}
}

func generateCoordinatesCloseToLocation(loc types.Location) types.Location {
	randomLat := (rand.Float64() - 0.5) * float64(rand.Int31n(100))
	randomLng := (rand.Float64() - 0.5) * float64(rand.Int31n(100))

	return types.Location{
		Lat: loc.Lat + randomLat,
		Lng: loc.Lng + randomLng,
		// PlaceID: loc.PlaceID,
	}
}

func generateCoordinatesFarFromLocation(loc types.Location) types.Location {

	const earthRadiusKm = 6371.0

	distanceKm := 10 + rand.Float64()*40

	bearing := rand.Float64() * 2 * math.Pi

	latRad := loc.Lat * math.Pi / 180
	lngRad := loc.Lng * math.Pi / 180

	newLatRad := math.Asin(math.Sin(latRad)*math.Cos(distanceKm/earthRadiusKm) +
		math.Cos(latRad)*math.Sin(distanceKm/earthRadiusKm)*math.Cos(bearing))

	newLngRad := lngRad + math.Atan2(
		math.Sin(bearing)*math.Sin(distanceKm/earthRadiusKm)*math.Cos(latRad),
		math.Cos(distanceKm/earthRadiusKm)-math.Sin(latRad)*math.Sin(newLatRad),
	)

	newLat := newLatRad * 180 / math.Pi
	newLng := newLngRad * 180 / math.Pi

	return types.Location{
		Lat: newLat,
		Lng: newLng,
	}
}

type Path struct {
	Origin      types.Location
	Destination types.Location
	myInt       interface {
		Area() int
	}
}

func rankRidesByProximityToPod(ridesArray []types.RideObject, pod types.Pod) []types.RideObject {
	podMidpoint := calc.GetMidpoint(pod.PodOrigin, pod.PodDestination)

	fmt.Println(podMidpoint)

	rankedRides := make([]types.RideObject, len(ridesArray))
	for i, ride := range ridesArray {
		rideMidpoint := calc.GetMidpoint(ride.Origin, ride.Destination)
		distance := math.Sqrt(math.Pow(rideMidpoint["x"]-podMidpoint["x"], 2) + math.Pow(rideMidpoint["y"]-podMidpoint["y"], 2))
		bearing := calc.CalculateAngleBetweenRides(ride.Origin, ride.Destination, pod.PodOrigin, pod.PodDestination)

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

func generateRandomRides(number int8, local types.Location) []types.RideObject {
	rides := []types.RideObject{}
	for range number {
		origin := generateCoordinatesFarFromLocation(local)
		destination := generateCoordinatesFarFromLocation(local)

		if rand.Intn(2) == 1 {
			origin, destination = destination, origin
		}

		ride := createRide("2023-10-27T10:00:00Z", origin, destination, 4)
		rides = append(rides, ride)
	}
	return rides
}

func main() {

	seedOrigin := types.Location{Lat: 32.520845925634895, Lng: -92.71762474132422}
	seedDestination := types.Location{Lat: 32.52454000792701, Lng: -92.7070350339111}

	initRide := createRide("2023-10-27T10:00:00Z", seedOrigin, seedDestination, 4)

	pod := types.CreatePod(initRide)

	randomRides := generateRandomRides(3, seedOrigin)

	rankedRides := rankRidesByProximityToPod(randomRides, pod)

	for i := range 2 {
		pod.PodRides[rankedRides[i].RideID] = rankedRides[i]
	}

	podAndRides := map[string]any{"randomRides": randomRides, "pod": pod, "ranked": rankRidesByProximityToPod(randomRides, pod)}

	helpers.UseHandler(podAndRides)

	print("server started running")

	http.ListenAndServe(":5000", nil)

	directions.GetMapDirections(pod)

}
