package main

import (
	"container/heap"
	"fmt"
)

// lift; panel where we can enter numbers on which floor we want to go
// output: reduce lift movement; reduce no. of rounds

// 10 people: 2, 4, 6, 6, 8, 2, 3, 7, 8
// priority queue to stop the lift; top min queue, down max queue

// 1. use map to store floor no.

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	old := *h
	len := len(old)
	val := old[len-1]
	*h = old[:len-1]

	return val
}

func buildElevator(topQueue intHeap) {
	for len(topQueue) > 0 {
		fmt.Println(heap.Pop(&topQueue))
	}
}

func elevator() {
	var minPQ intHeap
	heap.Init(&minPQ)

	elevMap := make(map[int]bool)

	arr := []int{2, 4, 6, 6, 8, 2, 3, 7, 8}

	for i := range arr {
		if _, ok := elevMap[arr[i]]; !ok {
			heap.Push(&minPQ, arr[i])

			elevMap[arr[i]] = true
		}
	}

	buildElevator(minPQ)
}
