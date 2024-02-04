package main

import "fmt"

func generatePrimeNumbersTill100() {
	fmt.Print(1, 2, 3)

	for i := 4; i <= 100; i++ {
		flag := false
		for j := 2; j < i/2; j++ {
			if i%j == 0 {
				flag = true
				break
			}
		}

		if !flag {
			fmt.Print("\t", i)
		}
	}
}
