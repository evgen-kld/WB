package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetString("458743336571232454664456", 10)

	b := new(big.Int)
	b.SetString("55788963352546451346", 10)

	fmt.Println(big.NewInt(0).Add(a, b)) // a+b
	fmt.Println(big.NewInt(0).Mul(a, b)) // a*b
	fmt.Println(big.NewInt(0).Div(a, b)) // a/b
	fmt.Println(big.NewInt(0).Sub(a, b)) // a-b
}
