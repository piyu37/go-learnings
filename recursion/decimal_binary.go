package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(log V)
 * ----------------------------
 * Where V is the numerical value of the input integer `v`.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(log V)
 *    - At every recursive step, the input `v` is divided by 2.
 *    - The number of times you can halve a number before it reaches 0 is
 *      mathematically defined as the base-2 logarithm of that number: log₂(V).
 *    - Therefore, the recursion tree is just a single straight line that is
 *      exactly log₂(V) nodes long.
 *
 * 2. Work Per State -> O(1)
 *    - Inside each recursive call, the algorithm performs basic integer arithmetic:
 *      modulo (`v % 2`), division (`v / 2`), multiplication (`10 * ...`), and addition.
 *    - Because these are primitive mathematical operations on integers, they
 *      execute in constant O(1) time.
 *
 * Total Time = O(log V) * O(1) = O(log V)
 *
 * SPACE COMPLEXITY: O(log V)
 * ----------------------------
 * 1. Recursion Stack: The maximum depth of the call stack matches the number of
 *    recursive calls. Since the algorithm continuously halves `v` until it hits 0,
 *    the call stack grows to exactly log₂(V) levels deep.
 * 2. Auxiliary Space: No additional data structures (like slices, arrays, or maps)
 *    are created. The memory footprint per stack frame is strictly limited to
 *    primitive integers, taking O(1) extra space.
 *
 * Total Space = Stack O(log V) + Auxiliary O(1) = O(log V)
 */
func convertDecimalToBinary(v int) int {
	if v < 1 {
		return v
	}

	mod := v % 2

	return 10*convertDecimalToBinary(v/2) + mod
}

// convert decimal to binary
func decimalBinary() {
	fmt.Println(convertDecimalToBinary(7))
}
