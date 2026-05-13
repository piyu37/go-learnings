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
 *    - To form unique combinations, the algorithm fundamentally makes a binary
 *      choice for every single item in the array: INCLUDE it or SKIP it.
 *    - 2 choices for the 1st item * 2 choices for the 2nd ... (N times) = 2^N subsets.
 *    - Even though it uses a `for` loop, passing `i+1` forces the algorithm to
 *      "never go backward," which structurally creates this exact 2^N decision tree.
 *    - Note: Sorting takes O(N log N) initially, but the O(2^N) subset generation
 *      completely dominates the time complexity.
 *
 * 2. Work Per State -> O(N)
 *    - The standard recursive branching and the duplicate check take O(1) time.
 *    - When a valid combination is found (`target == 0`), copying the `combination`
 *      slice into the `temp` slice takes up to O(N) time, since a combination can
 *      contain at most N elements.
 *
 * Total Time = O(2^N) * O(N) = O(N * 2^N)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes at most N levels deep in the worst case.
 * 2. Auxiliary Space: The `combination` slice holds a maximum of N integers.
 *
 * Total Space = O(N) + O(N) = O(N)
 */
func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)

	var findCombination func(candidates []int, target int, idx int, combination []int)

	findCombination = func(candidates []int, target int, idx int, combination []int) {
		if target == 0 {
			temp := make([]int, len(combination))
			copy(temp, combination)
			result = append(result, temp)

			return
		}

		if target < 0 {
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}

			findCombination(candidates, target-candidates[i], i+1, append(combination, candidates[i]))
		}
	}

	findCombination(candidates, target, 0, []int{})

	return result
}

// https://leetcode.com/problems/combination-sum-ii/description/
func combinationSum2Main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	fmt.Println(combinationSum2(candidates, target))
}
