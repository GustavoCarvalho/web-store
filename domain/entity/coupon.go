package entity

import "time"

type Coupon struct {
	Name        string
	Value       int64
	ExpiredDate time.Time
}

type Now func() time.Time

func CouponIsExpired(c Coupon, today time.Time) bool {
	if c.ExpiredDate.UnixMilli() < 0 {
		return false
	}
	return c.ExpiredDate.UnixMilli() < today.UnixMilli()
}
