package main

import "fmt"

func Powerset(array []int) [][]int {
	result := make([][]int, 0)

	result = append(result, []int{})

	for i := range array {
		val := array[i]

		for j := range result {
			tempArr := result[j]

			tempArr = append(tempArr, val)
			result = append(result, tempArr)
		}
	}

	return result
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/25_powerset
func powersetMain() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(Powerset(arr))
}
