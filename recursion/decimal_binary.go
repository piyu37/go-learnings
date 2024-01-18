package main

import "fmt"

func convertDecimalToBinary(v int) int {
	if v < 1 {
		return v
	}

	mod := v % 2

	return 10*convertDecimalToBinary(v/2) + mod
}

// convert decimal to binary
func decimalBinary() {
	fmt.Println(convertDecimalToBinary(7))
}
