package tilecover

import (
	"testing"

	"github.com/starboard-nz/orb"
)

func TestGeometry(t *testing.T) {
	for _, g := range orb.AllGeometries {
		Geometry(g, 1)
	}
}
