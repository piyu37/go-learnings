package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	substrIdx := 0
	for i := range t {
		if s[substrIdx] == t[i] {
			substrIdx++
		}

		if len(s) == substrIdx {
			return true
		}
	}

	return len(s) == substrIdx
}

// https://leetcode.com/problems/is-subsequence/?envType=study-plan-v2&envId=top-interview-150
func isSubsequenceMain() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
