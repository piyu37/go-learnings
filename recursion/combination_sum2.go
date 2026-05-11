package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)

	var findCombination func(candidates []int, target int, idx int, combination []int)

	findCombination = func(candidates []int, target int, idx int, combination []int) {
		if target == 0 {
			temp := make([]int, len(combination))
			copy(temp, combination)
			result = append(result, temp)

			return
		}

		if target < 0 {
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}

			findCombination(candidates, target-candidates[i], i+1, append(combination, candidates[i]))
		}
	}

	findCombination(candidates, target, 0, []int{})

	return result
}

// https://leetcode.com/problems/combination-sum-ii/description/
func combinationSum2Main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	fmt.Println(combinationSum2(candidates, target))
}
