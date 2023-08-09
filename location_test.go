package location

import (
	"testing"
)

func TestDistance(t *testing.T) {
	loc1 := Location{Latitude: 52.5200, Longitude: 13.4050} // Berlin, Germany
	loc2 := Location{Latitude: 48.8566, Longitude: 2.3522}  // Paris, France

	distance := Distance(loc1, loc2)
	t.Logf("Distance between Berlin and Paris: %f km", distance)
}
