package main

import (
	"fmt"
)

func maxPathSum(board [][]int32, p, q int32) int32 {
	rows := len(board)
	cols := len(board[0])

	// Top-down DP using O(m) space
	prev := make([]int32, cols)
	curr := make([]int32, cols)
	prev[p] = board[0][p]

	for i := 1; i < rows; i++ {
		for j := 0; j < cols; j++ {
			best := int32(0)
			for _, dj := range []int{-1, 0, 1} {
				nj := j + dj
				if nj >= 0 && nj < cols {
					best = max(best, prev[nj])
				}
			}
			if best > 0 {
				curr[j] = board[i][j] + best
			} else {
				curr[j] = 0
			}
		}
		prev, curr = curr, prev // swap for next row
	}

	maxTop := int32(0)
	for _, val := range prev {
		maxTop = max(maxTop, val)
	}

	// Bottom-up DP using O(m) space
	prev = make([]int32, cols)
	curr = make([]int32, cols)
	prev[q] = board[rows-1][q]

	for i := rows - 2; i >= 0; i-- {
		for j := 0; j < cols; j++ {
			best := int32(0)
			for _, dj := range []int{-1, 0, 1} {
				nj := j + dj
				if nj >= 0 && nj < cols {
					best = max(best, prev[nj])
				}
			}
			if best > 0 {
				curr[j] = board[i][j] + best
			} else {
				curr[j] = 0
			}
		}
		prev, curr = curr, prev // swap for next row
	}

	maxBottom := int32(0)
	for _, val := range prev {
		maxBottom = max(maxBottom, val)
	}

	return max(maxTop, maxBottom)
}

// You are given a rectangular grid board where each cell contains a positive integer. The top-left corner of the grid is
// at position (0, 0). You can start from either:

// The top row at column p, i.e., position (0, p), or
// The bottom row at column q, i.e., position (rows - 1, q)
// At each step, you can move only one cell per row, within the bounds of the grid. The goal is to determine the
// maximum score that can be obtained by traversing the grid from either starting point.
// The score is the sum of all integers in the cells visited during the traversal.

// 🔀 Movement Rules
// From a cell (i, j) in the top row path, you can move to:
// (i + 1, j - 1)
// (i + 1, j)
// (i + 1, j + 1)

// From a cell (i, j) in the bottom row path, you can move to:
// (i - 1, j - 1)
// (i - 1, j)
// (i - 1, j + 1)

// 🧪 Example
// Input:
// board = [
//   [1, 2, 3],
//   [4, 5, 6],
//   [7, 8, 9],
// ]
// p = 1
// q = 0
// Starting at (0, 1): 2 → 6 → 9 → Score = 17
// Starting at (2, 0): 7 → 5 → 3 → Score = 15

// Output: 17
// 📌 Constraints
// 2 < n, m < 501 — Grid size

// 0 < board[i][j] < 501 — Valid positive cell values

// 0 <= p, q <= m - 1 — Valid starting column indices
func maxPathSumMain() {
	board := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	p, q := int32(1), int32(0)
	fmt.Println(maxPathSum(board, p, q)) // Output: 17

	board = [][]int32{
		{9, 4, 7},
		{2, 1, 3},
		{1, 4, 2},
	}
	p, q = int32(2), int32(1)
	fmt.Println(maxPathSum(board, p, q)) // Output: 15
}
