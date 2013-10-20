# *fpd.Decimal [![Build Status](https://travis-ci.org/oguzbilgic/fpd.png?branch=master)](https://travis-ci.org/oguzbilgic/fpd)

Package implements fixed-point decimal 

## Usage

```go
package main

import "github.com/oguzbilgic/fpd"

func main() {
	// Buy price of the security: $136.02
	buyPrice := fpd.New(13602000, -5)

	// Sell price of the security: $137.699
	sellPrice := fpd.New(13769900, -5)

	// Volume traded: 0.01
	volume := fpd.New(1000000, -8)

	// Trade fee percentage: 0.6%
	feePer := fpd.New(6, -3)

	buyCost := buyPrice.Mul(volume)
	buyFee := buyPrice.Mul(volume).Mul(feePer)
	sellRevenue := sellPrice.Mul(volume)
	sellFee := sellPrice.Mul(volume).Mul(feePer)

	// Initall account balance: $2.00000
	balance := fpd.New(200000, -5)

	balance = balance.Sub(buyCost)
	balance = balance.Sub(buyFee)
	balance = balance.Add(sellRevenue)
	balance = balance.Sub(sellFee)

	// Final balance
	fmt.Println(balance)
	// Did this trade turn into profit? :)
}
```

## Documentation

http://godoc.org/github.com/oguzbilgic/fpd

## License

The MIT License (MIT)
