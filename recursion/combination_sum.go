package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	findCombinations(0, candidates, target, []int{}, &result)

	return result
}

func findCombinations(idx int, candidates []int, target int, currComb []int, result *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		temp := make([]int, len(currComb))
		copy(temp, currComb)
		*result = append(*result, temp)
		return
	}

	for i := idx; i < len(candidates); i++ {
		currComb = append(currComb, candidates[i])
		findCombinations(i, candidates, target-candidates[i], currComb, result)
		currComb = currComb[:len(currComb)-1]
	}
}

// https://leetcode.com/problems/combination-sum/description/?envType=study-plan-v2&envId=top-interview-150
func combinationSumMain() {
	candidates := []int{2, 3, 5}
	target := 8

	fmt.Println(combinationSum(candidates, target))
}
