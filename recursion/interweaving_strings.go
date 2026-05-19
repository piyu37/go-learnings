package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(2^(N+M))
 * ----------------------------
 * Where N is the length of string `one` and M is the length of string `two`.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(2^(N+M))
 *    - This is an unoptimized recursive solution (no memoization).
 *    - At each step, if the current character in `three` matches the current
 *      character in both `one` and `two`, the algorithm branches into two
 *      separate recursive paths.
 *    - The recursion tree continues until it forms the full string `three`,
 *      meaning the maximum depth of the tree is N + M.
 *    - A recursion tree with a depth of N+M and a branching factor of 2 gives
 *      us an upper bound of 2^(N+M) possible states in the worst-case scenario
 *      (e.g., when all strings consist of the same repeating character).
 *
 * 2. Work Per State -> O(1)
 *    - Inside each recursive call, the algorithm only performs basic integer
 *      addition (`k := i + j`), array indexing (`one[i] == three[k]`), and
 *      boolean evaluations.
 *    - Because it passes indexes (`i` and `j`) rather than creating new string
 *      substrings, these operations take constant O(1) time.
 *
 * Total Time = O(2^(N+M)) states * O(1) work per state = O(2^(N+M))
 *
 * SPACE COMPLEXITY: O(N+M)
 * ----------------------------
 * 1. Recursion Stack: In the worst-case scenario, the algorithm successfully
 *    matches characters one by one until it reaches the end of `three`. The
 *    call stack will go exactly N + M levels deep before hitting the base case
 *    (`k == len(three)`). This requires O(N+M) memory for the stack frames.
 * 2. Auxiliary Space: The algorithm passes strings by reference and uses primitive
 *    integer indices (`i`, `j`, `k`). It does not allocate any new maps, arrays,
 *    or strings, taking O(1) extra space.
 *
 * Total Space = Stack O(N+M) + Auxiliary O(1) = O(N+M)
 */
func areInterweaven(one, two, three string, i, j int) bool {
	k := i + j

	if k == len(three) {
		return true
	}

	if i < len(one) && one[i] == three[k] {
		if areInterweaven(one, two, three, i+1, j) {
			return true
		}
	}

	if j < len(two) && two[j] == three[k] {
		return areInterweaven(one, two, three, i, j+1)
	}

	return false
}

func InterweavingStrings(one, two, three string) bool {
	if len(three) != len(one)+len(two) {
		return false
	}

	return areInterweaven(one, two, three, 0, 0)
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/30_interweaving_strings
func interweavingStringsMain() {
	one := "algoexpert"
	two := "your-dream-job"
	three := "your-algodream-expertjob"

	fmt.Println(InterweavingStrings(one, two, three))
}
