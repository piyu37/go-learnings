package main

import (
	"fmt"
	"practice/common"
)

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N!)
 * ----------------------------
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(N!)
 *    - Row 0: N possible column choices
 *    - Row 1: N-1 possible choices (1 column blocked)
 *    - Row 2: N-2 possible choices
 *    - ...
 *    - Row N: 1 choice
 *    - Max states = N * (N-1) * (N-2) * ... * 1 = N!
 *
 * 2. Work Per State -> O(1)
 *    - At every recursive step, `isValidPosition()`, `placeQueen()`, and `removeQueen()`
 *      perform map lookups, insertions, and deletions to track blocked paths.
 *    - In Go, these map operations execute in O(1) average time.
 *
 * Total Time = O(N!) * O(1) = O(N!)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes exactly N levels deep (one for each row).
 * 2. Auxiliary Space: The three maps (`blockedCols`, `blockedLeftDiagonals`,
 *    `blockedRightDiagonals`) each store a maximum of N boolean values.
 *
 * Total Space = O(N) + O(N) = O(N)
 */
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

	for j := range n {
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

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N * N!)
 * ----------------------------
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(N!)
 *    - Row 0: N possible column choices
 *    - Row 1: N-1 possible choices (1 column blocked)
 *    - Row 2: N-2 possible choices
 *    - ...
 *    - Row N: 1 choice
 *    - Max states = N * (N-1) * (N-2) * ... * 1 = N!
 *
 * 2. Work Per State -> O(N)
 *    - At every recursive step, `canBePlaced()` loops through all previously
 *      placed queens to check for column/diagonal conflicts. This takes up to O(N) time.
 *    - Additionally, when a valid board is found, copying the `placement` slice
 *      into the `result` array takes O(N) time.
 *
 * Total Time = O(N!) * O(N) = O(N * N!)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes exactly N levels deep (one for each row).
 * 2. Auxiliary Space: The `placement` slice holds a maximum of N integers.
 *
 * Total Space = O(N) + O(N) = O(N)
 */
func solveNQueens(n int) [][]string {
	result := [][]int{}
	addQueenPlacements(n, 0, []int{}, &result)

	restr := make([][]string, 0)

	for _, combinations := range result {
		arr := make([]string, n)
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
