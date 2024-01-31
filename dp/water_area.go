package main

import "fmt"

func findWaterArea(heights []int) int {
	surfaceArea := 0

	leftMaxArr := make([]int, len(heights))

	leftMax := 0
	for i := range heights {
		leftMaxArr[i] = leftMax

		if heights[i] > leftMax {
			leftMax = heights[i]
		}
	}

	rightMaxArr := make([]int, len(heights))

	rightMax := 0
	for i := len(heights) - 1; i >= 0; i-- {
		rightMaxArr[i] = rightMax

		if heights[i] > rightMax {
			rightMax = heights[i]
		}
	}

	for i := range heights {
		height := heights[i]

		minHeight := min(leftMaxArr[i], rightMaxArr[i])
		if minHeight > height {
			surfaceArea += minHeight - height
		}
	}

	return surfaceArea
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/13_water_area
func waterArea() {
	heights := []int{0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3}
	fmt.Println(findWaterArea(heights))
}
