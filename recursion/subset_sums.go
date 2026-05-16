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
 * The total time is calculated by: (Number of states explored) * (Work done per state) + Sorting
 *
 * 1. Number of States -> O(2^N)
 *    - The algorithm uses a classic "Pick or Skip" strategy. For every element in
 *      the array (of length N), it makes exactly 2 recursive calls: one including
 *      the element's value in the sum, and one excluding it.
 *    - This creates a full binary decision tree with a maximum depth of N, resulting
 *      in exactly 2^N possible states (leaf nodes).
 *
 * 2. Work Per State -> O(1)
 *    - Standard recursive branching and integer addition take O(1) time.
 *    - At the base case (`idx == len(arr)`), appending the final sum to the
 *      `result` slice takes amortized O(1) time (as we discussed previously!).
 *    - Recursion Phase Total = O(2^N) * O(1) = O(2^N)
 *
 * 3. Sorting Phase -> O(N * 2^N)
 *    - After the recursion finishes, the `result` array contains exactly 2^N elements.
 *    - Sorting an array of size K takes O(K log K) time.
 *    - Substituting K = 2^N gives: O(2^N * log(2^N)).
 *    - Using logarithm rules, log(2^N) becomes N * log(2), which simplifies to O(N).
 *    - Therefore, the sorting step takes O(2^N * N) time.
 *
 * Total Time = Recursion O(2^N) + Sorting O(N * 2^N) = O(N * 2^N)
 *
 * SPACE COMPLEXITY: O(2^N)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes exactly N levels deep before hitting
 *    the base case, taking O(N) space.
 * 2. Auxiliary Space: The `result` slice stores the sum of every possible subset.
 *    Because there are 2^N subsets, it holds exactly 2^N integers.
 *
 * Total Space = O(N) + O(2^N) = O(2^N)
 */
func subsetSums(arr []int) []int {
	result := make([]int, 0)

	findSum(arr, 0, 0, &result)

	sort.Ints(result)

	return result
}

func findSum(arr []int, idx, sum int, result *[]int) {
	if idx == len(arr) {
		*result = append(*result, sum)
		return
	}

	findSum(arr, idx+1, sum, result)
	findSum(arr, idx+1, sum+arr[idx], result)
}

// https://www.geeksforgeeks.org/problems/subset-sums2234/1
func subsetSumsMain() {
	arr := []int{2, 3}
	result := subsetSums(arr)
	fmt.Println(result)
}
