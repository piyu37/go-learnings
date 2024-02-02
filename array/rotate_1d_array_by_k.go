package main

import "fmt"

func rotateArrayByK() {
	var (
		N, K int = 6, 2
	)
	// fmt.Scan(&T)
	// fmt.Scan(&N, &K)
	arr := []int{1, 2, 3, 4, 5, 6}

	// for i := range arr {
	// 	fmt.Scan(&arr[i])
	// }

	for i := range arr {
		fmt.Printf("%d ", arr[(i+(N-K))%N])
	}
}
