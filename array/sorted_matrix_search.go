package main

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
