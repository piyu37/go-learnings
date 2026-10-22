package main

import "fmt"

func characterReplacement(s string, k int) int {
	l, r, maxLen, maxFreq := 0, 0, 0, 0
	var charMap [26]int

	for r < len(s) {
		charMap[s[r]-'A']++
		maxFreq = max(maxFreq, charMap[s[r]-'A'])

		if r-l+1-maxFreq > k {
			charMap[s[l]-'A']--
			l++
		}

		if r-l+1-maxFreq <= k {
			maxLen = max(maxLen, r-l+1)
		}

		r++
	}

	return maxLen
}

// https://leetcode.com/problems/longest-repeating-character-replacement/description/
func longestRepeatingCharReplacement() {
	s := "ABAB"
	k := 2

	fmt.Println(characterReplacement(s, k))
}
