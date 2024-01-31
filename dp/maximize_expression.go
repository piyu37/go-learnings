package main

import (
	"fmt"
	"math"
)

func maximizeExpression(arr []int) int {
	if len(arr) < 4 {
		return 0
	}

	aMaxArr := []int{arr[0]}
	aMinBMaxArr := []int{math.MinInt}
	aMinBPlusCMaxArr := []int{math.MinInt, math.MinInt}
	aMinBPlusCMinDMaxArr := []int{math.MinInt, math.MinInt, math.MinInt}

	for i := 1; i < len(arr); i++ {
		m := max(arr[i-1], arr[i])

		aMaxArr = append(aMaxArr, m)
	}

	for i := 1; i < len(arr); i++ {
		m := max(aMinBMaxArr[i-1], aMaxArr[i-1]-arr[i])

		aMinBMaxArr = append(aMinBMaxArr, m)
	}

	for i := 2; i < len(arr); i++ {
		m := max(aMinBPlusCMaxArr[i-1], aMinBMaxArr[i-1]+arr[i])

		aMinBPlusCMaxArr = append(aMinBPlusCMaxArr, m)
	}

	for i := 3; i < len(arr); i++ {
		m := max(aMinBPlusCMinDMaxArr[i-1], aMinBPlusCMaxArr[i-1]-arr[i])

		aMinBPlusCMinDMaxArr = append(aMinBPlusCMinDMaxArr, m)
	}

	return aMinBPlusCMinDMaxArr[len(aMinBPlusCMinDMaxArr)-1]
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/18_maximize_expression
func maximizeExpressionMain() {
	arr := []int{3, 6, 1, -3, 2, 7}
	fmt.Println(maximizeExpression(arr))
}
