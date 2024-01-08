package main

// https://www.hackerearth.com/practice/algorithms/dynamic-programming/introduction-to-dynamic-programming-1/practice-problems/algorithm/max-seperations-4ed0b552/
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
