package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	result := ""
	endIdx := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			continue
		} else {
			result += s[i+1 : endIdx]
			if endIdx-i-1 > 0 {
				result += " "
			}

			endIdx = i
		}
	}

	result += s[0:endIdx]

	if result[len(result)-1] == ' ' {
		return result[:len(result)-1]
	}

	return result
}

func reverseWordsUsingInbuilt(s string) string {
	words := strings.Fields(s) // splits on whitespace, drops empty strings, so leading/trailing/multiple spaces are handled for free
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// https://leetcode.com/problems/reverse-words-in-a-string/description/
func reverseWordsMain() {
	str := "   the sky is blue    "
	fmt.Println(reverseWords(str))
	fmt.Println(reverseWordsUsingInbuilt(str))
}
