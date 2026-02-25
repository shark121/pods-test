package testing

import (
	"math"
	"math/rand"

	"github.com/google/uuid"
	types "github.com/shark121/pods-test/types"
)

func CreateUser(name string, age int16, defaultLocation types.Location, dob string) types.User {
	return types.User{Name: name, Age: age, DefaultLocation: defaultLocation, Dob: dob}
}

func GetDirection(origin types.Location, destination types.Location) float64 {
	// fmt.Println(origin, destination)
	return 0.000
}

func GetDistance(origin types.Location, destination types.Location) float64 {
	// fmt.Println(origin, destination)
	return 0.000
}

func CreateRide(rideTime string, origin types.Location, destination types.Location, rideCapacity int8) types.RideObject {
	rideStatus := "pending"
	rideID := uuid.New().String()
	direction := GetDirection(origin, destination)
	distance := GetDistance(origin, destination)

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

func GenerateCoordinatesCloseToLocation(loc types.Location) types.Location {
	randomLat := (rand.Float64() - 0.5) * float64(rand.Int31n(100))
	randomLng := (rand.Float64() - 0.5) * float64(rand.Int31n(100))

	return types.Location{
		Lat: loc.Lat + randomLat,
		Lng: loc.Lng + randomLng,
		// PlaceID: loc.PlaceID,
	}
}

func GenerateCoordinatesFarFromLocation(loc types.Location) types.Location {

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

func GenerateRandomRides(number int8, local types.Location) []types.RideObject {
	rides := []types.RideObject{}
	for range number {
		origin := GenerateCoordinatesFarFromLocation(local)
		destination := GenerateCoordinatesFarFromLocation(local)

		if rand.Intn(2) == 1 {
			origin, destination = destination, origin
		}

		ride := CreateRide("2023-10-27T10:00:00Z", origin, destination, 4)
		rides = append(rides, ride)
	}
	return rides
}
