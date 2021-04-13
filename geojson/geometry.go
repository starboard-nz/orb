package geojson

import (
	"encoding/json"
	"errors"

	"github.com/xerra-eo/orb"
)

// ErrInvalidGeometry will be returned if a the json of the geometry is invalid.
var ErrInvalidGeometry = errors.New("geojson: invalid geometry")

// A Geometry matches the structure of a GeoJSON Geometry.
type Geometry struct {
	Type        string       `json:"type"`
	Coordinates orb.Geometry `json:"coordinates,omitempty"`
	Geometries  []*Geometry  `json:"geometries,omitempty"`
}

// NewGeometry will create a Geometry object but will convert
// the input into a GoeJSON geometry. For example, it will convert
// Rings and Bounds into Polygons.
func NewGeometry(g orb.Geometry) *Geometry {
	jg := &Geometry{}
	switch g := g.(type) {
	case orb.Ring:
		jg.Coordinates = orb.Polygon{g}
	case orb.Bound:
		jg.Coordinates = g.ToPolygon()
	case orb.Collection:
		for _, c := range g {
			jg.Geometries = append(jg.Geometries, NewGeometry(c))
		}
		jg.Type = g.GeoJSONType()
	default:
		jg.Coordinates = g
	}

	if jg.Coordinates != nil {
		jg.Type = jg.Coordinates.GeoJSONType()
	}
	return jg
}

// Geometry returns the orb.Geometry for the geojson Geometry.
// This will convert the "Geometries" into a orb.Collection if applicable.
func (g Geometry) Geometry() orb.Geometry {
	if g.Coordinates != nil {
		return g.Coordinates
	}

	c := make(orb.Collection, 0, len(g.Geometries))
	for _, geom := range g.Geometries {
		c = append(c, geom.Geometry())
	}
	return c
}

// MarshalJSON will marshal the geometry into the correct json structure.
func (g Geometry) MarshalJSON() ([]byte, error) {
	if g.Coordinates == nil && len(g.Geometries) == 0 {
		return []byte(`null`), nil
	}

	ng := &jsonGeometryMarshall{}
	switch g := g.Coordinates.(type) {
	case orb.Ring:
		ng.Coordinates = orb.Polygon{g}
	case orb.Bound:
		ng.Coordinates = g.ToPolygon()
	case orb.Collection:
		ng.Geometries = make([]*Geometry, 0, len(g))
		for _, c := range g {
			ng.Geometries = append(ng.Geometries, NewGeometry(c))
		}
		ng.Type = g.GeoJSONType()
	default:
		ng.Coordinates = g
	}

	if ng.Coordinates != nil {
		ng.Type = ng.Coordinates.GeoJSONType()
	}

	if len(g.Geometries) > 0 {
		ng.Geometries = g.Geometries
		ng.Type = orb.Collection{}.GeoJSONType()
	}
	return json.Marshal(ng)
}

