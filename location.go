package location

import "math"

// p is the constant for converting degrees to radians
const p = math.Pi / 180

var (
	// EarthRadius is the radius of the earth in kilometers
	// change this if you want to use miles or going to another planet
	EarthRadius = 6371.0
)

// Location represents a location on earth
// Latitude and Longitude are in degrees
type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

// Distance calculates the distance between two locations using the Haversine formula
// returns the distance in kilometers
func Distance(startLocation, endLocation Location) float64 {
	return EarthRadius * 2 * math.Asin(math.Sqrt(0.5-math.Cos((endLocation.Latitude-startLocation.Latitude)*p)/2+
		math.Cos(startLocation.Latitude*p)*math.Cos(endLocation.Latitude*p)*
			(1-math.Cos((endLocation.Longitude-startLocation.Longitude)*p))/2))
}
