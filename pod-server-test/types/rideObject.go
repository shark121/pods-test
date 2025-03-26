package types

type location struct {
	lat     float64
	long    float64
	placeId string
}

type RideObject struct {
	rideId       string
	rideTime     string
	rideStatus   string
	origin       location
	destination  location
	rideCapacity int16
	direction    float64
	rideDistance float64
}
