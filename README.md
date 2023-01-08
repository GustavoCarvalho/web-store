# web-store

## Step by step

### Root folder
```ssh
go mod init github.com/GustavoCarvalho/web-store
mkdir domain
mkdir entity
```

## DDD
[x] You must not place an order with a valid CPF/SSN
[x] You must place an order with 3 items (price, quantity, description)
[x] Must place an order with discount coupon
[x] Create valid and expired coupons
[x] Must not apply an expired discount coupon
[-] You must calculate the freight, considering the dimensions and weight of the items
[-] Must return the minimum value of 10 if the accumulated value is lower

## Dependencies
```ssh
go get regexp
go get strconv
go get github.com/stretchr/testify
```

## TDD
- Nao deve escrever nenuma linha se codigo sem que antes exista um codigo de falha.
- Voce nao deve escrever mais do que necessario para evidenciar uma fallha.
- Voce nao deve escrever maais codigo do que o suficiente  para fazer o codigo passar.

## References
https://go.dev/ref/spec
https://github1s.com/eyazici90/go-ddd/blob/HEAD/internal/domain/order_test.go
https://github1s.com/klassmann/cpfcnpj/blob/HEAD/cpf.go#L27
