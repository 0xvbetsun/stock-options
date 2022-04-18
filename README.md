
<div align="center">
	<h1><img alt="stock-options logo" src="/logo.png" height="300" /><br />
		Golang Stock Options Calculator
	</h1>
</div>


![Build Status](https://github.com/vbetsun/stock-options/workflows/CI/badge.svg)
[![coverage](https://codecov.io/gh/vbetsun/stock-options/branch/master/graph/badge.svg)](https://codecov.io/gh/vbetsun/stock-options)
[![GoReport](https://goreportcard.com/badge/github.com/vbetsun/stock-options)](https://goreportcard.com/report/github.com/vbetsun/stock-options)
![license](https://img.shields.io/github/license/vbetsun/stock-options)
[![GoDoc](https://pkg.go.dev/badge/github.com/vbetsun/stock-options)](https://pkg.go.dev/github.com/vbetsun/stock-options)

## Install

```sh
# Go 1.16+
go install github.com/vbetsun/stock-options@latest

# Go version < 1.16
go get -u github.com/vbetsun/stock-options 
```

## Usage

```golang
package main

import (
	"fmt"

	so "github.com/vbetsun/stock-options"
)

func main() {
	var strike, stock, premium float64 = 50, 70, 10
	bep, err := so.BreakEvenPoint(so.Call, strike, premium)
	if err != nil {
		// handle error
	}
	fmt.Printf("%.2f", bep)
	// Output: 60.00
	payoff, err := so.PayoffFromBuying(so.Call, strike, stock, premium)
	if err != nil {
		// handle error
	}
	fmt.Printf("%.2f", payoff)
	// Output: 10.00
}
```

## License

Golang Stock Options Calculator is provided under the [MIT License](LICENSE)