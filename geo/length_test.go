package geo

import (
	"testing"

	"github.com/xerra-eo/orb"
)

func TestLength(t *testing.T) {
	for _, g := range orb.AllGeometries {
		// should not panic with unsupported type
		Length(g)
	}
}

func TestLengthHaversign(t *testing.T) {
	for _, g := range orb.AllGeometries {
		// should not panic with unsupported type
		LengthHaversign(g)
	}
}
