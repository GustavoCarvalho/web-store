package infra

import "github.com/GustavoCarvalho/web-store/domain/entity"

func SaveOrder(order *entity.Order) []*entity.Order {
	var orders []*entity.Order
	orders = append(orders, order)
	return orders
}

func CountOrders(orders []*entity.Order) int {
	return len(orders)
}
