package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N * K)
 * ----------------------------
 * Where N is the height of the staircase `h` and K is `maxSteps`.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(N)
 * - The state is defined by the current step (or remaining height) `h`.
 * - Thanks to the `cache` map (memoization), the algorithm guarantees that it
 * will only fully evaluate the number of ways to climb from a specific step
 * exactly once.
 * - Since `h` decreases down to 0, there are exactly N unique states that
 * need to be computed.
 *
 * 2. Work Per State -> O(K)
 * - Inside each recursive call (when the result is not already cached), the
 * algorithm runs a `for` loop that iterates exactly `maxSteps` (K) times.
 * - Inside this loop, it performs constant time O(1) operations: recursive calls
 * (which immediately return from the cache if already computed), map insertions,
 * and integer addition.
 * - Because it loops K times, the work done to compute a single new state is O(K).
 *
 * Total Time = O(N) unique states * O(K) work per state = O(N * K)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: In the worst-case scenario (always exploring the `h - 1`
 * branch first), the recursion tree goes exactly N levels deep before hitting
 * the base case (`h < 0` or a pre-seeded `h == 0`). This requires O(N) space
 * for the call stack.
 * 2. Auxiliary Space: The `cache` map stores the result for every single step
 * from 1 up to N. This mapping of step-to-result consumes O(N) memory.
 *
 * Total Space = Stack O(N) + Cache Map O(N) = O(N)
 */
func StaircaseTraversal(height int, maxSteps int) int {

	return staircaseHelper(height, maxSteps, map[int]int{0: 1, 1: 1})
}

func staircaseHelper(h, maxSteps int, cache map[int]int) (n int) {
	if h < 0 {
		return 0
	}

	if v, ok := cache[h]; ok {
		return v
	}

	for i := 1; i <= maxSteps; i++ {
		output := staircaseHelper(h-i, maxSteps, cache)

		cache[h-i] = output

		n += output
	}

	return n
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/45_staircase_traversal
func staircaseTraversalMain() {
	fmt.Println(StaircaseTraversal(5, 3))
}
