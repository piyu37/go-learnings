package main

import (
	"fmt"
	"math"
)

func sameBST(arr1, arr2 []int) bool {
	return isSame(len(arr1), len(arr2), 0, 0,
		math.MinInt, math.MaxInt, math.MinInt, math.MaxInt, arr1, arr2)
}

func isSame(count1, count2, firstValue1Idx, firstValue2Idx,
	arr1MinLimit, arr1MaxLimit, arr2MinLimit, arr2MaxLimit int,
	arr1, arr2 []int) bool {
	if count1 == 0 && count2 == 0 {
		return true
	}

	if count1 != count2 || arr1[firstValue1Idx] != arr2[firstValue2Idx] {
		return false
	}

	if count1 == 1 && count2 == 1 {
		return true
	}

	leftCount1 := 0
	leftFirstValueSaved1 := false
	var leftFirstValue1Idx int
	leftCount2 := 0
	leftFirstValueSaved2 := false
	var leftFirstValue2Idx int
	rightCount1 := 0
	rightFirstValueSaved1 := false
	var rightFirstValue1Idx int
	rightCount2 := 0
	rightFirstValueSaved2 := false
	var rightFirstValue2Idx int

	for i := firstValue1Idx + 1; i < len(arr1); i++ {
		if arr1[i] < arr1MinLimit || arr1[i] > arr1MaxLimit {
			continue
		}

		if arr1[firstValue1Idx] > arr1[i] {
			if !leftFirstValueSaved1 {
				leftFirstValue1Idx = i
				leftFirstValueSaved1 = true
			}
			leftCount1++
		} else {
			if !rightFirstValueSaved1 {
				rightFirstValue1Idx = i
				rightFirstValueSaved1 = true
			}
			rightCount1++
		}
	}

	for i := firstValue2Idx + 1; i < len(arr2); i++ {
		if arr2[i] < arr2MinLimit || arr2[i] > arr2MaxLimit {
			continue
		}

		if arr2[firstValue2Idx] > arr2[i] {
			if !leftFirstValueSaved2 {
				leftFirstValue2Idx = i
				leftFirstValueSaved2 = true
			}
			leftCount2++
		} else {
			if !rightFirstValueSaved2 {
				rightFirstValue2Idx = i
				rightFirstValueSaved2 = true
			}
			rightCount2++
		}
	}

	isSameBST := isSame(leftCount1, leftCount2, leftFirstValue1Idx, leftFirstValue2Idx,
		arr1MinLimit, arr1[firstValue1Idx], arr2MinLimit, arr2[firstValue2Idx], arr1, arr2)
	if !isSameBST {
		return false
	}

	return isSame(rightCount1, rightCount2, rightFirstValue1Idx, rightFirstValue2Idx,
		arr1[firstValue1Idx], arr1MaxLimit, arr2[firstValue2Idx], arr2MaxLimit, arr1, arr2)
}

func sameBST2(arr1, arr2 []int) bool {
	return areSameBSTs(arr1, arr2, 0, 0, math.MinInt32, math.MaxInt32)
}

func areSameBSTs(arrayOne, arrayTwo []int, rootIdxOne, rootIdxTwo int, minVal, maxVal int) bool {
	if rootIdxOne == -1 || rootIdxTwo == -1 {
		return rootIdxOne == rootIdxTwo
	}

	if arrayOne[rootIdxOne] != arrayTwo[rootIdxTwo] {
		return false
	}

	leftRootIdxOne := getIdxOfFirstSmaller(arrayOne, rootIdxOne, minVal)
	leftRootIdxTwo := getIdxOfFirstSmaller(arrayTwo, rootIdxTwo, minVal)
	rightRootIdxOne := getIdxOfFirstBiggerOrEqual(arrayOne, rootIdxOne, maxVal)
	rightRootIdxTwo := getIdxOfFirstBiggerOrEqual(arrayTwo, rootIdxTwo, maxVal)

	currentValue := arrayOne[rootIdxOne]
	leftAreSame := areSameBSTs(arrayOne, arrayTwo, leftRootIdxOne, leftRootIdxTwo, minVal, currentValue)
	rightAreSame := areSameBSTs(arrayOne, arrayTwo, rightRootIdxOne, rightRootIdxTwo, currentValue, maxVal)
	return leftAreSame && rightAreSame
}

func getIdxOfFirstSmaller(array []int, startingIdx, minVal int) int {
	// Find the index of the first smaller value after the startingIdx.
	// Make sure that this value is greater than or equal to the minVal,
	// which is the value of the previous parent node in the BST. If it
	// isn't, then that value is located in the left subtree of the
	// previous parent node.
	for i := startingIdx + 1; i < len(array); i++ {
		if array[i] < array[startingIdx] && array[i] >= minVal {
			return i
		}
	}
	return -1
}

func getIdxOfFirstBiggerOrEqual(array []int, startingIdx, maxVal int) int {
	// Find the index of the first bigger/equal value after the startingIdx.
	// Make sure that this value is smaller than maxVal, which is the value
	// of the previous parent node in the BST. If it isn't, then that value
	// is located in the right subtree of the previous parent node.
	for i := startingIdx + 1; i < len(array); i++ {
		if array[i] >= array[startingIdx] && array[i] < maxVal {
			return i
		}
	}
	return -1
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/06_same_bsts
func sameBSTMain() {
	arr1 := []int{10, 15, 8, 11, 12, 94, 81, 5, 2}
	arr2 := []int{10, 8, 5, 15, 2, 12, 11, 94, 81}

	fmt.Println(sameBST(arr1, arr2))
	fmt.Println(sameBST2(arr1, arr2))
}
