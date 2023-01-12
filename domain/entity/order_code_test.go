package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrderCode(t *testing.T) {
	date := time.Date(2021, 12, 10, 23, 0, 0, 0, time.UTC)
	sequence := 1
	o := NewOrderCode(date, int64(sequence))
	code := o.Value
	assert.Equal(t, code, "202100000001")
}
