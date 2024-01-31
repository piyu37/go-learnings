package main

import "fmt"

func Max(N int, K int, arr []int) int {
	leftOdd, leftEven := 0, 0
	rightOdd, rightEven := 0, 0
	left := 0
	right := N - 1
	partitionCount := 0
	indexPartitionArr := make([]int, 101)
	for left < right {
		if arr[left]%2 == 0 {
			rightEven++
		} else {
			rightOdd++
		}

		if arr[right]%2 == 0 {
			rightEven++
		} else {
			rightOdd++
		}

		left++
		right--
	}

	kTotal := 0
	for i := 0; i < N-2; i++ {
		if arr[i]%2 == 0 {
			leftEven++
			rightEven--
		} else {
			leftOdd++
			rightOdd--
		}

		if i%2 != 0 {
			if leftEven == leftOdd && rightEven == rightOdd {
				kDiff := abs(arr[i] - arr[i+1])
				kTotal += kDiff
				indexPartitionArr[kDiff] += 1
				partitionCount++
			}
		}
	}

	i := 100
	for i >= 0 {
		for kTotal > K && indexPartitionArr[i] > 0 {
			kTotal -= i
			partitionCount--
			indexPartitionArr[i]--
		}

		i--
	}

	return partitionCount
}

func abs(v1 int) int {
	if v1 >= 0 {
		return v1
	}

	return -v1
}

// https://www.hackerearth.com/practice/algorithms/dynamic-programming/introduction-to-dynamic-programming-1/practice-problems/algorithm/max-seperations-4ed0b552/
func maxSeparation() {
	N := 100
	K := 100
	ar := []int{60, 83, 82, 16, 17, 7, 89, 6, 83, 100, 85, 41, 72, 44, 23, 28, 64, 84, 3, 23, 33, 52, 93,
		30, 81, 38, 67, 25, 26, 97, 94, 78, 41, 74, 74, 17, 53, 51, 54, 17, 20, 81, 95, 76, 42, 16, 16, 56,
		74, 69, 30, 9, 82, 91, 32, 13, 47, 45, 97, 40, 56, 57, 27, 28, 84, 98, 91, 5, 61, 20, 3, 43, 42, 26,
		83, 40, 34, 100, 5, 63, 62, 61, 72, 5, 32, 58, 93, 79, 7, 18, 50, 43, 17, 24, 77, 73, 87, 74, 98, 2}
	var out_ int = Max(N, K, ar)
	fmt.Println(out_)
}
