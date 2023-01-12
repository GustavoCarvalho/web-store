package entity

import "time"

type Order struct {
	Cpf       Cpf
	Items     []Item
	Coupon    Coupon
	IssueDate time.Time
	OrderCode string
	Sequence  int64
}

func NewOrder(cpf Cpf, items []Item, coupom Coupon, issueDate time.Time, sequence int64) (*Order, error) {
	o := Order{
		Cpf:       cpf,
		Items:     items,
		Coupon:    coupom,
		IssueDate: issueDate,
		OrderCode: NewOrderCode(issueDate, sequence).Value,
	}
	return &o, nil
}

func OrderGetTotal(o *Order) float64 {
	var sum float64
	for i := 0; i < len(o.Items); i++ {
		sum = sum + o.Items[i].Price
	}
	if CouponIsExpired(o.Coupon, time.Now()) {
		return sum
	}
	return sum - sum*(float64(o.Coupon.Value)/100)
}
