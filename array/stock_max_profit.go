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

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150
func stockMaxProfit() {
	fmt.Println(maxProfit([]int{1, 2, 4, 3, 5, 6}))
}
