package main

import "fmt"

func prefixSumMatrix(arr [][]int) [][]int {
	prefixMatrix := make([][]int, len(arr))

	for r := range arr {
		prefixMatrix[r] = make([]int, len(arr[0]))

		for c := 0; c < len(arr[0]); c++ {
			prefixMatrix[r][c] = arr[r][c]

			if r > 0 {
				prefixMatrix[r][c] += prefixMatrix[r-1][c]
			}

			if c > 0 {
				prefixMatrix[r][c] += prefixMatrix[r][c-1]
			}

			if r > 0 && c > 0 {
				prefixMatrix[r][c] -= prefixMatrix[r-1][c-1]
			}
		}
	}

	return prefixMatrix
}

// https://www.geeksforgeeks.org/dsa/prefix-sum-2d-array/
func prefixSumMatrixMain() {
	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	fmt.Println(prefixSumMatrix(arr))
}
