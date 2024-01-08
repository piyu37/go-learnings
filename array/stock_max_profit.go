package main

import "fmt"

func maxProfit(prices []int) int {
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}

	return max
}

func stockMaxProfit() {
	fmt.Println(maxProfit([]int{1, 2, 4, 3, 5, 6}))
}
