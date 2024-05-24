package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := make(map[string]int)
	for _, w := range strs {
		h := hash(w)
		idx, ok := m[h]
		if ok {
			res[idx] = append(res[idx], w)
		} else {
			res = append(res, []string{w})
			m[h] = len(res) - 1
		}
	}

	return res
}

func hash(s string) string {
	res := make([]byte, 26)
	for _, c := range s {
		res[c-'a'] += 1
	}
	return string(res)
}

// https://leetcode.com/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-interview-150
func groupAnagramsMain() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
