package main

import (
	"fmt"
	"practice/common"
)

func NonAttackingQueens(n int) int {
	blockedCols := make(map[int]bool)
	blockedLeftDiagonals := make(map[int]bool)  // anti diagonal or up diagonal
	blockedRightDiagonals := make(map[int]bool) // main diagonal or down diagonal

	return calculateNonAttackingPositions(0, blockedCols, blockedLeftDiagonals, blockedRightDiagonals, n)
}

func calculateNonAttackingPositions(row int, blockedCols, blockedLeftDiagonals, blockedRightDiagonals map[int]bool, n int) int {
	result := 0

	if row == n {
		return 1
	}

	for j := 0; j < n; j++ {
		if isValidPosition(row, j, blockedCols, blockedLeftDiagonals, blockedRightDiagonals) {
			placeQueen(row, j, blockedCols, blockedLeftDiagonals, blockedRightDiagonals)

			result += calculateNonAttackingPositions(row+1, blockedCols, blockedLeftDiagonals, blockedRightDiagonals, n)

			removeQueen(row, j, blockedCols, blockedLeftDiagonals, blockedRightDiagonals)
		}
	}

	return result
}

func isValidPosition(row, col int, blockedCols, blockedLeftDiagonals, blockedRightDiagonals map[int]bool) bool {
	if blockedCols[col] {
		return false
	}

	if blockedLeftDiagonals[row+col] {
		return false
	}

	if blockedRightDiagonals[row-col] {
		return false
	}

	return true
}

func placeQueen(row, col int, blockedCols, blockedLeftDiagonals, blockedRightDiagonals map[int]bool) {
	blockedCols[col] = true
	blockedLeftDiagonals[row+col] = true
	blockedRightDiagonals[row-col] = true
}

func removeQueen(row, col int, blockedCols, blockedLeftDiagonals, blockedRightDiagonals map[int]bool) {
	delete(blockedCols, col)
	delete(blockedLeftDiagonals, row+col)
	delete(blockedRightDiagonals, row-col)
}

func solveNQueens(n int) [][]string {
	result := [][]int{}
	addQueenPlacements(n, 0, []int{}, &result)

	restr := make([][]string, 0)

	for _, combinations := range result {
		arr := make([]string, n, n)
		for i := range n {
			comb := combinations[i]
			for j := range n {
				if j == comb {
					arr[i] += "Q"
				} else {
					arr[i] += "."
				}
			}
		}

		restr = append(restr, arr)
	}

	return restr
}

func addQueenPlacements(n, i int, placement []int, result *[][]int) {
	if i == n {
		temp := make([]int, len(placement))
		copy(temp, placement)

		*result = append(*result, temp)

		return
	}

	for j := range n {
		if canBePlaced(i, j, placement) {
			addQueenPlacements(n, i+1, append(placement, j), result)
		}
	}
}

func canBePlaced(i, j int, placement []int) bool {
	for row, col := range placement {
		if j == col || (common.Abs(i-row) == common.Abs(j-col)) {
			return false
		}
	}

	return true
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/30_non_attacking_queens
func attackingQueen() {
	fmt.Println(NonAttackingQueens(4))
	fmt.Println(NonAttackingQueens(8))

	fmt.Println(solveNQueens(4))
}
