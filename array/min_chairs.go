package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func findMinChairs(timestamps [][]int) int {
	queue := []int{}
	minChairs := 0
	count := 0

	sort.Slice(timestamps, func(i, j int) bool {
		return timestamps[i][1] < timestamps[j][1]
	})

	for i := range timestamps {
		timestamp := timestamps[i]

		for len(queue) > 0 {
			if queue[0] <= timestamp[0] {
				queue = queue[1:]
				count--
			} else {
				break
			}
		}

		count++
		queue = append(queue, timestamp[1])
		if count > minChairs {
			minChairs = count
		}
	}

	return minChairs
}

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
	list := *h
	temp := list[len(list)-1]
	*h = list[:len(list)-1]
	return temp
}

// just to check how much I remember priority queue; above solution is fine
func findMinChairs2(timestamps [][]int) int {
	queue := []int{}
	minChairs := 0
	count := 0
	endTime := &intHeap{}
	for i := range timestamps {
		timestamp := timestamps[i]

		*endTime = append(*endTime, timestamp[1])
	}

	heap.Init(endTime)

	for i := range timestamps {
		timestamp := timestamps[i]

		for len(queue) > 0 {
			if queue[0] <= timestamp[0] {
				queue = queue[1:]
				count--
			} else {
				break
			}
		}

		count++
		queue = append(queue, timestamp[1])
		if count > minChairs {
			minChairs = count
		}
	}

	return minChairs
}

// Find min no. of chairs needed for employees
// timestamps [0900, 1230] [1000, 1200] [1210, 1230] [1230 1400] [1400 1700]
// min = 1200; stack = 1200, 1230,
// if arr[0] < min; count = 2, 1
func minChairs() {
	timestamps := [][]int{{900, 1230}, {1030, 1200}, {1200, 1230}, {1230, 1400}, {1400, 1800}, {1900, 2100}}
	fmt.Println(findMinChairs(timestamps))
	fmt.Println(findMinChairs2(timestamps))
}
