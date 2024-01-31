package main

import "fmt"

func longestCommonSubsequence(str1, str2 string) []string {
	arr := make([][]int, len(str2)+1)
	arr[0] = make([]int, len(str1)+1)
	for i := 1; i < len(arr); i++ {
		arr[i] = make([]int, len(str1)+1)
		for j := 1; j < len(arr[i]); j++ {
			if str2[i-1] == str1[j-1] {
				arr[i][j] = 1 + arr[i-1][j-1]
			} else {
				arr[i][j] = max(arr[i-1][j], arr[i][j-1])
			}
		}
	}

	longestSubsequence := make([]string, arr[len(str2)][len(str1)])
	count := arr[len(str2)][len(str1)]

	row := len(str2)
	col := len(str1)
	for count > 0 {
		if str2[row-1] == str1[col-1] {
			count--
			longestSubsequence[count] = string(str2[row-1])
			row--
			col--
		} else {
			if arr[row-1][col] > arr[row][col-1] {
				row--
			} else {
				col--
			}
		}
	}

	return longestSubsequence
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/11_longest_common_subsequence
func longestCommonSubsequenceMain() {
	str1 := "CCCDDEGDHAGKGLWAJWKJAWGKGWJAKLGGWAFWLFFWAGJWKAGTUV"
	str2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println(longestCommonSubsequence(str1, str2))
}
