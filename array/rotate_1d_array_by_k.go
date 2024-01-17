package main

import "fmt"

func rotateArrayByK() {
	var (
		T, N, K int
	)
	fmt.Scan(&T)
	fmt.Scan(&N, &K)
	arr := make([]int, N)

	for i := range arr {
		fmt.Scan(&arr[i])
	}

	for i := range arr {
		fmt.Printf("%d ", arr[(i+(N-K))%N])
	}
}
