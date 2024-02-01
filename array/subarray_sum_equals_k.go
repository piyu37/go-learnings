package main

import "fmt"

// For only positive values in array
func subarraySum(nums []int, k int) int {
	maxCount := 0
	sum := 0
	left := 0
	right := 0

	for left < len(nums) {
		for right < len(nums) && sum < k {
			sum += nums[right]
			right++
		}

		if sum == k && right > 0 {
			maxCount++
		}

		for left <= right && left < len(nums) && sum >= k {
			sum -= nums[left]
			left++
		}

		if right == left {
			sum += nums[right]
			right++
		}
	}

	return maxCount
}

// For both positive & negative values in array
func subarraySum2(arr []int, k int) int {
	sumMap := map[int]int{}
	count := 0
	sum := 0

	for i := range arr {
		sum += arr[i]

		if sum == k {
			count++
		}

		if v, ok := sumMap[sum-k]; ok {
			count += v
		}

		sumMap[sum] += 1
	}

	return count
}

// https://leetcode.com/problems/subarray-sum-equals-k/description/
func subarraySumEqualsK() {
	arr0 := []int{1, 2, 3}
	fmt.Println(subarraySum(arr0, 3))
	arr := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(subarraySum2(arr, 0))
}
