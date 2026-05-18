package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY:
 * Tight & Accurate Bound: O(k * C(n, k))  <-- (Use this one!)
 * Loose Upper Bound: O(k * n^k)
 * ----------------------------
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(C(n, k)) vs O(n^k)
 *    - Accurate Bound (C(n, k)): Because the algorithm passes `i+1` into the recursive
 *      call, it strictly moves forward and never revisits elements or creates duplicate
 *      sets. The exact number of valid combinations generated is mathematically
 *      "N choose K", written as C(n, k) or n! / (k! * (n-k)!). This is the most
 *      accurate representation of the algorithm's execution.
 *    - Loose Bound (n^k): If we ignore the mathematical pruning and just look at the
 *      raw shape of the recursion tree, it goes `k` levels deep, with a maximum
 *      branching factor of roughly `n`. This gives a theoretical ceiling of n^k states.
 *      However, this is a mathematically lazy bound that vastly overestimates the work.
 *
 * 2. Work Per State -> O(k)
 *    - At every recursive step, updating the array takes O(1) time.
 *    - At the base case (`idx == k`), the algorithm allocates a new `temp` slice
 *      and uses `copy()` to freeze the result.
 *    - Copying exactly `k` elements takes O(k) time and happens at every valid leaf node.
 *
 * Total Time:
 * Most Accurate: O(C(n, k)) * O(k) = O(k * C(n, k))
 * Mathematically Lazy: O(n^k) * O(k) = O(k * n^k)
 *
 * SPACE COMPLEXITY: O(n)
 * ----------------------------
 * 1. Recursion Stack: The call stack depth is exactly bounded by `k` levels.
 * 2. Auxiliary Space: The initial `nums` array consumes O(n) space, and the reusable
 *    `currCombination` array consumes O(k) space. Since k is always <= n, O(n) is
 *    the dominant upper bound.
 *
 * Total Space = Stack O(k) + Arrays O(n) = O(n)
 */
func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	nums := make([]int, n)

	for i := range nums {
		nums[i] = i + 1
	}

	currCombination := make([]int, k)

	createCombinations(k, 0, 0, nums, currCombination, &result)

	return result
}

func createCombinations(k, idx, numIdx int, nums, currCombination []int,
	result *[][]int) {
	if idx == k {
		temp := make([]int, k)
		copy(temp, currCombination)
		*result = append(*result, temp)
		return
	}

	for i := numIdx; i < len(nums); i++ {
		currCombination[idx] = nums[i]
		createCombinations(k, idx+1, i+1, nums, currCombination, result)
	}
}


// Same as above in terms of worst case
func combineOptimized(n int, k int) [][]int {
	result := make([][]int, 0)
	currCombination := make([]int, k)
	createCombinations2(n, k, 0, 1, currCombination, &result)

	return result
}

func createCombinations2(n, k, currIdx, currNum int, currCombination []int, result *[][]int) {
	if currIdx == k {
		temp := make([]int, k)
		copy(temp, currCombination)
		*result = append(*result, temp)
		return
	}

	need := k - currIdx
	remain := n - currNum + 1
	available := remain - need

	for i := currNum; i <= currNum+available; i++ {
		currCombination[currIdx] = i
		createCombinations2(n, k, currIdx+1, i+1, currCombination, result)
	}
}

// https://leetcode.com/problems/combinations/description/?envType=study-plan-v2&envId=top-interview-150
func combinations() {
	n := 4
	k := 3

	fmt.Println(combine(n, k))

	fmt.Println(combineOptimized(n, k))
}
