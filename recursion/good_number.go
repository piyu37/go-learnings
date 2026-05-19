package main

import (
	"fmt"
)

const mod = 1000000007

func pow(x, y int) int {
	if y == 0 {
		return 1
	}

	ans := pow(x, y/2)
	ans *= ans
	ans %= mod

	if y%2 != 0 {
		ans *= x
	}

	ans %= mod

	return ans

}

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(log N)
 * ----------------------------
 * Where N is the input number `n`.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(log N)
 *    - The `countGoodNumbers` function makes two calls to the `pow` function,
 *      where the exponent `y` is roughly N/2.
 *    - The `pow` function uses a technique called Binary Exponentiation
 *      (or Fast Exponentiation). Instead of multiplying `x` by itself `y` times
 *      (which would take O(N) steps), it recursively divides the exponent `y`
 *      by 2 at each step (`pow(x, y/2)`).
 *    - A number can only be divided by 2 approximately log₂(N) times before
 *      it reaches 0. Therefore, the recursion tree is strictly a single path
 *      going down log₂(N) levels.
 *
 * 2. Work Per State -> O(1)
 *    - Inside each recursive call of `pow`, the algorithm performs a few
 *      basic arithmetic operations: multiplication (`*=`), modulo (`%=`),
 *      and an `if` condition check.
 *    - All of these operations execute in constant time regardless of how
 *      large N gets.
 *
 * Total Time = O(log N) unique states * O(1) work per state = O(log N)
 *
 * SPACE COMPLEXITY: O(log N)
 * ----------------------------
 * 1. Recursion Stack: Because the `pow` function recursively calls itself by
 *    halving the exponent, the maximum depth of the call stack is equal to the
 *    number of divisions by 2. This means the stack goes exactly log₂(N) levels
 *    deep before hitting the base case (`y == 0`). This takes O(log N) space.
 * 2. Auxiliary Space: No additional data structures (like arrays or maps) are
 *    created during the execution.
 *
 * Total Space = Stack O(log N) + Auxiliary O(1) = O(log N)
 */
func countGoodNumbers(n int64) int {
	even := (n + 1) / 2
	odd := n / 2

	return pow(5, int(even)) * pow(4, int(odd)) % mod
}

// https://leetcode.com/problems/count-good-numbers/
func goodNumber() {
	fmt.Println(countGoodNumbers(4))
}
