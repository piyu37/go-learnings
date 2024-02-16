package main

import "fmt"

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	nums := make([]int, n)

	for i := range nums {
		nums[i] = i + 1
	}

	currCombination := make([]int, k)

	createCombinations(k, 0, 0, nums, currCombination, &result)

	return result
}

func createCombinations(k, idx, numIdx int, nums, currCombination []int,
	result *[][]int) {
	if idx == k {
		temp := make([]int, k)
		copy(temp, currCombination)
		*result = append(*result, temp)
		return
	}

	for i := numIdx; i < len(nums); i++ {
		currCombination[idx] = nums[i]
		createCombinations(k, idx+1, i+1, nums, currCombination, result)
	}
}

func combineOptimized(n int, k int) [][]int {
	result := make([][]int, 0)
	currCombination := make([]int, k)
	createCombinations2(n, k, 0, 1, currCombination, &result)

	return result
}

func createCombinations2(n, k, currIdx, currNum int, currCombination []int, result *[][]int) {
	if currIdx == k {
		temp := make([]int, k)
		copy(temp, currCombination)
		*result = append(*result, temp)
		return
	}

	need := k - currIdx
	remain := n - currNum + 1
	available := remain - need

	for i := currNum; i <= currNum+available; i++ {
		currCombination[currIdx] = i
		createCombinations2(n, k, currIdx+1, i+1, currCombination, result)
	}
}

// https://leetcode.com/problems/combinations/description/?envType=study-plan-v2&envId=top-interview-150
func combinations() {
	n := 4
	k := 3

	fmt.Println(combine(n, k))

	fmt.Println(combineOptimized(n, k))
}
