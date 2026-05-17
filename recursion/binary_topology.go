package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N^2)
 * ----------------------------
 * The total time is calculated by finding the exact aggregate work done
 * across all memoized states, rather than a simple multiplication.
 *
 * 1. Number of Unique States -> O(N)
 *    - Thanks to the `cache` map (memoization), the algorithm guarantees that it
 *      never computes the topology count for the same number of nodes twice.
 *    - It evaluates exactly N+1 unique states (from 0 up to N).
 *
 * 2. Variable Work Per State -> 1 to N iterations
 *    - When a state is NOT in the cache, the algorithm executes the loop
 *      `for i := range n`.
 *    - The amount of work is NOT a static O(N) for every state. Instead, it grows:
 *      - `calculateTopology(1)` loops 1 time.
 *      - `calculateTopology(2)` loops 2 times.
 *      - ...
 *      - `calculateTopology(N)` loops N times.
 *
 * 3. Total Aggregate Work -> O(N^2)
 *    - To find the exact total time across the entire program execution, we
 *      sum the loop iterations for all states:
 *      Total Iterations = 1 + 2 + 3 + ... + N
 *    - In mathematics, the sum of this arithmetic series is exacty: N(N+1)/2
 *    - Expanding this gives: (N^2)/2 + N/2.
 *    - Because Big-O notation drops constants (like the /2) and non-dominant
 *      terms (like the + N/2), the final complexity simplifies to exactly O(N^2).
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: The maximum depth of the call stack is N. This occurs
 *    because evaluating `n` requires evaluating `n-1`, which requires `n-2`,
 *    all the way down to 0 before the first branch fully resolves.
 * 2. Auxiliary Space: The `cache` map stores exactly N+1 integer key-value
 *    pairs (from 0 to N) to prevent redundant calculations.
 *
 * Total Space = Stack O(N) + Cache O(N) = O(N)
 */
func NumberOfBinaryTreeTopologies(n int) int {
	return calculateTopology(n, map[int]int{
		0: 1,
		1: 1,
	})
}

func calculateTopology(n int, cache map[int]int) int {
	if v, ok := cache[n]; ok {
		return v
	}

	var count int

	for i := range n {
		left := calculateTopology(i, cache)
		right := calculateTopology(n-i-1, cache)

		count += left * right
	}

	cache[n] = count

	return count
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/29_number_of_binary_tree_topologies
func binaryTopology() {
	fmt.Println(NumberOfBinaryTreeTopologies(3))
}
