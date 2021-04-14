package orb

import (
	"fmt"
	"math"
)

const (
	EAST_MAX  = 180.
	WEST_MAX  = -180.
	NORTH_MAX = 90.
	SOUTH_MAX = -90.
)

var emptyBound = Bound{Min: Point{1, 1}, Max: Point{-1, -1}}

// MultiBounds is a slice of PolyBounds that is used to contain
// bounds of each Polygon within a MultiPolygon.
type MultiBounds []PolyBounds

// PolyBounds is a slice that represent the bounds of each
// geofence of a polygon, the first being the polygon's exterior,
// and the rest being each hole within the polygon.
type PolyBounds []Bound

// A Bound represents a closed box or rectangle.
// To create a bound with two points you can do something like:
//	orb.MultiPoint{p1, p2}.Bound()
type Bound struct {
	Min, Max Point
}

// A bound representation useful for rendering bounding boxes.
type NWSEBound struct {
	NW, SE Point
}

// GeoJSONType returns the GeoJSON type for the object.
func (b Bound) GeoJSONType() string {
	return "Polygon"
}

// Dimensions returns 2 because a Bound is a 2d object.
func (b Bound) Dimensions() int {
	return 2
}

// ToPolygon converts the bound into a Polygon object.
func (b Bound) ToPolygon() Polygon {
	return Polygon{b.ToRing()}
}

// ToRing converts the bound into a loop defined
// by the boundary of the box.
func (b Bound) ToRing() Ring {
	return Ring{
		b.Min,
		Point{b.Max[0], b.Min[1]},
		b.Max,
		Point{b.Min[0], b.Max[1]},
		b.Min,
	}
}

// Extend grows the bound to include the new point.
func (b Bound) Extend(point Point) Bound {
	// already included, no big deal
	if b.Contains(point) {
		return b
	}

	return Bound{
		Min: Point{
			math.Min(b.Min[0], point[0]),
			math.Min(b.Min[1], point[1]),
		},
		Max: Point{
			math.Max(b.Max[0], point[0]),
			math.Max(b.Max[1], point[1]),
		},
	}
}

// Union extends this bound to contain the union of this and the given bound.
func (b Bound) Union(other Bound) Bound {
	if other.IsEmpty() {
		return b
	}

	b = b.Extend(other.Min)
	b = b.Extend(other.Max)
	b = b.Extend(other.LeftTop())
	b = b.Extend(other.RightBottom())

	return b
}

// Contains determines if the point is within the bound.
// Points on the boundary are considered within.
func (b Bound) Contains(point Point) bool {
	if point[1] < b.Min[1] || b.Max[1] < point[1] {
		return false
	}

	if point[0] < b.Min[0] || b.Max[0] < point[0] {
		return false
	}

	return true
}

// Intersects determines if two bounds intersect.
// Returns true if they are touching.
func (b Bound) Intersects(bound Bound) bool {
	if (b.Max[0] < bound.Min[0]) ||
		(b.Min[0] > bound.Max[0]) ||
		(b.Max[1] < bound.Min[1]) ||
		(b.Min[1] > bound.Max[1]) {
		return false
	}

	return true
}

// Pad extends the bound in all directions by the given value.
func (b Bound) Pad(d float64) Bound {
	b.Min[0] -= d
	b.Min[1] -= d

	b.Max[0] += d
	b.Max[1] += d

	return b
}

// Center returns the center of the bounds by "averaging" the x and y coords.
func (b Bound) Center() Point {
	return Point{
		(b.Min[0] + b.Max[0]) / 2.0,
		(b.Min[1] + b.Max[1]) / 2.0,
	}
}

// Top returns the top of the bound.
func (b Bound) Top() float64 {
	return b.Max[1]
}

// Bottom returns the bottom of the bound.
func (b Bound) Bottom() float64 {
	return b.Min[1]
}

// Right returns the right of the bound.
func (b Bound) Right() float64 {
	return b.Max[0]
}

// Left returns the left of the bound.
func (b Bound) Left() float64 {
	return b.Min[0]
}

