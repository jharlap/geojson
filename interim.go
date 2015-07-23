package geojson

import (
	"encoding/json"
	"fmt"
)

type withCoords interface {
	getType() string
	getCoords() json.RawMessage
	clearCoords()
}

type interimGeometry struct {
	Type         string
	Coordinates  json.RawMessage
	Point        Point
	LineString   Line
	Polygon      Polygon
	MultiPolygon []Polygon
}

func (g *interimGeometry) getType() string {
	return g.Type
}

func (g *interimGeometry) getCoords() json.RawMessage {
	return g.Coordinates
}

func (g *interimGeometry) clearCoords() {
	g.Coordinates = nil
}

func unmarshalCoordinatesInto(w withCoords, g *interimGeometry) error {
	c := w.getCoords()
	if c == nil {
		return nil
	}

	var err error
	switch w.getType() {
	case "Point":
		err = json.Unmarshal(c, &g.Point)
	case "MultiPoint":
		err = json.Unmarshal(c, &g.LineString)
	case "LineString":
		err = json.Unmarshal(c, &g.LineString)
	case "Polygon":
		err = json.Unmarshal(c, &g.Polygon)
	case "MultiPolygon":
		err = json.Unmarshal(c, &g.MultiPolygon)
	default:
		err = fmt.Errorf("Unknown GeoJSON type: %s", w.getType())
	}

	if err != nil {
		return err
	}

	w.clearCoords()
	return nil
}

func (g *interimGeometry) unmarshalCoordinates() error {
	return unmarshalCoordinatesInto(g, g)
}

func (c *interimContainer) unmarshalCoordinates() error {
	return unmarshalCoordinatesInto(c, &c.Geometry)
}

func (c *interimContainer) unmarshal() error {
	var err error
	switch c.Type {
	case "Feature":
		err = c.Geometry.unmarshalCoordinates()
	case "FeatureCollection":
		for j := range c.Features {
			f := &c.Features[j]
			if err = f.Geometry.unmarshalCoordinates(); err != nil {
				fmt.Printf("Failed to convert %s: %s", f.Geometry.Type, err)
			}
		}
	default:
		err = c.unmarshalCoordinates()
	}
	return err
}

type interimContainer struct {
	Type        string
	Coordinates json.RawMessage
	Properties  map[string]interface{}
	Geometry    interimGeometry
	Features    []interimContainer
}

func (c *interimContainer) clearCoords() {
	c.Coordinates = nil
}

func (c *interimContainer) getCoords() json.RawMessage {
	return c.Coordinates
}

func (c *interimContainer) getType() string {
	return c.Type
}

func (c *interimContainer) toContainer() Container {
	r := Container{
		Type:       (*c).Type,
		Properties: (*c).Properties,
		Geometry: Geometry{
			Type:         (*c).Geometry.Type,
			Point:        (*c).Geometry.Point,
			LineString:   (*c).Geometry.LineString,
			Polygon:      (*c).Geometry.Polygon,
			MultiPolygon: (*c).Geometry.MultiPolygon,
		},
	}

	for _, f := range (*c).Features {
		r.Features = append(r.Features, f.toContainer())
	}

	return r
}

// unmarshalBytes unmarshals GeoJSON into an interimContainer and returns the interimContainer
func unmarshalBytes(d []byte) (*interimContainer, error) {
	var c interimContainer
	if err := json.Unmarshal(d, &c); err != nil {
		return nil, err
	}

	if err := c.unmarshal(); err != nil {
		return nil, err
	}

	return &c, nil
}
