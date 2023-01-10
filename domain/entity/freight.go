package entity

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
	freight := volume * 1000 * (density / 100)
	return freight * float64(item.Quantity)
}
