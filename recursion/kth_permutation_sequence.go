package main

import (
	"fmt"
	"strconv"
)

func getPermutation(n int, k int) string {
	nums := ""
	for i := 1; i <= n; i++ {
		nums += strconv.Itoa(i)
	}

	return getKthPermutation(n, k-1, nums)

}

func getKthPermutation(n, k int, nums string) string {
	if n == 1 {
		return nums
	}

	nFact := 1
	for i := 2; i <= n; i++ {
		nFact *= i
	}

	partition := nFact / n
	idx := k / partition
	result := string(nums[idx])
	nums = nums[:idx] + nums[idx+1:]
	k = k % partition

	result += getKthPermutation(n-1, k, nums)

	return result
}

// https://leetcode.com/problems/permutation-sequence/description/
func kthPermutationSequence() {
	fmt.Println(getPermutation(4, 9))
	fmt.Println(getPermutation(4, 17))
	fmt.Println(getPermutation(3, 3))
}
