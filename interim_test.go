package geojson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterimUnmarshalBytesPoint(t *testing.T) {
	data := []byte(`{"type":"Point", "coordinates":[1.1,2]}`)

	result, err := unmarshalBytes(data)
	assert.NoError(t, err)

	expected := &interimContainer{
		Type: "Point",
		Geometry: interimGeometry{
			Point: Point{1.1, 2.0},
		},
	}
	assert.Equal(t, expected, result)
}

func TestInterimToContainer(t *testing.T) {
	i := &interimContainer{
		Type: "Point",
		Geometry: interimGeometry{
			Point: Point{1.1, 2.0},
		},
	}

	e := Container{
		Type: "Point",
		Geometry: Geometry{
			Point: Point{1.1, 2.0},
		},
	}

	c := i.toContainer()
	assert.Equal(t, e, c)
}
