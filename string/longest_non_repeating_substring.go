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

func longestNonRepeatingSubstring() {
	s := "abcabcbb"

	fmt.Println(lengthOfLongestSubstring(s))
}
