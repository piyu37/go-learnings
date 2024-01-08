package main

import "fmt"

// move negative elements to left without changing order
func moveNegativeElements(arr []int) {
	n := len(arr)
	left := 0

	count := 0
	for i := 0; i < n; i++ {
		count++
		if arr[i] < 0 {
			if i != left {
				val := arr[i]
				k := i
				for k > left {
					count++
					arr[k] = arr[k-1]
					k--
				}

				arr[k] = val
			}

			left++
		}
	}

	fmt.Println(count)
}
