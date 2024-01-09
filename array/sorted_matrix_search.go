package main

import "fmt"

func SearchInSortedMatrix(matrix [][]int, target int) []int {
	result := []int{-1, -1}
	rowLen := len(matrix)
	if rowLen == 0 {
		return result
	}

	colLen := len(matrix[0])
	row, col := 0, colLen-1

	for row < rowLen && col >= 0 {
		if matrix[row][col] == target {
			result = []int{row, col}

			return result
		}

		if target < matrix[row][col] {
			col--
		} else {
			row++
		}
	}

	return result
}

func sortedMatrixSearch() {
	matrix := [][]int{
		{1, 4, 7, 12, 15, 1000},
		{2, 5, 19, 31, 32, 1001},
		{3, 8, 24, 33, 35, 1002},
		{40, 41, 42, 44, 45, 1003},
		{99, 100, 103, 106, 128, 1004},
	}
	fmt.Println(SearchInSortedMatrix(matrix, 4))
}
