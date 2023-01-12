# web-store

## Step by step

### Root folder
```ssh
go mod init github.com/GustavoCarvalho/web-store
mkdir domain
mkdir entity
```

## DDD
- [x] You must not place an order with a valid CPF/SSN
- [x] You must place an order with 3 items (price, quantity, description)
- [x] Must place an order with discount coupon
- [x] Create valid and expired coupons
- [x] Must not apply an expired discount coupon
- [x] You must calculate the freight, considering the dimensions and weight of the items: volume*1000*(density/100)
- [x] Must return the minimum value of 10 if the accumulated value is lower
- [x] The order code is formed by AAAAPPPPPPPP where AAAA represents the year and PPPPPPPP represents a sequence of the order Implement a decoupled persistence mechanism using a database

## Dependencies
```ssh
go get regexp
go get strconv
go get github.com/stretchr/testify
```

## References
https://go.dev/ref/spec
https://github1s.com/eyazici90/go-ddd/blob/HEAD/internal/domain/order_test.go
https://github1s.com/klassmann/cpfcnpj/blob/HEAD/cpf.go#L27
