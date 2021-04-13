package geo

import (
	"github.com/xerra-eo/orb"
	"github.com/xerra-eo/orb/internal/length"
)

// Length returns the length of the boundary of the geometry
// using the geo distance function.
func Length(g orb.Geometry) float64 {
	return length.Length(g, Distance)
}

// LengthHaversign returns the length of the boundary of the geometry
// using the geo haversine formula
func LengthHaversign(g orb.Geometry) float64 {
	return length.Length(g, DistanceHaversine)
}
