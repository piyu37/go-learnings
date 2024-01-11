package main

import "fmt"

func LargestRectangleUnderSkyline(buildings []int) int {
	largestRectangle := 0

	stack := make([]int, 0)

	extendedBuildings := append(buildings, 0)

	for i := range extendedBuildings {
		for len(stack) > 0 && buildings[stack[len(stack)-1]] >= extendedBuildings[i] {
			poppedElementHeight := buildings[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = width - stack[len(stack)-1] - 1
			}

			maxArea := poppedElementHeight * width

			if maxArea > largestRectangle {
				largestRectangle = maxArea
			}
		}

		stack = append(stack, i)
	}

	return largestRectangle
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/42_largest_rectangle_under_skyline
func largestBuilding() {
	arr := []int{1, 3, 3, 2, 4, 1, 5, 3, 2}

	fmt.Println(LargestRectangleUnderSkyline(arr))
}
