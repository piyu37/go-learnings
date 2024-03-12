package main

import "fmt"

// when using extra space
func rotateArrayByKWithExtraSpace() {
	var (
		N, K int = 6, 2
	)
	// fmt.Scan(&T)
	// fmt.Scan(&N, &K)
	arr := []int{1, 2, 3, 4, 5, 6}

	// for i := range arr {
	// 	fmt.Scan(&arr[i])
	// }

	for i := range arr {
		fmt.Printf("%d ", arr[(i+(N-K))%N])
	}
}

func rotateArrayByKWithoutExtraSpace(nums []int, k int) {
	numsLen := len(nums)

	if numsLen < 2 {
		return
	}

	k = k % numsLen

	left := 0
	right := numsLen - k - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	left = numsLen - k
	right = numsLen - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	left = 0
	right = numsLen - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// https://leetcode.com/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
func rotate1DArrayByKMain() {
	rotateArrayByKWithExtraSpace()
	arr := []int{1, 2, 3, 4, 5, 6}
	k := 2
	rotateArrayByKWithoutExtraSpace(arr, k)
	fmt.Println(arr)
}
