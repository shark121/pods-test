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

func CalculateAngleBetweenRides(rideOneOrigin, rideOneDestination, rideTwoOrigin, rideTwoDestination types.Location) float64 {
	bearing1 := CalculateBearing(rideOneOrigin, rideOneDestination)
	bearing2 := CalculateBearing(rideTwoOrigin, rideTwoDestination)

	angleDifference := math.Abs(bearing1 - bearing2)
	if angleDifference > 180 {
		return 360 - angleDifference
	}
	return angleDifference
}
