package main

import (
	"fmt"
	"math"
)

// getSum returns the sum of submatrix [(r1, c1) ... (r2, c2)]
// using the 0-indexed 2D prefix sum array, where pre[r][c] is the
// sum of mat[0..r][0..c].
func getSum(pre [][]int, r1, c1, r2, c2 int) int {
	sum := pre[r2][c2]

	if r1 > 0 {
		sum -= pre[r1-1][c2]
	}

	if c1 > 0 {
		sum -= pre[r2][c1-1]
	}

	if r1 > 0 && c1 > 0 {
		sum += pre[r1-1][c1-1]
	}

	return sum
}

func maxSubMatSum(mat [][]int, queries [][2]int) []int {
	n := len(mat)
	m := len(mat[0])

	// Build 2D prefix sum array.
	pre := make([][]int, n)

	for r := range mat {
		pre[r] = make([]int, len(mat[0]))

		for c := 0; c < len(mat[0]); c++ {
			pre[r][c] = mat[r][c]

			if r > 0 {
				pre[r][c] += pre[r-1][c]
			}

			if c > 0 {
				pre[r][c] += pre[r][c-1]
			}

			if r > 0 && c > 0 {
				pre[r][c] -= pre[r-1][c-1]
			}
		}
	}

	res := make([]int, 0, len(queries))

	// Process each query independently.
	for _, qr := range queries {
		a, b := qr[0], qr[1]

		mx := math.MinInt

		// Try every possible a x b submatrix.
		for i := 0; i <= n-a; i++ {
			for j := 0; j <= m-b; j++ {
				sumVal := getSum(pre, i, j, i+a-1, j+b-1)
				if sumVal > mx {
					mx = sumVal
				}
			}
		}

		res = append(res, mx)
	}

	return res
}

// getSum2 returns the sum of submatrix [(r1, c1) ... (r2, c2)]
// using the 2D prefix sum array.
func getSum2(pre [][]int, r1, c1, r2, c2 int) int {
	return pre[r2+1][c2+1] -
		pre[r1][c2+1] -
		pre[r2+1][c1] +
		pre[r1][c1]
}

func maxSubMatSum2(mat [][]int, queries [][2]int) []int {
	n := len(mat)
	m := len(mat[0])

	// Build 2D prefix sum array.
	pre := make([][]int, n+1)
	for i := range pre {
		pre[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			pre[i][j] = mat[i-1][j-1] +
				pre[i-1][j] +
				pre[i][j-1] -
				pre[i-1][j-1]
		}
	}

	res := make([]int, 0, len(queries))

	// Process each query independently.
	for _, qr := range queries {
		a, b := qr[0], qr[1]

		mx := math.MinInt

		// Try every possible a x b submatrix.
		for i := 0; i <= n-a; i++ {
			for j := 0; j <= m-b; j++ {
				sumVal := getSum2(pre, i, j, i+a-1, j+b-1)
				if sumVal > mx {
					mx = sumVal
				}
			}
		}

		res = append(res, mx)
	}

	return res
}

// https://www.geeksforgeeks.org/dsa/maximum-sum-submatrix/
func maxSumSubmatrixWithQuery() {
	mat := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	queries := [][2]int{
		{2, 2},
		{1, 1},
		{3, 3},
	}

	// maxSubMatSum: prefix array is n x m (0-indexed), pre[r][c] = sum of mat[0..r][0..c].
	// getSum needs `if r1 > 0` / `if c1 > 0` checks, since pre[r1-1] / pre[..][c1-1]
	// would go out of bounds for the first row/col.
	ans := maxSubMatSum(mat, queries)

	fmt.Println(ans)

	// maxSubMatSum2: prefix array is (n+1) x (m+1) (1-indexed, row 0 & col 0 are all 0),
	// pre[i][j] = sum of mat[0..i-1][0..j-1]. The zero padding removes the boundary
	// checks, so getSum2 is a single formula: pre[r2+1][c2+1] - pre[r1][c2+1] - pre[r2+1][c1] + pre[r1][c1].
	// Same time complexity for both; version 2 just has cleaner code.
	fmt.Println(maxSubMatSum2(mat, queries))
}
