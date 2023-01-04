# web-store

## Step by step

### Root folder
```ssh
go mod init github.com/GustavoCarvalho/web-store
mkdir domain
mkdir entity
```


## DDD
[-] You must not place an order with a valid CPF/SSN
[-] You must place an order with 3 items (price, quantity, description)
[-] Must place an order with discount coupon

## Dependencies
```ssh
go get regexp
go get strconv
go get github.com/stretchr/testify
```


## References
https://github1s.com/eyazici90/go-ddd/blob/HEAD/internal/domain/order_test.go
https://github1s.com/klassmann/cpfcnpj/blob/HEAD/cpf.go#L27
