package main

import (
	"fmt"
	"math"
)

func maxMatrixSum(matrix [][]int) int64 {
	var min int64 = math.MaxInt32
	var sum int64
	countNegative := 0
	for i := range matrix {
		for j := range matrix[0] {
			num := int64(abs(matrix[i][j]))
			sum += num

			if num < min {
				min = num
			}

			if matrix[i][j] < 0 {
				countNegative++
			}
		}
	}

	if countNegative%2 == 0 {
		return sum
	}

	return sum - (2 * min)
}

// https://leetcode.com/problems/maximum-matrix-sum/
func maxMatrixSumMain() {
	m := [][]int{
		{1, 2, 3},
		{-1, -2, -3},
		{1, 2, 3},
	}

	fmt.Println(maxMatrixSum(m))
}
