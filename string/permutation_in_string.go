package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	s1Map := [26]int{}
	s2Map := [26]int{}

	for i := range s1 {
		s1Map[s1[i]-'a']++
		s2Map[s2[i]-'a']++
	}

	if s1Map == s2Map {
		return true
	}

	s1Len := len(s1)
	for i := s1Len; i < len(s2); i++ {
		s2Map[s2[i-s1Len]-'a']--
		s2Map[s2[i]-'a']++

		if s1Map == s2Map {
			return true
		}
	}

	return false
}

// https://leetcode.com/problems/permutation-in-string/description/
func permutationInString() {
	s1 := "ab"
	s2 := "eidbaooo"

	fmt.Println(checkInclusion(s1, s2))
}
