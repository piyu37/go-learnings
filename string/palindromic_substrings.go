package main

import "fmt"

func palindromicSubstrings(str string) []int {
	result := make([]int, len(str))
	result[0] = 1

	palArr := make([][]int, len(str))

	for i := 0; i < len(str); i++ {
		palArr[i] = make([]int, len(str))

		palArr[i][i] = 1

		if i+1 < len(str) && str[i] == str[i+1] {
			palArr[i][i+1] = 1
		}
	}

	substringSize := 2
	i, j := 0, substringSize
	for i != 0 || j < len(str) {
		for j < len(str) {

			if str[i] == str[j] && palArr[i+1][j-1] == 1 {
				palArr[i][j] = 1
			}

			i++
			j++
		}

		i = 0
		substringSize++
		j = substringSize
	}

	for i := 0; i < len(str); i++ {
		palindromeLength := 1
		for j := i + 1; j < len(str); j++ {
			if palArr[i][j] == 1 {
				palindromeLength = j - i + 1
				if result[j] < palindromeLength {
					result[j] = palindromeLength
				}
			} else if result[j] < palindromeLength {
				result[j] = palindromeLength
			}
		}
	}

	return result
}

func palindromicSubstringsMain() {
	str := "ababababababababababababab"

	fmt.Println(palindromicSubstrings(str))
}
