package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	o, _ := fakeOrder()
	assert.NotNil(t, o)
}

func fakeOrder() (*Order, error) {
	cpfString := "099.075.865-60"
	cpf, _ := NewCPF(cpfString)
	o, _ := NewOrder(cpf)
	return o, nil
}
