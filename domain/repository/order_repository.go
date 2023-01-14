package repository

import "github.com/GustavoCarvalho/web-store/domain/entity"

type OrderRepository interface {
	save(order entity.Order)
	count() int64
}
