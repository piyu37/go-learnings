package main

import (
	"fmt"
	"sort"
)

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N * 2^N)
 * ----------------------------
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(2^N)
 *    - To form subsets, the algorithm essentially makes a binary choice for every
 *      item in the array: INCLUDE it or SKIP it.
 *    - Even though it uses a `for` loop, passing `i+1` forces the algorithm to
 *      "never go backward," structurally creating a decision tree with up to 2^N
 *      unique combinations in the worst case (if all numbers are distinct).
 *    - Note: Sorting takes O(N log N) initially, but the O(2^N) subset generation
 *      completely dominates the time complexity.
 *
 * 2. Work Per State -> O(N)
 *    - The duplicate check (`nums[i] == nums[i-1]`) takes O(1) time.
 *    - At EVERY recursive call, before the loop starts, the current `subset`
 *      is copied into `temp` and appended to the result.
 *    - Since a subset can contain up to N elements, `copy(temp, subset)` takes
 *      up to O(N) time at every single state.
 *
 * Total Time = O(2^N) * O(N) = O(N * 2^N)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes at most N levels deep in the worst case.
 * 2. Auxiliary Space: The `subset` slice holds a maximum of N integers at any time.
 *
 * Total Space = O(N) + O(N) = O(N)
 */
func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	sort.Ints(nums)
	findUniqueSubsets(nums, 0, []int{}, &result)

	return result
}

func findUniqueSubsets(nums []int, idx int, subset []int, result *[][]int) {
	temp := make([]int, len(subset))
	copy(temp, subset)

	*result = append(*result, temp)

	for i := idx; i < len(nums); i++ {
		if i > idx && nums[i] == nums[i-1] {
			continue
		}

		findUniqueSubsets(nums, i+1, append(subset, nums[i]), result)
	}
}

// https://leetcode.com/problems/subsets-ii/description/
func subsetsMain() {
	nums := []int{1, 2, 2}
	fmt.Println(subsets(nums))
}
