package main

import "fmt"

func NonAttackingQueens(n int) int {
	blockedCols := make(map[int]bool)
	blockedUpDiagonals := make(map[int]bool)
	blockedDownDiagonals := make(map[int]bool)

	return calculateNonAttackingPositions(0, blockedCols, blockedUpDiagonals, blockedDownDiagonals, n)
}

func calculateNonAttackingPositions(row int, blockedCols, blockedUpDiagonals, blockedDownDiagonals map[int]bool, n int) int {
	result := 0

	if row == n {
		return 1
	}

	for j := 0; j < n; j++ {
		if isValidPosition(row, j, blockedCols, blockedUpDiagonals, blockedDownDiagonals) {
			placeQueen(row, j, blockedCols, blockedUpDiagonals, blockedDownDiagonals)

			result += calculateNonAttackingPositions(row+1, blockedCols, blockedUpDiagonals, blockedDownDiagonals, n)

			removeQueen(row, j, blockedCols, blockedUpDiagonals, blockedDownDiagonals)
		}
	}

	return result
}

func isValidPosition(row, col int, blockedCols, blockedUpDiagonals, blockedDownDiagonals map[int]bool) bool {
	if blockedCols[col] {
		return false
	}

	if blockedUpDiagonals[row+col] {
		return false
	}

	if blockedDownDiagonals[row-col] {
		return false
	}

	return true
}

func placeQueen(row, col int, blockedCols, blockedUpDiagonals, blockedDownDiagonals map[int]bool) {
	blockedCols[col] = true
	blockedUpDiagonals[row+col] = true
	blockedDownDiagonals[row-col] = true
}

func removeQueen(row, col int, blockedCols, blockedUpDiagonals, blockedDownDiagonals map[int]bool) {
	delete(blockedCols, col)
	delete(blockedUpDiagonals, row+col)
	delete(blockedDownDiagonals, row-col)
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/30_non_attacking_queens
func attackingQueen() {
	fmt.Println(NonAttackingQueens(8))
}
