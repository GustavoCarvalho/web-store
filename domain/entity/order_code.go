package entity

import (
	"fmt"
	"time"
)

type OrderCode struct {
	Sequence int64
	Date     time.Time
	Value    string
}

func NewOrderCode(date time.Time, sequence int64) OrderCode {
	code := fmt.Sprintf("%d", date.Year()) + fmt.Sprintf("%08d", 1)
	orderCode := OrderCode{
		Sequence: sequence,
		Date:     date,
		Value:    code,
	}
	return orderCode
}
