package types

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"Lng"`
}

type RideObject struct {
	RideID       string    `json:"rideId"`
	RideTime     string    `json:"rideTime"`
	RideStatus   string    `json:"rideStatus"`
	Origin       Location  `json:"origin"`
	Destination  Location  `json:"destination"`
	RideCapacity int8      `json:"rideCapacity"`
	Direction    float64   `json:"direction"`
	RideDistance float64   `json:"rideDistance"`
	RideBearing  *float64  `json:"bearing"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Pod struct {
	PodOrigin      Location              `json:"origin"`
	PodDestination Location              `json:"destination"`
	PodCapacity    int8                  `json:"podCapacity"`
	PodStatus      string                `json:"podStatus"`
	PodID          string                `json:"podId"`
	PodRides       map[string]RideObject `json:"podRides"`
	Waypoints      []Location            `json:"waypoints"`
	PodDirection   float64               `json:"podDirection"`
	PodDistance    float64               `json:"podDistance"`
	CreatedAt      time.Time             `json:"createdAt"`
}

func (p *Pod) AddRide(ride RideObject) {
	p.PodRides[ride.RideID] = ride
}

func (p *Pod) RemoveRide(ride RideObject) {
	delete(p.PodRides, ride.RideID)
}

func CreatePod(ride RideObject) Pod {
	podStatus := "pending"
	podID := uuid.New().String()
	podRides := map[string]RideObject{ride.RideID: ride}
	podDirection := ride.Direction
	podDistance := ride.RideDistance

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
		CreatedAt:      time.Now(),
	}
}
