package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N * N!)
 * ----------------------------
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(N!)
 *    - The algorithm generates permutations by swapping elements in-place.
 *    - For the first position, the loop runs N times (N choices). For the second
 *      position, it runs N-1 times, and so on down to 1 choice for the last position.
 *    - This creates a decision tree with exactly N * (N-1) * (N-2) * ... * 1 = N!
 *      leaf nodes, representing all valid permutations.
 *
 * 2. Work Per State -> O(N)
 *    - The in-place swapping (`array[i], array[j] = ...`) takes O(1) time.
 *    - However, at the base case (`len(array)-1 == i`), a complete permutation is
 *      formed. Copying the `array` into the `temp` slice takes O(N) time.
 *    - Because this O(N) copy happens exactly N! times (at every leaf node), it
 *      dominates the time complexity.
 *
 * Total Time = O(N!) * O(N) = O(N * N!)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes exactly N levels deep (one for each index).
 * 2. Auxiliary Space: The `temp` slice holds exactly N integers during the copy operation.
 *    (Note: Because you use in-place swapping, you avoid creating new arrays at
 *    every recursive step, keeping the memory overhead incredibly lean!)
 *
 * Total Space = O(N) + O(N) = O(N)
 */
func GetPermutations(array []int) [][]int {
	if len(array) == 0 {
		return [][]int{}
	}

	perms := make([][]int, 0)
	helper(0, array, &perms)

	return perms
}

func helper(i int, array []int, perms *[][]int) {
	if len(array)-1 == i {
		temp := make([]int, len(array))
		copy(temp, array)
		*perms = append(*perms, temp)

		return
	}

	for j := i; j < len(array); j++ {
		array[i], array[j] = array[j], array[i]

		helper(i+1, array, perms)

		array[i], array[j] = array[j], array[i]
	}
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/24_permutations
func permutation() {
	arr := []int{1, 2, 3}

	fmt.Println(GetPermutations(arr))
}
