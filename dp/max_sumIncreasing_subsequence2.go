package main

func findMaxSumIncreasingSubsequence2(arr []int) ([]int, int) {
	msis := make([]int, len(arr))
	maxSum := 0
	maxSumArr := make([]int, 0)

	copy(msis, arr)

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && msis[i] < arr[i]+msis[j] {
				msis[i] = arr[i] + msis[j]
				if maxSum < arr[i]+msis[j] {
					maxSum = arr[i] + msis[j]
				}
			}
		}
	}

	temp := maxSum
	for i := len(msis) - 1; i >= 0; i-- {
		if msis[i] == temp {
			maxSumArr = append(maxSumArr, arr[i])
			temp -= arr[i]
		}
	}

	return maxSumArr, maxSum
}
