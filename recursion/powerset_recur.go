package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N * 2^N)
 * ----------------------------
 * The total time is calculated by: (Total number of subsets generated) * (Work to build each subset)
 *
 * 1. Number of Subsets -> O(2^N)
 *    - The algorithm builds the powerset from the bottom up. At each recursive
 *      return step, it takes all existing subsets and duplicates them, adding
 *      the current array element to the new copies.
 *    - This doubling effect (1 -> 2 -> 4 -> 8...) means that for an array of
 *      length N, the final `subsets` array will contain exactly 2^N subsets.
 *
 * 2. Work Per Subset -> O(N)
 *    - Inside the `for` loop, you create a copy of the existing subset:
 *      `newSubset := append([]int{}, currentSubset...)`.
 *    - Because a subset can contain up to N elements (and the average subset
 *      contains N/2 elements), copying it into a new slice takes O(N) time.
 *    - You perform this O(N) copy operation to generate every single one of
 *      the 2^N subsets.
 *
 * Total Time = O(2^N) * O(N) = O(N * 2^N)
 *
 * SPACE COMPLEXITY: O(N * 2^N)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes exactly N levels deep (O(N)).
 * 2. Auxiliary/Output Space: The `subsets` array holds the final 2^N combinations.
 *    Since the average combination has a length of N/2, the total number of integers
 *    stored in memory across all subsets is exactly (N * 2^(N-1)).
 *    This simplifies to an overall space bound of O(N * 2^N).
 *
 * Total Space = Stack O(N) + Output O(N * 2^N) = O(N * 2^N)
 */
func PowersetRecur(array []int) [][]int {
	return powersetHelper(array, len(array)-1)
}

func powersetHelper(array []int, index int) [][]int {
	if index < 0 {
		return [][]int{{}}
	}

	subsets := powersetHelper(array, index-1)

	val := array[index]

	for j := range subsets {
		currentSubset := subsets[j]
		newSubset := append([]int{}, currentSubset...)
		newSubset = append(newSubset, val)
		subsets = append(subsets, newSubset)
	}

	return subsets
}

// This approach we follow in many recursions; like pick and not pick
// example of real world is like when we try clothes, to wear new onle we need to remove old
// following the same in below approach too

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N * 2^N)
 * ----------------------------
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(2^N)
 *    - This uses a classic "Pick or Skip" backtracking tree. For every element,
 *      it branches exactly twice (pick it, or don't pick it).
 *    - For an array of size N, this generates exactly 2^N leaf nodes (subsets).
 *
 * 2. Work Per State -> O(N)
 *    - Pushing to the slice (`append`) and popping from it are amortized O(1).
 *    - At the base case (`index >= len(array)`), you now properly allocate a new
 *      array (`make`) and `copy` the elements.
 *    - Because a subset can contain up to N elements, this copy operation takes
 *      up to O(N) time. Since this O(N) copy happens at every single one of the
 *      2^N leaf nodes, it completely dominates the time complexity.
 *
 * Total Time = O(2^N) * O(N) = O(N * 2^N)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes exactly N levels deep.
 * 2. Auxiliary Space: The `tempArr` slice holds a maximum of N integers, and is
 *    reused efficiently throughout the recursion.
 *
 * Total Space = O(N) + O(N) = O(N)
 */
func PowersetRecur2(array []int) [][]int {
	subsets := [][]int{}
	powersetHelper2(array, 0, []int{}, &subsets)
	return subsets
}

func powersetHelper2(array []int, index int, tempArr []int, subsets *[][]int) {
	if index >= len(array) {
		temp := make([]int, len(tempArr))
		copy(temp, tempArr)

		*subsets = append(*subsets, temp)
		return
	}

	tempArr = append(tempArr, array[index]) // pick
	powersetHelper2(array, index+1, tempArr, subsets)
	tempArr = tempArr[:len(tempArr)-1] // not pick
	powersetHelper2(array, index+1, tempArr, subsets)
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/25_powerset
func powersetRecurMain() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(PowersetRecur(arr))
	fmt.Println(PowersetRecur2(arr))
}
