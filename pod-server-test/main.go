package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"sort"

	"github.com/google/uuid"
	"github.com/pod-server-test/helpers"

	"github.com/pod-server-test/directions"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"Lng"`
	// PlaceID string  `json:"placeId"`
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
	RideBearing  *float64 `json:"bearing"`
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
	randomLng := (rand.Float64() - 0.5) * float64(rand.Int31n(100))

	return Location{
		Lat: loc.Lat + randomLat,
		Lng: loc.Lng + randomLng,
		// PlaceID: loc.PlaceID,
	}
}

func generateCoordinatesFarFromLocation(loc Location) Location {
	randomLatOffset := (rand.Float64() - 0.5) * float64(rand.Int31n(200))
	randomLngOffset := (rand.Float64() - 0.5) * float64(rand.Int31n(200))

	newLat := loc.Lat + randomLatOffset
	newLng := loc.Lng + randomLngOffset

	if rand.Intn(2) == 1 {
		newLat, newLng = newLng, newLat
	}

	return Location{
		Lat: newLat,
		Lng: newLng,
		// PlaceID: loc.PlaceID,
	}
}

type Path struct {
	Origin      Location
	Destination Location
	myInt       interface {
		Area() int
	}
}

func getMidpoint(origin Location, destination Location) map[string]float64 {

	return map[string]float64{
		"x": float64((origin.Lng + destination.Lng)) / 2,
		"y": float64((origin.Lat + destination.Lat)) / 2,
	}
}

func calculateBearing(origin, destination Location) float64 {
	toRadians := func(deg float64) float64 { return (deg * math.Pi) / 180 }
	toDegrees := func(rad float64) float64 { return (rad * 180) / math.Pi }

	lat1 := toRadians(origin.Lat)
	lat2 := toRadians(destination.Lat)
	deltaLng := toRadians(destination.Lng - origin.Lng)

	y := math.Sin(deltaLng) * math.Cos(lat2)
	x := math.Cos(lat1)*math.Sin(lat2) - math.Sin(lat1)*math.Cos(lat2)*math.Cos(deltaLng)

	bearing := toDegrees(math.Atan2(y, x))
	return math.Mod(bearing+360, 360)
}

func calculateAngleBetweenRides(rideOneOrigin, rideOneDestination, rideTwoOrigin, rideTwoDestination Location) float64 {
	bearing1 := calculateBearing(rideOneOrigin, rideOneDestination)
	bearing2 := calculateBearing(rideTwoOrigin, rideTwoDestination)

	angleDifference := math.Abs(bearing1 - bearing2)
	if angleDifference > 180 {
		return 360 - angleDifference
	}
	return angleDifference
}

func rankRidesByProximityToPod(ridesArray []RideObject, pod Pod) []RideObject {
	podMidpoint := getMidpoint(pod.PodOrigin, pod.PodDestination)

	fmt.Println(podMidpoint)

	rankedRides := make([]RideObject, len(ridesArray))
	for i, ride := range ridesArray {
		rideMidpoint := getMidpoint(ride.Origin, ride.Destination)
		distance := math.Sqrt(math.Pow(rideMidpoint["x"]-podMidpoint["x"], 2) + math.Pow(rideMidpoint["y"]-podMidpoint["y"], 2))
		bearing := calculateAngleBetweenRides(ride.Origin, ride.Destination, pod.PodOrigin, pod.PodDestination)

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

	seedOrigin := Location{Lat: -11.0522, Lng: 34.2437}

	initRide := createRide("2023-10-27T10:00:00Z", generateCoordinatesCloseToLocation(seedOrigin), generateCoordinatesFarFromLocation(seedOrigin), 4)

	pod := createPod(initRide)

	randomRides := generateRandomRides(3, seedOrigin)

	podAndRides := map[string]any{"randomRides": randomRides, "pod": pod, "ranked": rankRidesByProximityToPod(randomRides, pod)}

	helpers.UseHandler(podAndRides)

	print("server started running")

	http.ListenAndServe(":5000", nil)

	directions.GetMapDirections(randomRides)

}
