package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	maxVal := math.MinInt16

	subArr := math.MinInt16
	for i := 0; i < len(nums); i++ {
		subArr = max(nums[i], subArr+nums[i])
		if subArr > maxVal {
			maxVal = subArr
		}
	}

	return maxVal
}

func maxSubarraySumCircular(nums []int) int {
	maxVal, minVal := nums[0], nums[0]

	subArrMax, subArrMin := nums[0], nums[0]
	total := nums[0]
	for i := 1; i < len(nums); i++ {
		subArrMax = max(nums[i], subArrMax+nums[i])
		if subArrMax > maxVal {
			maxVal = subArrMax
		}

		subArrMin = min(nums[i], subArrMin+nums[i])
		if subArrMin < minVal {
			minVal = subArrMin
		}

		total += nums[i]
	}

	if maxVal < 0 {
		return maxVal
	}

	return max(maxVal, total-minVal)
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

// https://leetcode.com/problems/maximum-subarray/description/?envType=study-plan-v2&envId=top-interview-150
// https://leetcode.com/problems/maximum-sum-circular-subarray/description/?envType=study-plan-v2&envId=top-interview-150
func maxSubarrayKadane() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr))
	fmt.Println(maxSubarraySumCircular(arr))
	arr = []int{9, -4, -7, 9}
	fmt.Println(maxSubArray(arr))
	fmt.Println(maxSubarraySumCircular(arr))
	arr = []int{-3, -2, -3}
	fmt.Println(maxSubArray(arr))
	fmt.Println(maxSubarraySumCircular(arr))
}
