package main

import "fmt"

func SolveSudoku(board [][]int) [][]int {
	solvePartialSudoku(0, 0, board)

	return board
}

func solvePartialSudoku(row, col int, board [][]int) bool {
	currentRow := row
	currentCol := col

	if currentCol == len(board[row]) {
		currentRow++
		currentCol = 0

		if currentRow == len(board) {
			return true
		}
	}

	if board[currentRow][currentCol] == 0 {
		return tryDigitsAtPosition(currentRow, currentCol, board)
	}

	return solvePartialSudoku(currentRow, currentCol+1, board)
}

func tryDigitsAtPosition(row, col int, board [][]int) bool {
	for i := 1; i <= 9; i++ {
		if isValidForPosition(i, row, col, board) {
			board[row][col] = i
			if solvePartialSudoku(row, col+1, board) {
				return true
			}
		}
	}

	board[row][col] = 0

	return false
}

func isValidForPosition(value, row, col int, board [][]int) bool {
	for _, v := range board[row] {
		if v == value {
			return false
		}
	}

	for _, r := range board {
		v := r[col]
		if v == value {
			return false
		}
	}

	subGridForRow := row / 3
	subGridForCol := col / 3

	for i := 3 * subGridForRow; i < 3*subGridForRow+3; i++ {
		for j := 3 * subGridForCol; j < 3*subGridForCol+3; j++ {
			if board[i][j] == value {
				return false
			}
		}
	}

	return true
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/31_solve_sudoku
func sudoku() {
	board := [][]int{
		{7, 8, 0, 4, 0, 0, 1, 2, 0},
		{6, 0, 0, 0, 7, 5, 0, 0, 9},
		{0, 0, 0, 6, 0, 1, 0, 7, 8},
		{0, 0, 7, 0, 4, 0, 2, 6, 0},
		{0, 0, 1, 0, 5, 0, 9, 3, 0},
		{9, 0, 4, 0, 6, 0, 0, 0, 5},
		{0, 7, 0, 3, 0, 0, 0, 1, 2},
		{1, 2, 0, 0, 0, 7, 4, 0, 0},
		{0, 4, 9, 2, 0, 6, 0, 0, 7},
	}

	fmt.Println(SolveSudoku(board))
}
