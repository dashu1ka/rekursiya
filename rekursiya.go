package main

import (
	"fmt"
	"math/big"
)

func main() {
	f0 := big.NewInt(1)
	f1 := big.NewInt(3)

	a := []*big.Int{}

	for n := 0; len(a) < 40; n++ {
		var fn *big.Int
		if n == 0 {
			fn = new(big.Int).Set(f0)
		} else if n == 1 {
			fn = new(big.Int).Set(f1)
		} else {
			five := big.NewInt(5)
			prom := new(big.Int).Mul(five, f1)
			fn = new(big.Int).Add(prom, f0)
			f0.Set(f1)
			f1.Set(fn)
		}
		if fn.Bit(0) == 1 {
			a = append(a, new(big.Int).Set(fn))
		}
	}
	fmt.Println(a[39])
}
