package main

import "fmt"

// move negative elements to left without changing order
func moveNegativeElements(arr []int) {
	n := len(arr)
	left := 0
	for i := 0; i < n; i++ {
		if arr[i] < 0 {
			if i != left {
				val := arr[i]
				k := i
				for k > left {
					arr[k] = arr[k-1]
					k--
				}

				arr[k] = val
			}

			left++
		}
	}
}

func moveNegativeToEnd() {
	ar := []int{1, 2, 3, -4, -5, -6, -7}
	fmt.Println("Original array:", ar)
	moveNegativeElements(ar)
	fmt.Println("Modified array:", ar)
}
