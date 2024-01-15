package main

import "fmt"

// func partition(array []int, low, high int) int {
// 	s := rand.NewSource(time.Now().Unix())
// 	r := rand.New(s)
// 	p := low + r.Intn(high-low+1)

// 	array[high], array[p] = array[p], array[high]

// 	i := low - 1

// 	p = array[high]

// 	for j := low; j < high; j++ {
// 		if array[j] <= p {
// 			i++

// 			array[j], array[i] = array[i], array[j]
// 		}
// 	}

// 	i++

// 	array[i], array[high] = p, array[i]

// 	return i
// }

func randomizedIterQuickSort(array []int) []int {
	low := 0

	high := len(array) - 1

	if low >= high {
		return array
	}

	top := -1
	stack := make([]int, high+1)
	top++
	stack[top] = low
	top++
	stack[top] = high

	for top > 0 {
		high = stack[top]
		top--
		low = stack[top]
		top--

		p := partition(array, low, high)

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

	return array
}

func randomizedIterQuickSortMain() {
	fmt.Println(randomizedIterQuickSort([]int{8}))
}
