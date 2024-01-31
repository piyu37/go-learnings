package main

import "fmt"

func find_levenshtien_distance(str1, str2 string) int {
	str1 = fmt.Sprint(" ", str1)
	str2 = fmt.Sprint(" ", str2)
	currArr := make([]int, len(str1))

	for i := range str1 {
		currArr[i] = i
	}

	prevBlockVal := 1
	for i := 1; i < len(str2); i++ {
		prevArr := make([]int, len(currArr))
		copy(prevArr, currArr)
		currArr[0] = i
		for j := 1; j < len(str1); j++ {
			if str1[j] == str2[i] {
				prevBlockVal = prevArr[j-1]
			} else {
				prevBlockVal = min(min(prevArr[j], prevArr[j-1]), currArr[j-1]) + 1
			}

			currArr[j] = prevBlockVal
		}
	}

	return prevBlockVal
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}

	return v2
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/16_levenshtein_distance
func levenshtienDistance() {
	fmt.Println(find_levenshtien_distance("abc", "yabd"))
}
