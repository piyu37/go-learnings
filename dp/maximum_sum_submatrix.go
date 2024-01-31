package main

import "fmt"

func findMaxSumSubmatrix(matrix [][]int, size int) int {
	sums := make([][]int, len(matrix))

	maxSum := 0
	for i := 0; i < len(matrix); i++ {
		sums[i] = make([]int, len(matrix[0]))

		if i == 0 {
			sums[i][0] = matrix[i][0]
			for j := 1; j < len(matrix[0]); j++ {
				sums[i][j] = sums[i][j-1] + matrix[i][j]
			}

			continue
		}

		sums[i][0] = sums[i-1][0] + matrix[i][0]

		for j := 1; j < len(matrix[0]); j++ {
			sums[i][j] = sums[i][j-1] + sums[i-1][j] + matrix[i][j] - sums[i-1][j-1]
		}
	}

	for i := size - 1; i < len(sums); i++ {
		for j := size - 1; j < len(sums[0]); j++ {
			subMatrixSum := sums[i][j]
			if i >= size && j >= size {
				subMatrixSum = subMatrixSum - sums[i-size][j] - sums[i][j-size] + sums[i-size][j-size]
			} else if i >= size {
				subMatrixSum = subMatrixSum - sums[i-size][j]
			} else if j >= size {
				subMatrixSum = subMatrixSum - sums[i][j-size]
			}

			if subMatrixSum > maxSum {
				maxSum = subMatrixSum
			}
		}
	}

	return maxSum
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/19_maximum_sum_submatrix
func maximumSumSubmatrix() {
	matrix := [][]int{
		{5, 3, -1, 5},
		{-7, 3, 7, 4},
		{12, 8, 0, 0},
		{1, -8, -8, 2},
	}

	size := 2

	fmt.Println(findMaxSumSubmatrix(matrix, size))
}
