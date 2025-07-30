package geojson

import (
	"encoding/json"
	"fmt"

	"github.com/starboard-nz/orb"
)

// A Feature corresponds to GeoJSON feature object
type Feature struct {
	ID         interface{}  `json:"id,omitempty"`
	Type       string       `json:"type"`
	BBox       BBox         `json:"bbox,omitempty"`
	Geometry   orb.Geometry `json:"geometry"`
	Properties Properties   `json:"properties"`

	// ExtraMembers can be used to encoded/decode extra key/members in
	// the base of the feature collection. Note that keys of "type", "bbox"
	// and "features" will not work as those are reserved by the GeoJSON spec.
	ExtraMembers Properties `json:"-"`
}

// NewFeature creates and initializes a GeoJSON feature given the required attributes.
func NewFeature(geometry orb.Geometry) *Feature {
	return &Feature{
		Type:       "Feature",
		Geometry:   geometry,
		Properties: make(map[string]interface{}),
	}
}

// Point implements the orb.Pointer interface so that Features can be used
// with quadtrees. The point returned is the center of the Bound of the geometry.
// To represent the geometry with another point you must create a wrapper type.
func (f *Feature) Point() orb.Point {
	return f.Geometry.Bound().Center()
}

var _ orb.Pointer = &Feature{}

// MarshalJSON converts the feature object into the proper JSON.
// It will handle the encoding of all the child geometries.
// Alternately one can call json.Marshal(f) directly for the same result.
func (f Feature) MarshalJSON() ([]byte, error) {
	var tmp map[string]interface{}
	if f.ExtraMembers != nil {
		tmp = f.ExtraMembers.Clone()
	} else {
		tmp = make(map[string]interface{}, 4)
	}

	if f.ID != nil {
		tmp["id"] = f.ID
	}

	tmp["type"] = "Feature"
	tmp["properties"] = f.Properties

	if f.BBox != nil {
		tmp["bbox"] = f.BBox
	}

	if f.Geometry != nil {
		tmp["geometry"] = NewGeometry(f.Geometry)
	} else {
		tmp["geometry"] = nil
	}

	if len(f.Properties) == 0 {
		tmp["properties"] = nil
	}

	return json.Marshal(tmp)
}

// UnmarshalFeature decodes the data into a GeoJSON feature.
// Alternately one can call json.Unmarshal(f) directly for the same result.
func UnmarshalFeature(data []byte) (*Feature, error) {
	f := &Feature{}
	err := f.UnmarshalJSON(data)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalJSON handles the correct unmarshalling of the data
// into the orb.Geometry types.
func (f *Feature) UnmarshalJSON(data []byte) error {
	tmp := make(map[string]nocopyRawMessage, 5)

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	jf := &Feature{}
	for key, value := range tmp {
		switch key {
		case "id":
			err := json.Unmarshal(value, &jf.ID)
			if err != nil {
				return fmt.Errorf("geojson: invalid id: %w", err)
			}
		case "type":
			err := json.Unmarshal(value, &jf.Type)
			if err != nil {
				return fmt.Errorf("geojson: invalid type: %w", err)
			}
		case "bbox":
			err := json.Unmarshal(value, &jf.BBox)
			if err != nil {
				return fmt.Errorf("geojson: invalid bbox: %w", err)
			}
		case "geometry":
			geom := &Geometry{}
			err := json.Unmarshal(value, geom)
			if err != nil {
				return err
			}
			if geom.Coordinates == nil && geom.Geometries == nil {
				return ErrInvalidGeometry
			}
			jf.Geometry = geom.Geometry()
		case "properties":
			err := json.Unmarshal(value, &jf.Properties)
			if err != nil {
				return fmt.Errorf("geojson: invalid properties: %w", err)
			}
		default:
			if jf.ExtraMembers == nil {
				jf.ExtraMembers = Properties{}
			}

			var val interface{}
			err := json.Unmarshal(value, &val)
			if err != nil {
				return err
			}
			jf.ExtraMembers[key] = val
		}
	}

	if jf.Type != "Feature" {
		return fmt.Errorf("geojson: not a feature: type=%s", jf.Type)
	}

	*f = Feature{
		ID:           jf.ID,
		Type:         jf.Type,
		Properties:   jf.Properties,
		BBox:         jf.BBox,
		Geometry:     jf.Geometry,
		ExtraMembers: jf.ExtraMembers,
	}

	return nil
}
