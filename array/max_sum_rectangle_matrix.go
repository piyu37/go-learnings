package main

import (
	"fmt"
	"math"
)

// maxRectSum finds the max sum rectangle by reducing the 2D problem to 1D Kadane.
//
// Idea: every rectangle is defined by a left column (c1), a right column (c2),
// a top row and a bottom row.
//   - Fix the column pair (c1, c2) with two loops.
//   - Squash the columns c1..c2 into one 1D array: rowSubArraySum[r] holds the
//     sum of matrix[r][c1..c2]. When c2 moves right by one, just add
//     matrix[r][c2] to each entry, so there is no need to recompute.
//   - Any contiguous range of rows (top..bottom) in rowSubArraySum is the sum of
//     the rectangle bounded by rows top..bottom and columns c1..c2.
//   - So Kadane on rowSubArraySum gives the best top/bottom pair for this
//     column pair. The best result over all (c1, c2) is the answer.
//
// Example: c1=1, c2=3 -> rowSubArraySum = [2+(-1)+(-4), -3+4+2, 8+10+1, -1+1+7]
// = [-3, 3, 19, 7]. Kadane gives 3+19+7 = 29, i.e. rows 1..3, cols 1..3.
//
// Time: O(cols^2 * rows), Space: O(rows).
func maxRectSum(matrix [][]int) int {
	maxSum := math.MinInt16
	colLen := len(matrix[0])

	for c1 := range colLen {
		rowSubArraySum := make([]int, len(matrix))

		for c2 := c1; c2 < colLen; c2++ {

			for r := range matrix {
				rowSubArraySum[r] += matrix[r][c2]
			}

			subArrSum := math.MinInt16
			for i := range rowSubArraySum {
				subArrSum = max(rowSubArraySum[i], subArrSum+rowSubArraySum[i])
				if subArrSum > maxSum {
					maxSum = subArrSum
				}
			}
		}
	}

	return maxSum
}

// https://www.geeksforgeeks.org/dsa/maximum-sum-rectangle-in-a-2d-matrix/
func maxSumRectangleMatrixKadane() {
	matrix := [][]int{
		{1, 2, -1, -4, -20},
		{-8, -3, 4, 2, 1},
		{3, 8, 10, 1, 3},
		{-4, -1, 1, 7, -6},
	}

	fmt.Println(maxRectSum(matrix))

	fmt.Println(maxRectSumPos(matrix))
}

func maxRectSumPos(matrix [][]int) (int, int, int, int, int) {
	maxSum := math.MinInt16
	colLen := len(matrix[0])

	maxSumStartRowIdx, maxSumEndRowIdx, maxSumStartColIdx, maxSumEndColIdx := 0, 0, 0, 0

	for c1 := range colLen {
		rowSubArraySum := make([]int, len(matrix))

		for c2 := c1; c2 < colLen; c2++ {

			for r := range matrix {
				rowSubArraySum[r] += matrix[r][c2]
			}

			subArrSum := 0
			startRowIdx := 0
			for i := range rowSubArraySum {
				if subArrSum < 0 {
					subArrSum = 0
					startRowIdx = i
				}
				
				subArrSum += rowSubArraySum[i]

				if subArrSum > maxSum {
					maxSum = subArrSum
					maxSumStartRowIdx, maxSumEndRowIdx, maxSumStartColIdx, maxSumEndColIdx = startRowIdx, i, c1, c2
				}
			}
		}
	}

	return maxSum, maxSumStartRowIdx, maxSumStartColIdx, maxSumEndRowIdx, maxSumEndColIdx
}
