package main

import "fmt"

// ababab
// bab
// babad

func findLongestpalindrome(str string) string {
	result := make([][]int, len(str))

	for i := 0; i < len(str); i++ {
		result[i] = make([]int, len(str))

		result[i][i] = 1

		if i+1 < len(str) && str[i] == str[i+1] {
			result[i][i+1] = 1
		}
	}

	diff := 2
	for diff < len(str) {
		i := 0
		j := diff
		for j < len(str) {
			if result[i+1][j-1] == 1 && str[i] == str[j] {
				result[i][j] = 1
			}

			i++
			j++
		}

		diff++
	}

	startIndex := 0
	endIndex := 0
	maxCount := 1
	for i := 0; i < len(str); i++ {
		tmpStartIndex := i
		tmpEndIndex := i
		for j := i + 1; j < len(str); j++ {
			if result[i][j] == 1 {
				tmpEndIndex = j
			}
		}

		if tmpEndIndex-tmpStartIndex > maxCount {
			maxCount = tmpEndIndex - tmpStartIndex
			startIndex = tmpStartIndex
			endIndex = tmpEndIndex
		}
	}

	return str[startIndex : endIndex+1]
}

func longestPalindromString() {
	str := "abbabab"

	fmt.Println(findLongestpalindrome(str))
}
