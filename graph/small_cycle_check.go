package main

import "fmt"

func singleCycleCheck(arr []int) bool {
	index := arr[0]
	indexCount := 1

	for indexCount < len(arr) {
		if indexCount > 0 && index == 0 {
			return false
		}

		indexCount++

		index = index + arr[index]
		if index < 0 {
			index = len(arr) + index
		} else if index > len(arr) {
			index = index % len(arr)
		}
	}

	return index == 0
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/18_single_cycle_check
func smallCycleCheck() {
	arr := []int{2, 3, 1, -4, -4, 2}

	fmt.Println(singleCycleCheck(arr))

}
