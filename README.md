# go-upbit

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hexoul/go-upbit/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/hexoul/go-upbit)](https://goreportcard.com/report/github.com/hexoul/go-upbit)
[![GoDoc](https://godoc.org/github.com/hexoul/go-upbit?status.svg)](https://godoc.org/github.com/hexoul/go-upbit)

> Upbit API Client written in Golang

## Install

`go get -u github.com/hexoul/go-upbit`

## Getting started

- As library, start from `upbit.GetInstanceWithKey('YOUR_ACCESS_KEY', 'YOUR_SECRET_KEY')`
- As program, start from `upbit.GetInstance()` after executing `go run -upbit:accesskey=[YOUR_ACCESS_KEY] -upbit:secretkey=[YOUR_SECRET_KEY]`

```go
package main

import (
    "fmt"

    upbit "github.com/hexoul/go-upbit"
)

func init() {
    upbit.GetInstanceWithKey("YOUR_ACCESS_KEY", "YOUR_SECRET_KEY")
}

func main() {
    if ret, err := upbit.GetInstance().Accounts(); err == nil {
		for _, v := range ret {
			fmt.Println(v.Currency, v.Balance, v.Locked)
		}
	}
}
```

## Features

| Type           | Method | Endpoint                               | Done |
|----------------|--------|----------------------------------------|------|
| Asset | GET | /v1/accounts | ✔ |
| Order | GET | /v1/orders/chance | - |
| Order | GET | /v1/orders | - |
| Order | GET | /v1/order | - |
| Order | POST | /v1/order | - |
| Order | DELETE | /v1/order | - |

## Reference

[Upbit API](https://docs.upbit.com/v1.0.2/reference)