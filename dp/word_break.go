package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	return wordBreakRecursion(s, wordDict, 0, map[int]bool{})
}

func wordBreakRecursion(s string, wordDict []string, start int, memoization map[int]bool) bool {
	if start == len(s) {
		return true
	}

	if memoization[start] {
		return memoization[start]
	}

	for end := start; end < len(s); end++ {
		for _, w := range wordDict {
			if len(w)-1 > end {
				continue
			}

			exist, ok := memoization[end+1]
			if s[start:end+1] == w && (ok || wordBreakRecursion(s, wordDict, end+1, memoization)) {
				if ok {
					if exist {
						memoization[start] = true
						return true
					}

					memoization[start] = false
					continue
				}

				memoization[start] = true
				return true
			}
		}
	}

	memoization[start] = false

	return false
}

// https://leetcode.com/problems/word-break/description/?envType=study-plan-v2&envId=top-interview-150
func wordBreakMain() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))
}
