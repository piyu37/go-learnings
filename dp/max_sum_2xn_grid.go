package main

import "fmt"

// Picking any cell in column y blocks the other cell in column y (vertical)
// and every cell in columns y-1 and y+1 (horizontal + diagonal).
// So from each column we take at most one cell, the bigger one, and no two
// adjacent columns can both be used. That's exactly House Robber (rob2) where
// each column is a house worth max(mat[0][y], mat[1][y]).
// Time: O(n), Space: O(1)
func maxSum2xNGrid(mat [][]int) int {
	// prev2 = best total up to two columns back
	// prev1 = best total up to the previous column
	prev2, prev1 := 0, 0

	for y := 0; y < len(mat[0]); y++ {
		money := max(mat[0][y], mat[1][y])

		// skip this column -> keep prev1
		// use this column  -> prev2 + money
		best := max(prev1, prev2+money)
		prev2, prev1 = prev1, best
	}
	return prev1
}

// https://www.geeksforgeeks.org/dsa/maximum-sum-2-x-n-grid-no-two-elements-adjacent/
func maxSum2xNGridMain() {
	fmt.Println(maxSum2xNGrid([][]int{{1, 4, 5}, {2, 0, 0}})) // 7
	fmt.Println(maxSum2xNGrid([][]int{{1, 2}, {3, 4}}))       // 4
}
