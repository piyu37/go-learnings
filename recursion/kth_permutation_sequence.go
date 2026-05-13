package main

import (
	"fmt"
	"strconv"
)

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N^2)
 * ----------------------------
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(N)
 *    - The algorithm builds the final permutation one digit at a time.
 *    - Since it locks in one digit per recursive call and decreases `n` by 1
 *      each time, it makes exactly N recursive calls.
 *
 * 2. Work Per State -> O(N)
 *    - Calculating the factorial (`nFact`) using the `for` loop takes up to
 *      O(N) time at each step.
 *    - Removing the selected character and concatenating the remaining string
 *      (`nums[:idx] + nums[idx+1:]`) creates a new string in Go, which requires
 *      copying characters and takes O(N) time.
 *
 * Total Time = O(N) * O(N) = O(N^2)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes exactly N levels deep (one for each digit).
 * 2. Auxiliary Space: The `nums` string and `result` string hold a maximum of N characters.
 *
 * Total Space = O(N) + O(N) = O(N)
 */
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
