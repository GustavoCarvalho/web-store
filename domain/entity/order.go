package entity

type Order struct {
	Cpf Cpf
}

func NewOrder(cpf Cpf) (*Order, error) {
	o := Order{
		Cpf: cpf,
	}
	return &o, nil
}
