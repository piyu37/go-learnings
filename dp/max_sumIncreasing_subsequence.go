package main

import "fmt"

func findMaxSumIncreasingSubsequence(arr []int) ([]int, int) {
	return findMaxSum(arr, 0)
}

func findMaxSum(arr []int, idx int) ([]int, int) {
	if idx >= len(arr) {
		return []int{}, 0
	}

	maxValuesArr := []int{arr[idx]}
	maxValue := arr[idx]

	maxValuesArrTemp, maxValueTemp := findMaxSum(arr, idx+1)

	for j := 0; j < len(maxValuesArrTemp); j++ {
		if arr[idx] < maxValuesArrTemp[j] {
			maxValuesArr = append(maxValuesArr, maxValuesArrTemp[j])
			maxValue += maxValuesArrTemp[j]
		}
	}

	if maxValue < maxValueTemp {
		maxValuesArr = maxValuesArrTemp
		maxValue = maxValueTemp
	}

	return maxValuesArr, maxValue
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/10_max_sumIncreasing_subsequence
func maxSumIncreasingSubsequence() {
	arr := []int{-5, 10, 70, 20, 30, 50, 11, 30, 90}
	fmt.Println(findMaxSumIncreasingSubsequence(arr))

	fmt.Println(findMaxSumIncreasingSubsequence2(arr))
}
