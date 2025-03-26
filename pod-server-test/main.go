package main

import (
	"math/rand"
	"net/http"

	"github.com/google/uuid"
	"github.com/pod-server-test/helpers"
)

type Location struct {
	Lat     float64 `json:"lat"`
	Long    float64 `json:"long"`
	PlaceID string  `json:"placeId"`
}

type User struct {
	Name            string   `json:"name"`
	Age             int16    `json:"age"`
	DefaultLocation Location `json:"defaultLocation"`
	Dob             string   `json:"dob"`
}

func createUser(name string, age int16, defaultLocation Location, dob string) User {
	return User{Name: name, Age: age, DefaultLocation: defaultLocation, Dob: dob}
}

type RideObject struct {
	RideID       string   `json:"rideId"`
	RideTime     string   `json:"rideTime"`
	RideStatus   string   `json:"rideStatus"`
	Origin       Location `json:"origin"`
	Destination  Location `json:"destination"`
	RideCapacity int16    `json:"rideCapacity"`
	Direction    float64  `json:"direction"`
	RideDistance float64  `json:"rideDistance"`
}

func getDirection(origin Location, destination Location) float64 {
	// fmt.Println(origin, destination)
	return 0.000
}

func getDistance(origin Location, destination Location) float64 {
	// fmt.Println(origin, destination)
	return 0.000
}

func createRide(rideTime string, origin Location, destination Location, rideCapacity int16) RideObject {
	rideStatus := "pending"
	rideID := uuid.New().String()
	direction := getDirection(origin, destination)
	distance := getDistance(origin, destination)

	return RideObject{
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

type Pod struct {
	PodOrigin      Location              `json:"origin"`
	PodDestination Location              `json:"destination"`
	PodCapacity    int16                 `json:"podCapacity"`
	PodStatus      string                `json:"podStatus"`
	PodID          string                `json:"podId"`
	PodRides       map[string]RideObject `json:"podRides"`
	Waypoints      []Location            `json:"waypoints"`
	PodDirection   float64               `json:"podDirection"`
	PodDistance    float64               `json:"podDistance"`
	Time           string                `json:"time"`
}

func createPod(ride RideObject) Pod {
	podStatus := "pending"
	podID := uuid.New().String()
	podRides := map[string]RideObject{ride.RideID: ride}
	podDirection := ride.Direction
	podDistance := ride.RideDistance
	time := ride.RideTime

	return Pod{
		PodOrigin:      ride.Origin,
		PodDestination: ride.Destination,
		PodCapacity:    ride.RideCapacity,
		PodStatus:      podStatus,
		PodID:          podID,
		PodRides:       podRides,
		Waypoints:      []Location{},
		PodDirection:   podDirection,
		PodDistance:    podDistance,
		Time:           time,
	}
}

func (p *Pod) addRide(ride RideObject) {
	p.PodRides[ride.RideID] = ride
}

func (p *Pod) removeRide(ride RideObject) {
	delete(p.PodRides, ride.RideID)
}

func generateCoordinatesCloseToLocation(loc Location) Location {
	randomLat := (rand.Float64() - 0.5) * float64(rand.Int31n(100))
	randomLong := (rand.Float64() - 0.5) * float64(rand.Int31n(100))

	return Location{
		Lat:     loc.Lat + randomLat,
		Long:    loc.Long + randomLong,
		PlaceID: loc.PlaceID,
	}
}

func generateCoordinatesFarFromLocation(loc Location) Location {
	randomLatOffset := (rand.Float64() - 0.5) * float64(rand.Int31n(200))
	randomLongOffset := (rand.Float64() - 0.5) * float64(rand.Int31n(200))

	newLat := loc.Lat + randomLatOffset
	newLong := loc.Long + randomLongOffset

	if rand.Intn(2) == 1 {
		newLat, newLong = newLong, newLat
	}

	return Location{
		Lat:     newLat,
		Long:    newLong,
		PlaceID: loc.PlaceID,
	}
}

func generateRandomRides(number int8, local Location) []RideObject {
	rides := []RideObject{}
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

	seedOrigin := Location{Lat: -11.0522, Long: 34.2437, PlaceID: "LA"}

	initRide := createRide("2023-10-27T10:00:00Z", generateCoordinatesCloseToLocation(seedOrigin), generateCoordinatesFarFromLocation(seedOrigin), 4)

	pod := createPod(initRide)

	randomRides := generateRandomRides(10, seedOrigin)

	podAndRides := map[string]any{"randomRides": randomRides, "pod": pod}

	helpers.UseHandler(podAndRides)

	http.ListenAndServe(":8080", nil)

}
