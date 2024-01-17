package main

import "fmt"

// Rearrange function rearranges the sorted array in maximum minimum form without using extra space
// e.g. 1 2 3 4 5 => 5 1 4 2 3
func rearrange(arr []int, n int) {
	// Initialize index of first minimum and first maximum element
	maxIdx := n - 1
	minIdx := 0

	// Store maximum element of array
	maxElem := arr[n-1] + 1

	// Traverse array elements
	for i := 0; i < n; i++ {
		// At even index: we have to put the maximum element
		if i%2 == 0 {
			arr[i] += (arr[maxIdx] % maxElem) * maxElem
			maxIdx--
		} else {
			// At odd index: we have to put the minimum element
			arr[i] += (arr[minIdx] % maxElem) * maxElem
			minIdx++
		}
	}

	// Array elements back to their original form
	for i := 0; i < n; i++ {
		arr[i] = arr[i] / maxElem
	}
}

func rearrangeMaxMin() {
	arr := []int{2, 3, 5, 7, 8, 9}
	n := len(arr)

	fmt.Println("Original Array")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

	rearrange(arr, n)

	fmt.Println("\nModified Array")
	for i := 0; i < n; i++ {
		fmt.Print(int(arr[i]), " ")
	}
}
