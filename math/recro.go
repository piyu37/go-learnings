package main

import "fmt"

func recro() {
	T := 1

	count := 0

	for count < T {
		N := 5

		// goodCount := 0

		// for i := 1; i <= N; i++ {
		// 	goodCount += 32 * (i - 1)
		// }

		// goodCount = goodCount + (4 * N)

		goodCount := 16*N*(N-1) + (4 * N)

		fmt.Println(goodCount)

		count++
	}
}
