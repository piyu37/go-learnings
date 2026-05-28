package main

import "fmt"

// ababab
// bab
// babad

func findLongestpalindromeUsingBottomUpDP(str string) string {
	result := make([][]int, len(str))
	startIndex := 0
	endIndex := 0

	for i := 0; i < len(str); i++ {
		result[i] = make([]int, len(str))

		result[i][i] = 1

		if i+1 < len(str) && str[i] == str[i+1] {
			result[i][i+1] = 1
			startIndex = i
			endIndex = i + 1
		}
	}

	diff := 2
	for diff < len(str) {
		i := 0
		j := diff
		for j < len(str) {
			if result[i+1][j-1] == 1 && str[i] == str[j] {
				result[i][j] = 1
				startIndex = i
				endIndex = j
			}

			i++
			j++
		}

		diff++
	}

	return str[startIndex : endIndex+1]
}

func findLongestpalindromeExpandFromCenter(str string) string {
	ans := ""

	for i := range str {
		odd := expand(i, i, str)
		if len(odd) > len(ans) {
			ans = odd
		}

		even := expand(i, i+1, str)
		if len(even) > len(ans) {
			ans = even
		}
	}

	return ans
}

func expand(left, right int, str string) string {
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}

	return str[left+1 : right]
}

func longestPalindromeManacherAlgoOptimal(s string) string {
	sPrime := "#"
	for _, c := range s {
		sPrime = sPrime + string(c) + "#"
	}

	n := len(sPrime)
	palindromeRadii := make([]int, n)
	center := 0
	radius := 0

	for i := 0; i < n; i++ {
		mirror := 2*center - i

		if i < radius {
			palindromeRadii[i] = min(radius-i, palindromeRadii[mirror])
		}

		for i+1+palindromeRadii[i] < n && i-1-palindromeRadii[i] >= 0 &&
			sPrime[i+1+palindromeRadii[i]] == sPrime[i-1-palindromeRadii[i]] {
			palindromeRadii[i]++
		}

		if i+palindromeRadii[i] > radius {
			center = i
			radius = i + palindromeRadii[i]
		}
	}

	maxLength := 0
	centerIndex := 0
	for i := 0; i < n; i++ {
		if palindromeRadii[i] > maxLength {
			maxLength = palindromeRadii[i]
			centerIndex = i
		}
	}

	startIndex := (centerIndex - maxLength) / 2
	longestPalindrome := s[startIndex : startIndex+maxLength]

	return longestPalindrome
}

// https://leetcode.com/problems/longest-palindromic-substring/
func longestPalindromString() {
	str := "abbabab"

	fmt.Println(findLongestpalindromeUsingBottomUpDP(str))

	str = "abbabab"

	fmt.Println(findLongestpalindromeExpandFromCenter(str))

	str = "abbabab"

	fmt.Println(longestPalindromeManacherAlgoOptimal(str))
}
