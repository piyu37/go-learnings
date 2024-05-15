package main

import (
	"fmt"
)

func longestKSubstr(s string, k int) int {
	i := 0
	j := 0
	ans := -1
	mp := make(map[byte]int)
	for j < len(s) {
		mp[s[j]]++
		for len(mp) > k {
			mp[s[i]]--
			if mp[s[i]] == 0 {
				delete(mp, s[i])
			}
			i++
		}
		if len(mp) == k {
			ans = maxBetweenTwo(ans, j-i+1)
		}
		j++
	}
	return ans
}

func maxBetweenTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://www.geeksforgeeks.org/find-the-longest-substring-with-k-unique-characters-in-a-given-string/
func longestSubstringKUniqueChar() {
	s := "aabacbebebe"
	k := 3
	result := longestKSubstr(s, k)
	fmt.Println("Length of longest substring with at most", k, "distinct characters:", result)
}
