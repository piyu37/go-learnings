package main

import "fmt"

func knapsack(items [][]int, cap int) (int, []int) {
	result := make([][]int, len(items)+1)

	result[0] = make([]int, cap+1)

	for i := 1; i <= len(items); i++ {
		item := items[i-1]
		result[i] = make([]int, cap+1)
		for c := 1; c <= cap; c++ {
			if c < item[1] {
				result[i][c] = result[i-1][c]
			} else {
				maxProfit := max(result[i-1][c], (result[i-1][c-item[1]] + item[0]))

				result[i][c] = maxProfit
			}
		}
	}

	maxProfit := result[len(items)][cap]
	profit := maxProfit
	i := len(items)
	j := cap
	arr := make([]int, i)

	for profit != 0 {
		if result[i][j] > result[i-1][j] {
			profit -= items[i-1][0]
			j -= items[i-1][1]
			arr[i-1] = 1
		} else {
			i--
		}
	}

	finalItems := make([]int, 0)

	for i := range arr {
		if arr[i] == 1 {
			finalItems = append(finalItems, i)
		}
	}

	return maxProfit, finalItems
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/14_knapsack_problem
func knapsackMain() {
	items := [][]int{{1, 2}, {2, 3}, {5, 4}, {7, 5}}
	cap := 8

	fmt.Println(knapsack(items, cap))
}
