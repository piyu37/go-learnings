package main

import "fmt"

// import "fmt"

// func partition(arr []int, l, h int) int {
// 	i := l - 1
// 	pivot := arr[h]

// 	for j := l; j < h; j++ {
// 		if pivot >= arr[j] {
// 			i++

// 			arr[j], arr[i] = arr[i], arr[j]
// 		}
// 	}

// 	i++
// 	arr[i], arr[h] = arr[h], arr[i]

// 	return i
// }

func quickSort(arr []int, low, high int) []int {
	top := -1
	stack := make([]int, high+1)

	top++
	stack[top] = low
	top++
	stack[top] = high

	for top >= 0 {
		high = stack[top]
		top--
		low = stack[top]
		top--

		p := partition(arr, low, high)

		if low < p-1 {
			top++
			stack[top] = low
			top++
			stack[top] = p - 1
		}

		if high > p+1 {
			top++
			stack[top] = p + 1
			top++
			stack[top] = high
		}
	}

	return arr
}

func iterQuickSort() {
	arr := []int{4, 3, 5, 1, 2, 3, 1, 3}

	arrLen := len(arr)

	arr = quickSort(arr, 0, arrLen-1)

	fmt.Println(arr)

}
