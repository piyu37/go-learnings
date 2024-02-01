package main

import "fmt"

func largestRectangleArea(heights []int) int {
	left := make([]int, 0)
	right := make([]int, 0)

	for range heights {
		right = append(right, len(heights))
		left = append(left, -1)
	}

	stack := []int{}

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	stack = []int{}

	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			left[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	maxArea := 0

	for i := range heights {
		area := heights[i] * (right[i] - left[i] - 1)

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func largestRectangleAreaOptimal(heights []int) int {
	stack := []int{}
	maxArea := 0
	for i := 0; i <= len(heights); i++ {
		for len(stack) > 0 && (i == len(heights) || heights[i] < heights[stack[len(stack)-1]]) {
			right := i
			curr := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			left := -1
			area := 0
			if len(stack) > 0 {
				left = stack[len(stack)-1]
				area = curr * (right - left - 1)
			} else {
				area = curr * right
			}

			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}

// stack related
// https://takeuforward.org/data-structure/area-of-largest-rectangle-in-histogram/
// https://www.codingninjas.com/studio/problems/largest-rectangle-in-a-histogram_1058184
func largestRectangleAreaMain() {
	heights := []int{2, 4}

	fmt.Println(largestRectangleArea(heights))
	fmt.Println(largestRectangleAreaOptimal(heights))
}
