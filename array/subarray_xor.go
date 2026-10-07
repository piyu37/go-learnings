package main

import "fmt"

func subarrayXor(N, K int, A []int) int {
	minLen := N + 1
	result := -1
	for i := 0; i < N; i++ {
		x := 0
		for j := i; j < N; j++ {
			x ^= A[j]
			if x >= K {
				L := j - i + 1
				if L < minLen {
					minLen = L
					result = x
				} else if L == minLen && x > result {
					result = x
				}
				break // longer subarrays from this i can't be shorter
			}
		}
	}
	return result
}

// Array Operations: Subarray XOR (Golang Coding)Problem Statement:
// You are given one array of integers of size N and one integer K. You must find the subarray with a minimum size with a
// bitwise XOR of all the values of the subarray greater than or equal to K. Find the bitwise XOR of all elements in this subarray.
// If there is no array with a bitwise XOR greater than or equal to K, print -1. If multiple subarrays with minimum size have a
// bitwise XOR greater than or equal to K, print the maximum XOR value from them.

// Constraints:2 <= N <= 10^5   0 <= A[i] <= 10^9   0 <= K <= 10^9
// Sample:Input: 5 7 (denotes N and K) \n 3 4 2 3 5 (denotes A[i])
// Output: 7
// Explanation: For array [3, 4, 2, 3, 5], the smallest length subarray with a bitwise XOR >= 7 is [3, 4]. The XOR of 3 and 4 is 7.

// Similar Concept: This involves determining contiguous XOR values rapidly, which requires storing Prefix XORs mapped
// within a Trie (Prefix Tree). The closest conceptual LeetCode problem is LeetCode 1803: Count Pairs With XOR in a Range.

// This problem can be solved via Prefix XOR with a Trie. DO it when you learn/solve about trie questions
// Ericsson Hackerrank test
func subarrayXORMain() {
	N, K := 5, 7
	A := []int{3, 4, 2, 3, 5}

	fmt.Println(subarrayXor(N, K, A))
}
