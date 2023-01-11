package entity

var DISTANCE = 1000.00

func FreightGetTotal(items []Item) float64 {
	var sum float64
	for i := 0; i < len(items); i++ {
		sum = sum + CalculateFreight(items[i])
	}
	return sum
}

func CalculateFreight(item Item) float64 {
	volume := DimensionGetVolume(&item.Product.Dimension)
	density := DimensionGetDensity(&item.Product.Dimension)
	freight := volume * DISTANCE * (density / 100)
	if freight*float64(item.Quantity) < 10 {
		return 10.0
	}
	return freight * float64(item.Quantity)
}