// UnmarshalGeometry decodes the data into a GeoJSON feature.
// Alternately one can call json.Unmarshal(g) directly for the same result.
func UnmarshalGeometry(data []byte) (*Geometry, error) {
	g := &Geometry{}
	err := json.Unmarshal(data, g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// UnmarshalJSON will unmarshal the correct geometry from the json structure.
func (g *Geometry) UnmarshalJSON(data []byte) error {
	jg := &jsonGeometry{}
	err := json.Unmarshal(data, jg)
	if err != nil {
		return err
	}

	switch jg.Type {
	case "Point":
		p := orb.Point{}
		err = json.Unmarshal(jg.Coordinates, &p)
		g.Coordinates = p
	case "MultiPoint":
		mp := orb.MultiPoint{}
		err = json.Unmarshal(jg.Coordinates, &mp)
		g.Coordinates = mp
	case "LineString":
		ls := orb.LineString{}
		err = json.Unmarshal(jg.Coordinates, &ls)
		g.Coordinates = ls
	case "MultiLineString":
		mls := orb.MultiLineString{}
		err = json.Unmarshal(jg.Coordinates, &mls)
		g.Coordinates = mls
	case "Polygon":
		p := orb.Polygon{}
		err = json.Unmarshal(jg.Coordinates, &p)
		g.Coordinates = p
	case "MultiPolygon":
		mp := orb.MultiPolygon{}
		err = json.Unmarshal(jg.Coordinates, &mp)
		g.Coordinates = mp
	case "GeometryCollection":
		g.Geometries = jg.Geometries
	default:
		return ErrInvalidGeometry
	}

	g.Type = g.Geometry().GeoJSONType()

	return nil
}

// A Point is a helper type that will marshal to/from a GeoJSON Point geometry.
type Point orb.Point

// Geometry will return the orb.Geometry version of the data.
func (p Point) Geometry() orb.Geometry {
	return orb.Point(p)
}

// MarshalJSON will convert the Point into a GeoJSON Point geometry.
func (p Point) MarshalJSON() ([]byte, error) {
	return json.Marshal(Geometry{Coordinates: orb.Point(p)})
}

// UnmarshalJSON will unmarshal the GeoJSON Point geometry.
func (p *Point) UnmarshalJSON(data []byte) error {
	g := &Geometry{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}

	point, ok := g.Coordinates.(orb.Point)
	if !ok {
		return errors.New("geojson: not a Point type")
	}

	*p = Point(point)
	return nil
}

// A MultiPoint is a helper type that will marshal to/from a GeoJSON MultiPoint geometry.
type MultiPoint orb.MultiPoint

// Geometry will return the orb.Geometry version of the data.
func (mp MultiPoint) Geometry() orb.Geometry {
	return orb.MultiPoint(mp)
}

// MarshalJSON will convert the MultiPoint into a GeoJSON MultiPoint geometry.
func (mp MultiPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(Geometry{Coordinates: orb.MultiPoint(mp)})
}

// UnmarshalJSON will unmarshal the GeoJSON MultiPoint geometry.
func (mp *MultiPoint) UnmarshalJSON(data []byte) error {
	g := &Geometry{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}

	multiPoint, ok := g.Coordinates.(orb.MultiPoint)
	if !ok {
		return errors.New("geojson: not a MultiPoint type")
	}

	*mp = MultiPoint(multiPoint)
	return nil
}

// A LineString is a helper type that will marshal to/from a GeoJSON LineString geometry.
type LineString orb.LineString

// Geometry will return the orb.Geometry version of the data.
func (ls LineString) Geometry() orb.Geometry {
	return orb.LineString(ls)
}

// MarshalJSON will convert the LineString into a GeoJSON LineString geometry.
func (ls LineString) MarshalJSON() ([]byte, error) {
	return json.Marshal(Geometry{Coordinates: orb.LineString(ls)})
}

// UnmarshalJSON will unmarshal the GeoJSON MultiPoint geometry.
func (ls *LineString) UnmarshalJSON(data []byte) error {
	g := &Geometry{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}

	lineString, ok := g.Coordinates.(orb.LineString)
	if !ok {
		return errors.New("geojson: not a LineString type")
	}

	*ls = LineString(lineString)
	return nil
}

// A MultiLineString is a helper type that will marshal to/from a GeoJSON MultiLineString geometry.
type MultiLineString orb.MultiLineString

// Geometry will return the orb.Geometry version of the data.
func (mls MultiLineString) Geometry() orb.Geometry {
	return orb.MultiLineString(mls)
}

// MarshalJSON will convert the MultiLineString into a GeoJSON MultiLineString geometry.
func (mls MultiLineString) MarshalJSON() ([]byte, error) {
	return json.Marshal(Geometry{Coordinates: orb.MultiLineString(mls)})
}

// UnmarshalJSON will unmarshal the GeoJSON MultiPoint geometry.
func (mls *MultiLineString) UnmarshalJSON(data []byte) error {
	g := &Geometry{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}

	multilineString, ok := g.Coordinates.(orb.MultiLineString)
	if !ok {
		return errors.New("geojson: not a MultiLineString type")
	}

	*mls = MultiLineString(multilineString)
	return nil
}

// A Polygon is a helper type that will marshal to/from a GeoJSON Polygon geometry.
type Polygon orb.Polygon

// Geometry will return the orb.Geometry version of the data.
func (p Polygon) Geometry() orb.Geometry {
	return orb.Polygon(p)
}

// MarshalJSON will convert the Polygon into a GeoJSON Polygon geometry.
func (p Polygon) MarshalJSON() ([]byte, error) {
	return json.Marshal(Geometry{Coordinates: orb.Polygon(p)})
}

// UnmarshalJSON will unmarshal the GeoJSON Polygon geometry.
func (p *Polygon) UnmarshalJSON(data []byte) error {
	g := &Geometry{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}

	polygon, ok := g.Coordinates.(orb.Polygon)
	if !ok {
		return errors.New("geojson: not a Polygon type")
	}

	*p = Polygon(polygon)
	return nil
}

// A MultiPolygon is a helper type that will marshal to/from a GeoJSON MultiPolygon geometry.
type MultiPolygon orb.MultiPolygon

// Geometry will return the orb.Geometry version of the data.
func (mp MultiPolygon) Geometry() orb.Geometry {
	return orb.MultiPolygon(mp)
}

// MarshalJSON will convert the MultiPolygon into a GeoJSON MultiPolygon geometry.
func (mp MultiPolygon) MarshalJSON() ([]byte, error) {
	return json.Marshal(Geometry{Coordinates: orb.MultiPolygon(mp)})
}

// UnmarshalJSON will unmarshal the GeoJSON MultiPolygon geometry.
func (mp *MultiPolygon) UnmarshalJSON(data []byte) error {
	g := &Geometry{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}

	multiPolygon, ok := g.Coordinates.(orb.MultiPolygon)
	if !ok {
		return errors.New("geojson: not a MultiPolygon type")
	}

	*mp = MultiPolygon(multiPolygon)
	return nil
}

type jsonGeometry struct {
	Type        string           `json:"type"`
	Coordinates nocopyRawMessage `json:"coordinates"`
	Geometries  []*Geometry      `json:"geometries,omitempty"`
}

type jsonGeometryMarshall struct {
	Type        string       `json:"type"`
	Coordinates orb.Geometry `json:"coordinates,omitempty"`
	Geometries  []*Geometry  `json:"geometries,omitempty"`
}

type nocopyRawMessage []byte

func (m *nocopyRawMessage) UnmarshalJSON(data []byte) error {
	*m = data
	return nil
}
