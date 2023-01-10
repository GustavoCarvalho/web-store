package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductWithDimensions(t *testing.T) {
	product := Product{
		Id:          "1",
		Name:        "Guitarra",
		Description: "Guitarra elétrica",
		Dimension: Dimension{
			Height: 100,
			Width:  30,
			Lenght: 10,
			Weight: 3,
		},
	}
	volume := DimensionGetVolume(&product.Dimension)
	density := DimensionGetDensity(&product.Dimension)
	assert.Equal(t, volume, 0.03)
	assert.Equal(t, density, float64(100))
}
