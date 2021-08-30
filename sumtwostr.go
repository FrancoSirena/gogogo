package main

import (
	"fmt"
	"math/big"
)

func addStrings(num1 string, num2 string) string {
	n1_conv, _ := new(big.Int).SetString(num1, 0)
	n2_conv, _ := new(big.Int).SetString(num2, 0)
	var r big.Int
	return r.Add(n1_conv, n2_conv).String()
}

func main() {
	fmt.Println(addStrings("11", "123"))
}
