package main

import (
	"fmt"
)

func Solution(N int, K int) string {
	if N < 1 || N > 2001 || K < 1 || K > 26 {
		return ""
	}

	palindrome := make([]byte, N)

	for i := 0; i < (N+1)/2; i++ {
		palindrome[i] = byte('a' + i%K)
		palindrome[N-i-1] = palindrome[i]
	}

	// for i := (N - 1) / 2; i >= 0; i-- {
	// 	palindrome[N-i-1] = palindrome[i]
	// }

	return string(palindrome)
}

// Create palindrome string of l length using k characters starting from a
func skillikzCreatePalindrome() {
	l := 13
	k := 7
	fmt.Println(Solution(l, k))
}
