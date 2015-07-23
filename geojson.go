package geojson

import (
	"encoding/json"
	"io"
)

// Point is a 2 element array of floats with order longitude, latitude
type Point [2]float64

// Line is a series of Points
type Line []Point

// Polygon is a series of Lines, the first being the exterior ring of the polygon and the rest are holes
type Polygon []Line

// Geometry is a generalized way to store GeoJSON geometry. When
// unmarshalling, if the geometry is a Point it will be stored in the Point
// member. If a LineString then it will be stored in the LineString member,
// and so on for Polygon and MultiPolygon
type Geometry struct {
	Type         string
	Point        Point
	LineString   Line
	Polygon      Polygon
	MultiPolygon []Polygon
}

// Container is a generalized way to store GeoJSON objects. When
// unmarshalling, a FeatureCollection will have its Features stored in the
// Features member. Otherwise, the described geometry will be stored in the
// Geometry member.
type Container struct {
	Type       string
	Properties map[string]interface{}
	Geometry   Geometry
	Features   []Container
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface
func (c *Container) UnmarshalJSON(data []byte) error {
	uc, err := unmarshalBytes(data)
	if err != nil {
		return err
	}

	*c = uc.toContainer()
	return nil
}

// Unmarshal unmarshals GeoJSON from a Reader into a Container and returns the Container
func Unmarshal(r io.Reader) (*Container, error) {
	var c Container
	dec := json.NewDecoder(r)
	if err := dec.Decode(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
