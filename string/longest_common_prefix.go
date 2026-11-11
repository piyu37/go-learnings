package main

import "fmt"

// Vertical scanning: compare column by column across all strings,
// stop at the first mismatch or when any string runs out.
// Time: O(n * (len(prefix) + 1)), worst case O(S) where S = total chars. Space: O(1).
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for _, s := range strs[1:] {
			if i == len(s) || s[i] != c {
			return strs[0][:i]
			}
		}
	}

	return strs[0]
}

// https://leetcode.com/problems/longest-common-prefix/description/
func longestCommonPrefixMain() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs)) // Output: fl

	strs = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs)) // Output: ""

	strs = []string{"interview", "internet", "interval", "internal"}
	fmt.Println(longestCommonPrefix(strs)) // Output: inter
}
