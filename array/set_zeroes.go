package main

import (
	"fmt"
	"math"
)

func setZeroes(matrix [][]int) {
	colLen := len(matrix[0])
	min := math.MinInt

	for row := range matrix {
		for col := 0; col < colLen; col++ {
			if matrix[row][col] == 0 {
				for i := col - 1; i >= 0; i-- {
					matrix[row][i] = min
				}

				for i := col + 1; i < colLen; i++ {
					if matrix[row][i] != 0 {
						matrix[row][i] = min
					}
				}

				break
			}
		}
	}

	for col := 0; col < colLen; col++ {
		for row := 0; row < len(matrix); row++ {
			if matrix[row][col] == 0 {
				for i := row - 1; i >= 0; i-- {
					matrix[i][col] = min
				}

				for i := row + 1; i < len(matrix); i++ {
					matrix[i][col] = min
				}

				matrix[row][col] = min

				break
			}
		}
	}

	for row := range matrix {
		for col := 0; col < colLen; col++ {
			if matrix[row][col] == min {
				matrix[row][col] = 0
			}
		}
	}
}

// https://leetcode.com/problems/set-matrix-zeroes/description/?envType=study-plan-v2&envId=top-interview-150
func setZeroesMain() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
