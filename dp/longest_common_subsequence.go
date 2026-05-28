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

func longestCommonSubsequence2(str1, str2 string) int {
	memoization := make(map[string]int)

	return lcs(len(str1)-1, len(str2)-1, str1, str2, memoization)
}

func lcs(i, j int, str1, str2 string, memoization map[string]int) int {
	if i < 0 || j < 0 {
		return 0
	}

	if v, ok := memoization[fmt.Sprintf("%d:%d", i, j)]; ok {
		return v
	}

	var result int
	if str1[i] == str2[j] {
		result = 1 + lcs(i-1, j-1, str1, str2, memoization)
	} else {
		result = max(lcs(i-1, j, str1, str2, memoization), lcs(i, j-1, str1, str2, memoization))
	}

	memoization[fmt.Sprintf("%d:%d", i, j)] = result

	return result
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/11_longest_common_subsequence
func longestCommonSubsequenceMain() {
	str1 := "CCCDDEGDHAGKGLWAJWKJAWGKGWJAKLGGWAFWLFFWAGJWKAGTUV"
	str2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println(longestCommonSubsequence(str1, str2))

	fmt.Println(longestCommonSubsequence2(str1, str2))
}
