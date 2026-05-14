package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N * 2^N)
 * ----------------------------
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(2^N)
 *    - In the worst-case scenario (e.g., a string of all identical characters like "aaaa"),
 *      every single substring is a valid palindrome.
 *    - For a string of length N, there are exactly N-1 positions where you could
 *      potentially place a "split" or "partition" to divide the string.
 *    - Since you have a binary choice (split or don't split) at each of these N-1
 *      positions, there are 2^(N-1) possible partitions, which simplifies to O(2^N).
 *
 * 2. Work Per State -> O(N)
 *    - At every step in the `for` loop, `isPalindrome()` takes up to O(N) time to
 *      validate the substring.
 *    - Creating the substring `string(s[idx:j+1])` takes up to O(N) time.
 *    - Finally, when a valid full partition is found (`len(s) == idx`), copying
 *      the `substr` slice into `temp` takes O(N) time.
 *
 * Total Time = O(2^N) * O(N) = O(N * 2^N)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes at most N levels deep (this happens
 *    when the string is partitioned into individual characters).
 * 2. Auxiliary Space: The `substr` slice holds a maximum of N strings at any time,
 *    and the total length of all those strings combined will always be exactly N.
 *
 * Total Space = O(N) + O(N) = O(N)
 */
func partition(s string) [][]string {
	result := make([][]string, 0)

	checkPalindromePartitioning(s, 0, []string{}, &result)

	return result
}

func checkPalindromePartitioning(s string, idx int, substr []string, result *[][]string) {
	if len(s) == idx {
		temp := make([]string, len(substr))
		copy(temp, substr)

		*result = append(*result, temp)
		return
	}

	for j := idx; j < len(s); j++ {
		if isPalindrome(s, idx, j) {
			checkPalindromePartitioning(s, j+1, append(substr, string(s[idx:j+1])), result)
		}
	}
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

// https://leetcode.com/problems/palindrome-partitioning/
func palindromePartitioning() {
	s := "aabb"
	fmt.Println(partition(s))
}
