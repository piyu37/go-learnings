package prefixSuffix

import (
	"fmt"
	"math"
)

func findLeftSmallRightLargeInArray(arr []int) int {
	maxLeftSmallRightLargeNum := math.MinInt

	rightMin := make([]int, len(arr))
	rightMin[len(arr)-1] = math.MaxInt

	for i := len(arr) - 2; i >= 0; i-- {
		rightMin[i] = min(arr[i+1], rightMin[i+1])
	}

	leftMax := math.MinInt

	for i := range arr {
		if arr[i] > leftMax && arr[i] < rightMin[i] {
			maxLeftSmallRightLargeNum = max(maxLeftSmallRightLargeNum, arr[i])
		}

		leftMax = max(leftMax, arr[i])
	}

	return maxLeftSmallRightLargeNum
}

// https://www.geeksforgeeks.org/dsa/find-the-element-before-which-all-the-elements-are-smaller-than-it-and-after-which-all-are-greater-than-it/
func FindLeftSmallRightLargeInArrayMain() {
	arr := []int{9, 1, 12, 35, 14, 25}
	fmt.Println(findLeftSmallRightLargeInArray(arr))
}
