package main

import "fmt"

func validWordAbbreviation(word string, abbr string) bool {
	num := 0
	wordIdx := 0
	for i := 0; i < len(abbr); i++ {
		ch := abbr[i]
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')

			if num > 20 || num == 0 {
				return false
			}
		} else {
			wordIdx += num
			num = 0
			if wordIdx >= len(word) || abbr[i] != word[wordIdx] {
				return false
			}

			wordIdx++
		}
	}

	wordIdx += num

	return wordIdx == len(word)
}

// https://leetcode.com/problems/valid-word-abbreviation/
func validWordAbbreviationMain() {
	word := "internationalization"
	abbr := "i12iz4n"

	fmt.Println(validWordAbbreviation(word, abbr))
}
