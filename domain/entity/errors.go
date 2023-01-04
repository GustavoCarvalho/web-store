package entity

import (
	"errors"
)

var (
	ErrNotPaid      = errors.New("order has not paid yet")
	ErrInvalidValue = errors.New("invalid value")
)
