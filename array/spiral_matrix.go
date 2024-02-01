package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	result := []int{}

	rowLenOrg := len(matrix)    // bottom
	colLenOrg := len(matrix[0]) // right
	rowOrg := 0                 // top
	colOrg := 0                 // left
	for rowOrg < rowLenOrg && colOrg < colLenOrg {
		i := colOrg
		for i < colLenOrg {
			result = append(result, matrix[rowOrg][i])
			i++
		}

		rowOrg++
		i = rowOrg
		for i < rowLenOrg {
			result = append(result, matrix[i][colLenOrg-1])
			i++
		}

		colLenOrg--
		if rowOrg < rowLenOrg {
			i = colLenOrg - 1
			for i >= colOrg {
				result = append(result, matrix[rowLenOrg-1][i])
				i--
			}

			rowLenOrg--
		}

		if colOrg < colLenOrg {
			i = rowLenOrg - 1
			for i >= rowOrg {
				result = append(result, matrix[i][colOrg])
				i--
			}
		}

		colOrg++
	}

	return result
}

// https://leetcode.com/problems/spiral-matrix/description/?envType=study-plan-v2&envId=top-interview-150
func spiralMatrix() {
	arr := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(arr))
}
