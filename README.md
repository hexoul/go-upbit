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
    upbit.GetInstanceWithKey("YOUR_API_KEY")
}

func main() {
}
```

## Features

| Type           | Endpoint                               | Done |
|----------------|----------------------------------------|------|
| Asset | /v1/accounts | - |
