package calc

import (
	"math"

	"github.com/pod-server-test/types"
)

func GetMidpoint(origin types.Location, destination types.Location) map[string]float64 {

	return map[string]float64{
		"x": float64((origin.Lng + destination.Lng)) / 2,
		"y": float64((origin.Lat + destination.Lat)) / 2,
	}
}

func CalculateBearing(origin, destination types.Location) float64 {

	lat1 := toRad(origin.Lat)
	lat2 := toRad(destination.Lat)
	deltaLng := toRad(destination.Lng - origin.Lng)

	y := math.Sin(deltaLng) * math.Cos(lat2)
	x := math.Cos(lat1)*math.Sin(lat2) - math.Sin(lat1)*math.Cos(lat2)*math.Cos(deltaLng)

	bearing := toDeg(math.Atan2(y, x))
	return math.Mod(bearing+360, 360)
}

func DistanceBetweenTwoPoints(loc1, loc2 types.Location) float64 {
	// return value in km
	const R = 6371
	dLat := toRad(loc2.Lat - loc1.Lat)
	dLng := toRad(loc2.Lng - loc1.Lng)
	lat1 := toRad(loc1.Lat)
	lat2 := toRad(loc2.Lat)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dLng/2)*math.Sin(dLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func CalculateAngleBetweenRides(rideOneOrigin, rideOneDestination, rideTwoOrigin, rideTwoDestination types.Location) float64 {
	bearing1 := CalculateBearing(rideOneOrigin, rideOneDestination)
	bearing2 := CalculateBearing(rideTwoOrigin, rideTwoDestination)

	angleDifference := math.Abs(bearing1 - bearing2)
	if angleDifference > 180 {
		return 360 - angleDifference
	}
	return angleDifference
}

func toRad(deg float64) float64 {
	return deg * math.Pi / 180
}

func toDeg(rad float64) float64 {
	return (rad * 180) / math.Pi
}
