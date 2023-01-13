package entity

type OrderRepository interface {
	save(order Order)
	count() int64
}
