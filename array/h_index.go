package main

import "fmt"

func hIndex(citations []int) int {
	sortedArr := make([]int, len(citations)+1)

	for i := range citations {
		if citations[i] > len(citations) {
			sortedArr[len(citations)] += 1
		} else {
			sortedArr[citations[i]]++
		}
	}

	count := 0
	i := len(citations)
	for i >= 0 {
		count += sortedArr[i]

		if count >= i {
			return i
		}

		i--
	}

	return 0
}

// https://leetcode.com/problems/h-index/description/?envType=study-plan-v2&envId=top-interview-150
func hIndexMain() {
	arr := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(arr))
}
