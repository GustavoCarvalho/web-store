package entity

type Order struct {
	Cpf    Cpf
	Items  []Item
	Coupon Coupon
}

func NewOrder(cpf Cpf, items []Item, coupom Coupon) (*Order, error) {
	o := Order{
		Cpf:    cpf,
		Items:  items,
		Coupon: coupom,
	}
	return &o, nil
}

func OrderGetTotal(o *Order) float64 {
	var sum float64
	for i := 0; i < len(o.Items); i++ {
		sum = sum + o.Items[i].Price
	}
	return sum - sum*(float64(o.Coupon.Value)/100)
}
