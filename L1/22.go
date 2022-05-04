package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("622246405745257275088696311157297823662689037", 10)
	b, _ := new(big.Int).SetString("254754826712765328764954672573567865976942151", 10)
	fmt.Printf("Quotient - %v\n", new(big.Int).Div(a, b))
	fmt.Printf("Product - %v\n", new(big.Int).Mul(a, b))
	fmt.Printf("Sum - %v\n", new(big.Int).Add(a, b))
	fmt.Printf("Difference - %v\n", new(big.Int).Sub(a, b))
}
