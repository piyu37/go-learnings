package main

import (
	"fmt"
	"strconv"
)

func isValidSudoku(board [][]byte) bool {
	rowArr := make([]int, 10)
	colArr := make([]int, 10)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[j][i] == 46 {
				continue
			}

			val, _ := strconv.Atoi(string(board[j][i]))
			if rowArr[val] == 0 {
				rowArr[val] = val
			} else {
				return false
			}
		}

		for j := 0; j < 9; j++ {
			if board[i][j] == 46 {
				continue
			}

			val, _ := strconv.Atoi(string(board[i][j]))
			if colArr[val] == 0 {
				colArr[val] = val
			} else {
				return false
			}
		}

		rowArr = make([]int, 10)
		colArr = make([]int, 10)
	}

	segArr := make([]int, 10)
	row := 0
	col := 0
	for row < 9 {
		for col < 9 {
			for i := row + 0; i < row+3; i++ {
				for j := col + 0; j < col+3; j++ {
					if board[i][j] == 46 {
						continue
					}

					val, _ := strconv.Atoi(string(board[i][j]))
					if segArr[val] == 0 {
						segArr[val] = val
					} else {
						return false
					}
				}
			}

			segArr = make([]int, 10)

			col += 3
		}

		col = 0
		row += 3
	}

	return true
}

// https://leetcode.com/problems/valid-sudoku/description/?envType=study-plan-v2&envId=top-interview-150
func validSudoku() {
	board := [][]byte{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	fmt.Println(isValidSudoku(board))
}
