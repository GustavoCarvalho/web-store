package entity

type Dimension struct {
	Width   float64
	Height  float64
	Lenght  float64
	Density float64
	Weight  float64
}

func DimensionGetVolume(d *Dimension) float64 {
	return d.Height / 100 * d.Width / 100 * d.Lenght / 100
}

func DimensionGetDensity(d *Dimension) float64 {
	if d.Weight != 0 {
		return float64(d.Weight / DimensionGetVolume(d))
	}
	return 0
}
