package main

import (
	"fmt"
)

const mod = 1000000007

func pow(x, y int) int {
	if y == 0 {
		return 1
	}

	ans := pow(x, y/2)
	ans *= ans
	ans %= mod

	if y%2 != 0 {
		ans *= x
	}

	ans %= mod

	return ans

}

func countGoodNumbers(n int64) int {
	even := (n + 1) / 2
	odd := n / 2

	return pow(5, int(even)) * pow(4, int(odd)) % mod
}

// https://leetcode.com/problems/count-good-numbers/
func goodNumber() {
	fmt.Println(countGoodNumbers(4))
}
