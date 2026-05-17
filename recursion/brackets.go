package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N * 2^(2N))
 * ----------------------------
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(2^(2N))
 *    - At every step, the algorithm makes up to 2 choices: add '(' or add ')'.
 *    - To form a complete valid sequence, it must make exactly 2N choices (N opens, N closes).
 *    - If unpruned, the decision tree would branch into a maximum of 2^(2N) states.
 *    - (Note: The conditions `open < n` and `close < open` aggressively prune invalid
 *      paths, bringing the actual visited states down to the N-th Catalan number,
 *      but 2^(2N) is standard for the upper bound).
 *
 * 2. Work Per State -> O(N)
 *    - In Go, strings are immutable. Doing `str+"("` or `str+")"` forces the program
 *      to allocate a brand new string and copy the existing characters over.
 *    - Since the string grows to a maximum length of 2N, this copy operation takes
 *      up to O(N) time at every single recursive step.
 *
 * Total Time = O(2^(2N)) * O(N) = O(N * 2^(2N))
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes exactly 2N levels deep (one for each character).
 * 2. Auxiliary Space: String creation takes extra memory per frame, but the footprint
 *    at any given maximum depth is bounded by O(N).
 *
 * Total Space = Stack O(N) + Extra Space O(N) = O(N)
 */
func printProperBrackets(n int) []string {
	result := make([]string, 0)

	printBrackets("", n, 0, 0, &result)

	return result
}

func printBrackets(str string, n, open, close int, result *[]string) {
	if open == n && close == n {
		*result = append(*result, str)
		return
	}

	if open < n {
		printBrackets(str+"(", n, open+1, close, result)
	}

	if close < open {
		printBrackets(str+")", n, open, close+1, result)
	}
}

// https://www.geeksforgeeks.org/print-all-combinations-of-balanced-parentheses/
func brackets() {
	n := 3
	fmt.Println(printProperBrackets(n))
}
