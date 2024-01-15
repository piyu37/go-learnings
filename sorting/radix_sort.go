package main

import "fmt"

func RadixSort(array []int) []int {
	if len(array) == 0 {
		return array
	}

	maxNum := max(array)

	length := 0

	for maxNum > 0 {
		maxNum /= 10
		countingSort(array, length)
		length++
	}

	return array
}

func countingSort(arr []int, digit int) {
	counts := make([]int, 10)

	sortedArr := make([]int, len(arr))

	col := pow(digit)

	for _, v := range arr {
		valueDigit := (v / col) % 10

		counts[valueDigit] += 1
	}

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		valueDigit := (arr[i] / col) % 10

		counts[valueDigit] -= 1

		sortedArr[counts[valueDigit]] = arr[i]
	}

	_ = copy(arr, sortedArr)
}

func max(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func pow(num int) int {
	pow := 1

	for num > 0 {
		pow *= 10
		num--
	}

	return pow
}

func radixSortMain() {
	arr := []int{8762, 654, 3008, 345, 87, 65, 234, 12, 2}
	fmt.Println(RadixSort(arr))
}
