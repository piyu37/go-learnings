package main

import "fmt"

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	if len(array) == 0 {
		return -1
	}

	return sumOfArray(array, 1)
}

func sumOfArray(array []interface{}, mul int) int {
	total := 0

	for _, data := range array {
		switch v := data.(type) {
		case int:
			total += v

		case SpecialArray:
			total += sumOfArray(v, mul+1)

		default:
			return total * mul
		}
	}

	return total * mul
}

// https://github.com/lee-hen/Algoexpert/tree/master/easy/08_product_sum
func productSumMain() {
	arr := SpecialArray{5, 2, SpecialArray{7, -1}, 3, SpecialArray{6, SpecialArray{-13, 8}, 4}}

	fmt.Println(ProductSum(arr))
}
