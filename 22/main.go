package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 21) // a = 2^21
	b := big.NewInt(1 << 22) // b = 2^22

	fmt.Printf("1st number: %s, 2nd number: %s\n\n", a.String(), b.String())

	res := big.NewInt(0)
	res.Mul(a, b)
	fmt.Println("multiplex:", res.String())

	res.Div(a, b)
	fmt.Println("division:", res.String())

	res.Add(a, b)
	fmt.Println("sum:", res.String())

	res.Sub(a, b)
	fmt.Println("subtract:", res.String())
}
