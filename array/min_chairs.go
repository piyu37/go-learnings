package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// Chronological Ordering Approach
func findMinChairs(timestamps [][]int) int {
	minChairs := 0

	startIntervals := make([]int, len(timestamps))
	endIntervals := make([]int, len(timestamps))

	for i := range timestamps {
		startIntervals[i] = timestamps[i][0]
		endIntervals[i] = timestamps[i][1]
	}

	sort.Ints(startIntervals)
	sort.Ints(endIntervals)

	i, j := 0, 0
	for i < len(startIntervals) {
		if startIntervals[i] < endIntervals[j] {
			minChairs++
		} else {
			j++
		}

		i++
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
	sort.Slice(timestamps, func(i, j int) bool {
		return timestamps[i][0] < timestamps[j][0]
	})

	freeChairs := &intHeap{}
	heap.Push(freeChairs, timestamps[0][1])

	for i := 1; i < len(timestamps); i++ {
		timestamp := timestamps[i]

		if (*freeChairs)[0] <= timestamp[0] {
			heap.Pop(freeChairs)
		}

		heap.Push(freeChairs, timestamp[1])
	}

	return freeChairs.Len()
}

// Find min no. of chairs needed for employees
// timestamps [0900, 1230] [1000, 1200] [1210, 1230] [1230 1400] [1400 1700]

// Similar kind of problem in diff words:
// the maximum number of intervals that overlap at any single point.
// https://leetcode.com/problems/meeting-rooms-ii/description/
// https://www.hack2hire.com/question-bank/companies/amazon/coding-questions/68eda9b25abdefad4f6693bd/practice
func minChairs() {
	timestamps := [][]int{{900, 1230}, {1030, 1200}, {1200, 1230}, {1230, 1400}, {1400, 1800}, {1900, 2100}}
	fmt.Println(findMinChairs(timestamps))
	fmt.Println(findMinChairs2(timestamps))

	timestamps = [][]int{{600, 1700}, {800, 900}, {1100, 1200}, {600, 900}}

	fmt.Println(findMinChairs(timestamps))
	fmt.Println(findMinChairs2(timestamps))

	// 1,4 2,3 3,4 5,6
	timestamps = [][]int{{1, 4}, {2, 3}, {3, 4}, {5, 10}, {6, 7}, {7, 10}, {9, 10}}

	fmt.Println(findMinChairs(timestamps))
}
