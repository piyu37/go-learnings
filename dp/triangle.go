package main

import (
	"fmt"
	"math"
)

// Not working properly for few edge cases; not able to understand why
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	if len(triangle) == 1 {
		return triangle[0][0]
	}

	minPath := findMinPath(1, 0, 1, 1, triangle, map[string]int{})
	result := triangle[0][0] + minPath

	return result
}

func findMinPath(i1, j1, i2, j2 int, triangle [][]int, memztn map[string]int) int {
	if i1 == len(triangle) {
		return 0
	}

	key1 := fmt.Sprintf("%d%d", i1, j1)
	key2 := fmt.Sprintf("%d%d", i2, j2)
	var min1, min2 int
	if v, ok := memztn[key1]; ok {
		min1 = v
	} else {
		min1 = triangle[i1][j1] + findMinPath(i1+1, j1, i1+1, j1+1, triangle, memztn)
		memztn[key1] = min1
	}

	if v, ok := memztn[key2]; ok {
		min2 = v
	} else {
		min2 = triangle[i2][j2] + findMinPath(i2+1, j2, i2+1, j2+1, triangle, memztn)
		memztn[key2] = min2
	}

	return min(min1, min2)
}

// best solution
func minimumTotal2(triangle [][]int) int {
	for row := 1; row < len(triangle); row++ {
		for col := 0; col <= row; col++ {
			smallestFromAbove := math.MaxInt32

			if col > 0 {
				smallestFromAbove = triangle[row-1][col-1]
			}

			if col < row {
				smallestFromAbove = min(smallestFromAbove, triangle[row-1][col])
			}

			path := smallestFromAbove + triangle[row][col]

			triangle[row][col] = path
		}
	}

	minVal := triangle[len(triangle)-1][0]

	arr := triangle[len(triangle)-1]

	for i := 1; i < len(arr); i++ {
		if arr[i] < minVal {
			minVal = arr[i]
		}
	}

	return minVal
}

// https://leetcode.com/problems/triangle/description/?envType=study-plan-v2&envId=top-interview-150
func triangle() {
	// Example usage
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle)) // Output: 11
	fmt.Println(minimumTotal2(triangle))
}
