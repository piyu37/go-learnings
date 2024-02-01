package main

import "fmt"

func rotate(matrix [][]int) {
	row := 0
	col := 0
	rowLen := len(matrix) - 1
	colLen := len(matrix[0]) - 1

	for row <= rowLen && col <= colLen {
		prevArr := []int{}

		for i := col; i <= colLen-1; i++ {
			prevArr = append(prevArr, matrix[row][i])
		}

		for i := row; i < rowLen; i++ {
			prev := matrix[i][colLen]
			matrix[i][colLen] = prevArr[i-row]
			prevArr[i-row] = prev
		}

		for i := colLen; i > col; i-- {
			prev := matrix[rowLen][i]
			matrix[rowLen][i] = prevArr[colLen-i]
			prevArr[colLen-i] = prev
		}

		for i := rowLen; i > row; i-- {
			prev := matrix[i][col]
			matrix[i][col] = prevArr[rowLen-i]
			prevArr[rowLen-i] = prev
		}

		for i := col; i < colLen; i++ {
			matrix[row][i] = prevArr[i-col]
		}

		row++
		col++
		rowLen--
		colLen--
	}
}

// https://leetcode.com/problems/rotate-image/description/?envType=study-plan-v2&envId=top-interview-150
func rotateImage() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(arr)
	fmt.Println(arr)
}
