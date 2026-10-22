package main

import "fmt"

func findAnagrams(s string, p string) []int {
	result := []int{}

	if len(p) > len(s) {
		return result
	}

	pMap := [26]int{}
	sMap := [26]int{}

	for i := range p {
		pMap[p[i]-'a']++
		sMap[s[i]-'a']++
	}

	pLen := len(p)
	i := 0
	for ; i < len(s)-len(p); i++ {
		if sMap == pMap {
			result = append(result, i)
		}

		sMap[s[i]-'a']--
		sMap[s[i+pLen]-'a']++
	}

	if sMap == pMap {
		result = append(result, i)
	}

	return result
}

// https://leetcode.com/problems/find-all-anagrams-in-a-string/description/
func anagramsInString() {
	s := "cbaebabacd"
	p := "abc"

	fmt.Println(findAnagrams(s, p))
}
