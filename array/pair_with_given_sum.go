package main

import "fmt"

// https://www.geeksforgeeks.org/count-pairs-with-given-sum/
// arr := []int{1, 2, 10, 40, 3, 9, 4}
// sum=12
func pairWithGivenSum() {
	arr := []int{1, 2, 10, 40, 3, 9, 4}
	sum := 12
	result := make([][]int, 0)

	intMap := make(map[int]int)

	for i := range arr {
		val := arr[i]

		if _, ok := intMap[sum-val]; ok {
			result = append(result, []int{val, sum - val})
		}

		intMap[val] = val
	}

	fmt.Println(result)
}
