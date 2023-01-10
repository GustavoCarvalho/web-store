package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDimensionGetVolume(t *testing.T) {
	dimension := Dimension{
		Width:  100,
		Height: 30,
		Lenght: 10,
	}
	volume := DimensionGetVolume(&dimension)
	assert.Equal(t, volume, float64(0.03))
}
