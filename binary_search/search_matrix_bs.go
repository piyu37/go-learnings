package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1
	row := -1

	for left <= right {
		mid := left + (right-left)/2

		if target >= matrix[mid+1][0] {
			if matrix[mid+1][0] == target {
				return true
			}

			left = mid + 1
		} else if target < matrix[mid][0] {
			right = mid - 1
		} else {
			if matrix[mid][0] == target {
				return true
			}

			row = mid
			left++
			right--
		}
	}

	left = 1
	right = len(matrix[0]) - 1

	for left <= right {
		mid := left + (right-left)/2

		if target > matrix[row][mid] {
			left = mid + 1
		} else if target < matrix[row][mid] {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}

// https://leetcode.com/problems/search-a-2d-matrix/description/?envType=study-plan-v2&envId=top-interview-150
func searchMatrixMain() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix(matrix, 3))
}
