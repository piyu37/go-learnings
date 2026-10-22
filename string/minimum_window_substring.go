package main

import "fmt"

func minWindow(s string, t string) string {
	tMap := make(map[byte]int, len(t))
	minLen := len(s)
	minSubstring := ""

	tLen := len(t)
	for i := range t {
		tMap[t[i]] += 1
	}

	i, j := 0, 0
	count := 0
	countMap := make(map[byte]int, len(t))
	for i < len(s) {

		for j < len(s) && count < tLen {
			if tMap[s[j]] > 0 {
				countMap[s[j]] += 1

				if countMap[s[j]] <= tMap[s[j]] {
					count++
				}
			}

			j++
		}

		if count == tLen {
			if j-i <= minLen {
				minSubstring = s[i:j]
				minLen = j - i
			}
		}

		if _, ok := countMap[s[i]]; ok {
			if countMap[s[i]] <= tMap[s[i]] {
				count--
			}

			countMap[s[i]] -= 1
		}

		i++
	}

	return minSubstring
}

func minWindowUsingSingleMap(s string, t string) string {
	if len(t) == 0 || len(t) > len(s) {
		return ""
	}

	// need[c] > 0  -> still need that many more of c
	// need[c] == 0 -> exactly satisfied
	// need[c] < 0  -> window has extra copies of c
	need := make(map[int32]int, len(t))
	for i := 0; i < len(t); i++ {
		need[int32(t[i])]++
	}

	missing := len(t) // total chars of t still missing from the window
	start, minLen := 0, len(s)+1

	i := 0
	for j := 0; j < len(s); j++ {
		// expand: s[j] enters the window
		c := int32(s[j])
		if n, ok := need[c]; ok { // ok == false -> c not in t, skip
			if n > 0 {
				missing--
			}
			need[c] = n - 1
		}

		// shrink while the window is valid
		for missing == 0 {
			if j-i+1 < minLen {
				start, minLen = i, j-i+1
			}

			// s[i] leaves the window
			c := int32(s[i])
			if n, ok := need[c]; ok {
				if n == 0 { // was exactly satisfied, now it's missing
					missing++
				}
				need[c] = n + 1
			}
			i++
		}
	}

	if minLen > len(s) {
		return ""
	}

	return s[start : start+minLen]
}

func minWindowUsingArray(s string, t string) string {
	var tList, countList [128]int
	minLen := len(s)
	minSubstring := ""

	tLen := len(t)
	for i := range t {
		tList[t[i]]++
	}

	i, j := 0, 0
	count := 0
	for i < len(s) {

		for j < len(s) && count < tLen {
			if tList[s[j]] > 0 {
				countList[s[j]]++

				if countList[s[j]] <= tList[s[j]] {
					count++
				}
			}

			j++
		}

		if count == tLen {
			if j-i <= minLen {
				minSubstring = s[i:j]
				minLen = j - i
			}
		}

		if countList[s[i]] > 0 {
			if countList[s[i]] <= tList[s[i]] {
				count--
			}

			countList[s[i]] -= 1
		}

		i++
	}

	return minSubstring
}

// https://leetcode.com/problems/minimum-window-substring/description/
func minWindowSubstring() {
	s := "ADOBECODEBANC"
	t := "ABC"

	fmt.Println(minWindow(s, t))
	fmt.Println(minWindowUsingSingleMap(s, t))
	fmt.Println(minWindowUsingArray(s, t))
}
