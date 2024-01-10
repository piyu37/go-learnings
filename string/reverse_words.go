package main

import "fmt"

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

func reverseWordsMain() {
	str := "   the sky is blue    "
	fmt.Println(reverseWords(str))
}
