package main

import "fmt"

func GetPermutations(array []int) [][]int {
	if len(array) == 0 {
		return [][]int{}
	}

	perms := helper(0, array, [][]int{})

	return perms
}

func helper(i int, array []int, perms [][]int) [][]int {
	if len(array)-1 == i {
		temp := make([]int, len(array))
		copy(temp, array)
		perms = append(perms, temp)

		return perms
	}

	for j := i; j < len(array); j++ {
		array[i], array[j] = array[j], array[i]

		perms = helper(i+1, array, perms)

		array[i], array[j] = array[j], array[i]
	}

	return perms
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/24_permutations
func permutation() {
	arr := []int{1, 2, 3}

	fmt.Println(GetPermutations(arr))
}
