package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("The result is a Complex Number")
	}
	return math.Sqrt(num), nil
}

// output:
// The result is a Complex Number
// 16
func handlingErrors() {
	output, err := Sqrt(-4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output)
	}

	output, err = Sqrt(256)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output)
	}
}
