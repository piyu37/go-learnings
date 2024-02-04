package main

import "fmt"

// Given a BST, Add the value greater than the current value in the given BST
// https://www.geeksforgeeks.org/add-greater-values-every-node-given-bst/

// program to return subarray which are giving max sum from array
// https://www.geeksforgeeks.org/print-the-maximum-subarray-sum/
// 2, 2, 0, -1, 3, -4 => [0, 4]
func MaxSum(arr []int) []int {
	start := 0
	end := len(arr) - 1
	maxSum := arr[0]
	for i := 0; i < len(arr); i++ {
		sum := arr[i]
		for j := i + 1; j < len(arr); j++ {
			sum += arr[j]
			if sum > maxSum {
				maxSum = sum
				start = i
				end = j
			}
		}
	}

	return []int{start, end}
}

func maxSumMain() {
	arr := []int{2, 2, 0, -1, 3, -4}
	fmt.Println(MaxSum(arr))
}
