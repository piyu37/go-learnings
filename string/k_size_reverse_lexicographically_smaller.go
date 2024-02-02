package main

import (
	"fmt"
	"strings"
)

func countNumWays(s string, k int) int32 {
	var count int32 = 0
	n := len(s)

	for i := 0; i+k <= n; i++ {
		substr := s[i : i+k]
		reversed := reverse(substr)

		if strings.Compare(reversed, substr) < 0 {
			count++
		}
	}

	return count
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// countNumWaysMain counts the number of substrings of a given length for which the
// reversed version is lexicographically smaller than the original substring
// e.g. ababa => ab => reverse(ab) = ba => false; not lexicographically small
// ba => reverse(ba) => ab => true
// ab => reverse(ab) = ba => false
// ba => reverse(ba) => ab => true
// hence output = 2
func kSizeReverseLexicographicallySmaller() {
	// Example 1
	s1 := "ababa"
	k1 := 2
	output1 := countNumWays(s1, k1)
	fmt.Printf("Output 1: %d\n", output1)

	// Example 2
	s2 := "amazon"
	k2 := 3
	output2 := countNumWays(s2, k2)
	fmt.Printf("Output 2: %d\n", output2)
}
