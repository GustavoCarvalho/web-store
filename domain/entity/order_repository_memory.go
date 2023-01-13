package entity

func SaveOrder(order Order) []Order {
	var orders []Order
	orders = append(orders, order)
	return orders
}

func CountOrders(orders []Order) int {
	return len(orders)
}
