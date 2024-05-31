package main

import "fmt"

func getNext(n int) int {
	totalSum := 0
	for n > 0 {
		d := n % 10
		n = n / 10
		totalSum += d * d
	}
	return totalSum
}

func isHappy(n int) bool {
	slowRunner := n
	fastRunner := getNext(n)
	for fastRunner != 1 && slowRunner != fastRunner {
		slowRunner = getNext(slowRunner)
		fastRunner = getNext(getNext(fastRunner))
	}
	return fastRunner == 1
}

// https://leetcode.com/problems/happy-number/description/?envType=study-plan-v2&envId=top-interview-150
func isHappyMain() {
	n := 19
	fmt.Println(isHappy(n))
}
