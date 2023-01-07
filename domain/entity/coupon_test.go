package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewCoupon(t *testing.T) {
	c := Coupon{
		Name:  "VALE20",
		Value: 20,
	}
	assert.Equal(t, c.Value, int64(20))
	assert.Equal(t, c.Name, "VALE20")
	assert.Equal(t, CouponIsExpired(c, time.Now()), false)
}

func TestCreateNewCouponExpired(t *testing.T) {
	c := Coupon{
		Name:        "VALE20",
		Value:       20,
		ExpiredDate: time.Date(2022, 12, 10, 23, 0, 0, 0, time.UTC),
	}
	assert.Equal(t, CouponIsExpired(c, time.Now()), true)
}
