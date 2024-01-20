package main

import (
	"container/heap"
	"fmt"
)

var MAX_VALUE int = 100_001

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	leftBound := min(candidates-1, n-1)
	rightBound := max(n-candidates, leftBound+1)
	leftHeap := IntHeap(costs[:leftBound+1])
	rightHeap := IntHeap(costs[rightBound:])

	heap.Init(&leftHeap)
	heap.Init(&rightHeap)
	cost := 0

	for i := 0; i < k; i++ {
		leftCheapest := MAX_VALUE
		if len(leftHeap) > 0 {
			leftCheapest = leftHeap[0]
		}
		rightCheapest := MAX_VALUE
		if len(rightHeap) > 0 {
			rightCheapest = rightHeap[0]
		}

		if leftCheapest <= rightCheapest {
			cost += heap.Pop(&leftHeap).(int)
			if leftBound < rightBound-1 {
				leftBound++
				heap.Push(&leftHeap, costs[leftBound])
			}
		} else {
			cost += heap.Pop(&rightHeap).(int)
			if leftBound < rightBound-1 {
				rightBound--
				heap.Push(&rightHeap, costs[rightBound])
			}
		}
	}
	return int64(cost)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func min(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}

	return v1
}

// https://leetcode.com/problems/total-cost-to-hire-k-workers/
func totalCostHireKWorkers() {
	arr := []int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}
	// fmt.Println(totalCost(arr, 11, 2))
	fmt.Println(totalCost(arr, 11, 2))
}
