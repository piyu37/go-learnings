package main

import (
	"fmt"
	"slices"
)

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N * 9^M)
 * ----------------------------
 * Where:
 * - M = The number of empty cells on the board (at most 81).
 * - N = The dimension of the board (N=9 for a standard Sudoku).
 *
 * (Note: Because a standard Sudoku board is strictly fixed at 9x9, theoretically
 * the time complexity is technically O(1). However, in algorithmic interviews,
 * it is standard to express it in terms of M empty cells).
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(9^M)
 *    - The algorithm uses classic Backtracking. For every empty cell, it attempts
 *      to place up to 9 different digits.
 *    - This creates a decision tree with a maximum branching factor of 9 and a
 *      maximum depth of M.
 *    - In the absolute mathematical worst case, this generates up to 9^M states.
 *      (In reality, the tight constraints of Sudoku rules prune the vast majority
 *      of these branches very early, making the actual execution much faster).
 *
 * 2. Work Per State -> O(N)
 *    - At each recursive step, `isValidForPosition()` checks the row, the column,
 *      and the 3x3 subgrid.
 *    - `slices.Contains(board[row], value)` takes O(N) time.
 *    - Iterating through the column takes O(N) time.
 *    - Iterating through the 3x3 subgrid takes O(N) time.
 *    - Therefore, the validation checks take O(N) time per state.
 *
 * Total Time = O(9^M) * O(N) = O(N * 9^M)
 *
 * SPACE COMPLEXITY: O(M)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes exactly 1 level deep for every empty
 *    cell on the board. In the worst case (an entirely empty board), the maximum
 *    depth is M (or 81).
 * 2. Auxiliary Space: The board is modified perfectly in-place. You do not create
 *    any new arrays or slices during the recursion.
 *
 * Total Space = Stack O(M) + Extra Space O(1) = O(M)
 */
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
	if slices.Contains(board[row], value) {
		return false
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

// Another approach to check validity for position better than isValidForPosition
func isValidForPosition2(board [][]byte, row, col int, val int) bool {
	for i := range board {
		if board[i][col] == byte(val+'0') {
			return false
		}

		if board[row][i] == byte(val+'0') {
			return false
		}

		if board[3*(row/3)+i/3][3*(col/3)+i%3] == byte(val+'0') {
			return false
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
