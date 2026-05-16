package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N * 2^N)
 * ----------------------------
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(2^N)
 *    - The algorithm uses a classic "Pick or Skip" strategy. For every element in
 *      the array (of length N), it makes exactly 2 recursive calls: one including
 *      the element, and one excluding it.
 *    - This creates a binary decision tree with a maximum depth of N, resulting in
 *      up to 2^N possible states in the worst case.
 *    - (Note: The `sum > k` check acts as an optimization to prune branches early,
 *      but in the mathematical worst-case—like if `k` is very large—it still
 *      approaches 2^N).
 *
 * 2. Work Per State -> O(N)
 *    - Standard recursive branching and integer addition take O(1) time.
 *    - However, when a valid subsequence is found (`idx >= len(arr)` and `sum == k`),
 *      copying the `subsequence` slice into the `temp` slice takes up to O(N) time,
 *      since the subsequence could theoretically contain all N elements.
 *
 * Total Time = O(2^N) * O(N) = O(N * 2^N)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes exactly N levels deep before hitting
 *    the base case.
 * 2. Auxiliary Space: The `subsequence` slice holds a maximum of N integers.
 *
 * Total Space = O(N) + O(N) = O(N)
 */
func buildSubsequence(idx int, arr []int, k, sum int, subsequence []int, result *[][]int) {
	if sum > k {
		return
	}

	if idx >= len(arr) {
		if sum == k {
			temp := make([]int, len(subsequence))
			copy(temp, subsequence)

			*result = append(*result, temp)
			return
		}

		return
	}

	buildSubsequence(idx+1, arr, k, sum+arr[idx], append(subsequence, arr[idx]), result)
	buildSubsequence(idx+1, arr, k, sum, subsequence, result)

}

// https://www.geeksforgeeks.org/dsa/find-all-subsequences-with-sum-equals-to-k/
func subsequenceSumEqualsK() {
	arr := []int{0, 1, 2, 3, 4, 5}
	k := 5

	result := [][]int{}
	buildSubsequence(0, arr, k, 0, make([]int, 0, 10), &result)

	fmt.Println(result)
}
