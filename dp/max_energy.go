package main

import (
	"fmt"
	"math"
)

func maxEnergy(mat [][]int) int {
	min := math.MinInt16
	maxEnergySaving := math.MinInt16

	storedValues := []int{min}

	for i := 0; i < len(mat); i++ {
		storedValues = append(storedValues, 100)
	}

	storedValues = append(storedValues, min)
	upArr := make([]int, len(storedValues))

	for i := 1; i <= len(mat); i++ {
		copy(upArr, storedValues)
		for j := 1; j <= len(mat[0]); j++ {
			max := max(max(upArr[j-1], upArr[j]), upArr[j+1])
			storedValues[j] = max - mat[i-1][j-1]
			if i == len(mat) && storedValues[j] > maxEnergySaving {
				maxEnergySaving = storedValues[j]
			}
		}
	}

	return maxEnergySaving
}

// variant of below questions:
// Find the maximum energy savings while traversing from the top row to the bottom row of the matrix
// possible moves:
// Down: (row+1,col)
// Down left diagonal: (row+1,col-1)
// Down right diagonal: (row+1, col+1)
// output for below input: 100 - (22+55+28+11) = -16
// https://takeuforward.org/data-structure/minimum-maximum-falling-path-sum-dp-12/
// https://www.codingninjas.com/studio/problems/maximum-path-sum-in-the-matrix_797998
func maxEnergyMain() {
	mat := [][]int{
		{51, 22, 69, 45},
		{55, 87, 81, 15},
		{63, 28, 66, 63},
		{49, 19, 11, 6},
	}

	fmt.Println(maxEnergy(mat))
}
