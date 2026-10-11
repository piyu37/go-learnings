package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	nonRepeatingMap := make(map[byte]int)
	result := 0

	count := 0
	left := 0
	for i := range s {
		if _, ok := nonRepeatingMap[s[i]]; ok {
			for s[left] != s[i] {
				delete(nonRepeatingMap, s[left])
				count--
				left++
			}

			delete(nonRepeatingMap, s[i])
			count--
			left++
		}

		nonRepeatingMap[s[i]] = i

		count++

		if count > result {
			result = count
		}
	}

	return result
}

func lengthOfLongestSubstringTimeN(s string) int {
	charMap := make(map[byte]int)
	maxLongestSubstring := 0

	left := 0

	for right := range s {
		if prev, ok := charMap[s[right]]; ok {
			left = max(left, prev)
		}

		charMap[s[right]] = right + 1
		maxLongestSubstring = max(maxLongestSubstring, right-left+1)
	}

	return maxLongestSubstring
}

func lengthOfLongestSubstringTimeNOptimal(s string) int {
	var lastSeen [128]int
	left, result := 0, 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		left = max(left, lastSeen[c])
		lastSeen[c] = right + 1
		result = max(result, right-left+1)
	}

	return result
}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func longestNonRepeatingSubstring() {
	s := "abcabcbb"

	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstringTimeN(s))
	fmt.Println(lengthOfLongestSubstringTimeNOptimal(s))
}
