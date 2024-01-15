package main

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

// func recursiveQuickSort(arr []int, low, high int) {
// 	if low < high {
// 		p := partition(arr, low, high)

// 		recursiveQuickSort(arr, low, p-1)
// 		recursiveQuickSort(arr, p+1, high)
// 	}
// }

// func main() {
// 	arr := []int{4, 3, 5, 1, 2, 3, 1, 3}

// 	arrLen := len(arr)

// 	recursiveQuickSort(arr, 0, arrLen-1)

// 	fmt.Println(arr)

// }