// LeftTop returns the upper left point of the bound.
func (b Bound) LeftTop() Point {
	return Point{b.Left(), b.Top()}
}

// RightBottom return the lower right point of the bound.
func (b Bound) RightBottom() Point {
	return Point{b.Right(), b.Bottom()}
}

// IsEmpty returns true if it contains zero area or if
// it's in some malformed negative state where the left point is larger than the right.
// This can be caused by padding too much negative.
func (b Bound) IsEmpty() bool {
	return b.Min[0] > b.Max[0] || b.Min[1] > b.Max[1]
}

// IsZero return true if the bound just includes just null island.
func (b Bound) IsZero() bool {
	return b.Max == Point{} && b.Min == Point{}
}

// Bound returns the the same bound.
func (b Bound) Bound() Bound {
	return b
}

// Equal returns if two bounds are equal.
func (b Bound) Equal(c Bound) bool {
	return b.Min == c.Min && b.Max == c.Max
}

// PolyBounds computes bounds for a slice of Polygons and returns a slice of orb.Bounds of
// the same length as poly. The slice's first element bounds the exterior polgon and each proceeding
// element bounds each hole within the polygon.
func PolygonBounds(poly Polygon) PolyBounds {
	bounds := make(PolyBounds, len(poly))
	for i, ring := range poly {
		bounds[i] = ring.Bound()
	}
	return bounds
}

// MultiPolyBounds computes bounds for a MultiPolygon and returns a 2 dimensional slice of orb.Bounds,
// one slice for each Polygon of the MultiPolygon.
func MultiPolyBounds(mp MultiPolygon) MultiBounds {
	bounds := make(MultiBounds, len(mp))
	for i, poly := range mp {
		bounds[i] = PolygonBounds(poly)
	}
	return bounds
}

func AOIBounds() {

}

// CaptureExteriorBounds collects the exteroir bounds of each polygon from processed bounds
// of every polygon in a multi-polygon.
func CaptureExteriorBounds(bounds MultiBounds) []Bound {

	results := make([]Bound, len(bounds))
	for i, geoBound := range bounds {
		results[i] = geoBound[0]
	}
	return results
}

// MultiPolygonExteriorBounds collects the exterior bounds of each polygon in multi-polygon,
// ignoring the bounds of polygon holes.
func MultiPolygonExteriorBounds(mp MultiPolygon) []Bound {
	bounds := MultiPolyBounds(mp)
	return CaptureExteriorBounds(bounds)
}

// BoundToNWSE covnerts a bound from the traditional format to an alternative bound format,
// useful for rendering.
func BoundToNWSE(bound Bound) NWSEBound {

	nLat := bound.Max.Lat()
	eLon := bound.Max.Lon()
	sLat := bound.Min.Lat()
	wLon := bound.Min.Lon()

	return NWSEBound{
		NW: Point{
			nLat, wLon,
		},
		SE: Point{
			sLat, eLon,
		},
	}
}

// AntimeridianBounds finds the bounds of a multi-polygon which crosses the anti-meridian.
// An error occurs if the multi-polygon has not been found to be split across the anti-meridian.
func AntimeridianBounds(mp MultiPolygon) (*Bound, error) {

	crossedAnti := false
	eLon, wLon := WEST_MAX, EAST_MAX
	nLat, sLat := SOUTH_MAX, NORTH_MAX
	exBounds := MultiPolygonExteriorBounds(mp)

	for _, bound := range exBounds {
		if bound.Min.Lon() == WEST_MAX {
			crossedAnti = true
		}
		nLat = math.Max(nLat, bound.Max.Lat())
		sLat = math.Min(sLat, bound.Min.Lat())

		if !crossedAnti {
			wLon = math.Min(wLon, bound.Min.Lon())
		} else {
			eLon = math.Max(eLon, bound.Max.Lon())
		}
	}

	if !crossedAnti {
		return nil, fmt.Errorf("expected split MultiPolygon across anti-meridian but is not found to be")
	}

	antiBounds := Bound{
		Min: Point{eLon, sLat},
		Max: Point{wLon, nLat},
	}

	return &antiBounds, nil
}
