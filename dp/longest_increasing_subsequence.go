package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	msis := make([]int, len(nums))
	maxLen := 1

	for i := range msis {
		msis[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && msis[i] < 1+msis[j] {
				msis[i] = 1 + msis[j]
				if maxLen < msis[i] {
					maxLen = msis[i]
				}
			}
		}
	}

	return maxLen
}

func lengthOfLISBetter(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	msis := make([]int, 0)
	msis = append(msis, nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] > msis[len(msis)-1] {
			msis = append(msis, nums[i])
		} else {
			j := 0
			for nums[i] > msis[j] {
				j++
			}

			msis[j] = nums[i]
		}
	}

	return len(msis)
}

func lengthOfLISOptimalUsingBS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	msis := make([]int, 0)
	msis = append(msis, nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] > msis[len(msis)-1] {
			msis = append(msis, nums[i])
		} else {
			j := search(msis, nums[i])
			msis[j] = nums[i]
		}
	}

	return len(msis)
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return left
}

// https://leetcode.com/problems/longest-increasing-subsequence/description/?envType=study-plan-v2&envId=top-interview-150
func longestIncreasingSubsequence() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println(lengthOfLIS(nums))

	nums = []int{8, 1, 6, 9, 2, 3, 10}

	fmt.Println(lengthOfLISBetter(nums))

	fmt.Println(lengthOfLISOptimalUsingBS(nums))
}
