package main

import "fmt"

func trap(height []int) int {
	low := 0
	high := len(height) - 1

	for i := range height {
		if height[i] > 0 {
			low = i
			break
		}
	}

	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > 0 {
			high = i
			break
		}
	}

	if low == high {
		return 0
	}

	result := 0
	trapMap := make(map[string]bool)

	left := low + 1
	right := high

	min := low
	sum := 0
	for left <= right {
		if height[left] < height[min] {
			sum += height[min] - height[left]
		} else {
			result += sum
			key := fmt.Sprintf("%d%d", min, left)
			trapMap[key] = true
			sum = 0
			min = left
		}

		left++
	}

	min = high
	sum = 0
	for high-1 >= low {
		if height[high-1] < height[min] {
			sum += height[min] - height[high-1]
		} else {
			key := fmt.Sprintf("%d%d", high-1, min)
			if !trapMap[key] {
				result += sum
			}
			sum = 0
			min = high - 1
		}

		high--
	}

	return result
}

func trap2(height []int) int {
	left := 0
	right := len(height) - 1

	leftMax := 0
	rightMax := 0
	sum := 0
	for left < right {
		if height[left] <= height[right] {
			if height[left] < leftMax {
				sum += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
			left++
		} else {
			if height[right] < rightMax {
				sum += rightMax - height[right]
			} else {
				rightMax = height[right]
			}

			right--
		}
	}

	return sum
}

// https://leetcode.com/problems/trapping-rain-water/description/?envType=study-plan-v2&envId=top-interview-150
func trappingRainWater() {
	height := []int{0, 3, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
	fmt.Println(trap2(height))
}
