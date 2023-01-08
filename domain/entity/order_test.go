package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewOrderWithProducts(t *testing.T) {
	product1 := Product{
		Id:          "1",
		Name:        "Guitarra",
		Description: "Guitarra elétrica",
	}
	product2 := Product{
		Id:          "2",
		Name:        "Cabo P2",
		Description: "Cabo de som",
	}
	product3 := Product{
		Id:          "3",
		Name:        "Paleta",
		Description: "Paleta de acrílico",
	}
	i1 := Item{
		Id:       "idItem1",
		Price:    2000,
		Quantity: 1,
		Product:  product1,
	}
	i2 := Item{
		Id:       "idItem2",
		Price:    20,
		Quantity: 1,
		Product:  product2,
	}
	i3 := Item{
		Id:       "idItem3",
		Price:    10,
		Quantity: 1,
		Product:  product3,
	}
	items := []Item{i1, i2, i3}
	o, _ := fakeOrder(items)
	total := OrderGetTotal(o)
	assert.Equal(t, total, float64(2030))
}
func TestNewOrderWithProductsWithValidCoupon(t *testing.T) {
	product1 := Product{
		Id:          "1",
		Name:        "Guitarra",
		Description: "Guitarra elétrica",
	}
	product2 := Product{
		Id:          "2",
		Name:        "Cabo P2",
		Description: "Cabo de som",
	}
	product3 := Product{
		Id:          "3",
		Name:        "Paleta",
		Description: "Paleta de acrílico",
	}
	i1 := Item{
		Id:       "idItem1",
		Price:    2000,
		Quantity: 1,
		Product:  product1,
	}
	i2 := Item{
		Id:       "idItem2",
		Price:    20,
		Quantity: 1,
		Product:  product2,
	}
	i3 := Item{
		Id:       "idItem3",
		Price:    10,
		Quantity: 1,
		Product:  product3,
	}
	items := []Item{i1, i2, i3}
	coupon := Coupon{
		Name:  "VALE20",
		Value: 20,
	}
	o, _ := fakeOrderWithCoupom(items, coupon)
	total := OrderGetTotal(o)
	assert.Equal(t, total, float64(1624))
}

func TestNewOrderWithProductNotApplyExpiredCoupon(t *testing.T) {
	product1 := Product{
		Id:          "1",
		Name:        "Guitarra",
		Description: "Guitarra elétrica",
	}
	product2 := Product{
		Id:          "2",
		Name:        "Cabo P2",
		Description: "Cabo de som",
	}
	product3 := Product{
		Id:          "3",
		Name:        "Paleta",
		Description: "Paleta de acrílico",
	}
	i1 := Item{
		Id:       "idItem1",
		Price:    2000,
		Quantity: 1,
		Product:  product1,
	}
	i2 := Item{
		Id:       "idItem2",
		Price:    20,
		Quantity: 1,
		Product:  product2,
	}
	i3 := Item{
		Id:       "idItem3",
		Price:    10,
		Quantity: 1,
		Product:  product3,
	}
	items := []Item{i1, i2, i3}
	coupon := Coupon{
		Name:        "VALE20",
		Value:       20,
		ExpiredDate: time.Date(2022, 12, 10, 23, 0, 0, 0, time.UTC),
	}
	o, _ := fakeOrderWithCoupom(items, coupon)
	total := OrderGetTotal(o)
	assert.Equal(t, total, float64(2030))
}

func fakeOrder(items []Item) (*Order, error) {
	cpfString := "099.075.865-60"
	cpf, _ := NewCPF(cpfString)
	coupon := Coupon{
		Name:  "",
		Value: 0,
	}
	o, _ := NewOrder(cpf, items, coupon)
	return o, nil
}

func fakeOrderWithCoupom(items []Item, coupon Coupon) (*Order, error) {
	cpfString := "099.075.865-60"
	cpf, _ := NewCPF(cpfString)
	o, _ := NewOrder(cpf, items, coupon)
	return o, nil
}
