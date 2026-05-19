package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N * M)
 * ----------------------------
 * Where N is the length of string `one` and M is the length of string `two`.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(N * M)
 *    - The state in this algorithm is defined by the combination of the two
 *      indices `i` and `j`.
 *    - Because of the `cache` (memoization matrix), the algorithm guarantees
 *      that it will compute the result for any specific `(i, j)` pair exactly once.
 *    - The index `i` can range from 0 to N, and `j` can range from 0 to M.
 *      Therefore, there are exactly (N+1) * (M+1) unique states that could
 *      possibly be evaluated.
 *
 * 2. Work Per State -> O(1)
 *    - Inside each recursive call (when not returning a cached result), the
 *      algorithm performs constant time operations: addition (`k := i + j`),
 *      bounds checking, character comparisons (`one[i] == three[k]`), and
 *      assigning boolean pointers to the cache.
 *    - Since the strings are not being sliced or copied—only indexed—the
 *      work done to process a single state is O(1).
 *
 * Total Time = O(N * M) unique states * O(1) work per state = O(N * M)
 *
 * SPACE COMPLEXITY: O(N * M)
 * ----------------------------
 * 1. Recursion Stack: In the worst-case path, the algorithm increments either
 *    `i` or `j` one at a time until it reaches the end of `three`. This means
 *    the maximum depth of the call stack is exactly N + M levels deep.
 *    This consumes O(N + M) space.
 * 2. Auxiliary Space: The algorithm explicitly allocates a 2D `cache` matrix
 *    of dimensions (N+1) by (M+1) to store pointers to booleans. This matrix
 *    dominates the auxiliary memory, consuming O(N * M) space.
 *
 * Total Space = Stack O(N + M) + Cache O(N * M) = O(N * M)
 */
func areInterweavenCache(one, two, three string, i, j int, cache [][]*bool) bool {
	if cache[i][j] != nil {
		return *cache[i][j]
	}

	k := i + j

	if k == len(three) {
		return true
	}

	if i < len(one) && one[i] == three[k] {
		result := areInterweavenCache(one, two, three, i+1, j, cache)
		cache[i][j] = &result
		if result {
			return true
		}
	}

	if j < len(two) && two[j] == three[k] {
		result := areInterweavenCache(one, two, three, i, j+1, cache)
		cache[i][j] = &result

		return result
	}

	result := false
	cache[i][j] = &result

	return result
}

func InterweavingStringsCache(one, two, three string) bool {
	if len(three) != len(one)+len(two) {
		return false
	}

	cache := make([][]*bool, len(one)+1)

	for i := 0; i < len(one)+1; i++ {
		cache[i] = make([]*bool, len(two)+1)
	}

	return areInterweavenCache(one, two, three, 0, 0, cache)
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/30_interweaving_strings
func interweavingStringsCacheMain() {
	one := "algoexpert"
	two := "your-dream-job"
	three := "your-algodream-expertjob"

	fmt.Println(InterweavingStringsCache(one, two, three))
}
