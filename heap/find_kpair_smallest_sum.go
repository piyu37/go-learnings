package main

import (
	"container/heap"
	"fmt"
)

type IntHeap2 [][]int

func (h IntHeap2) Len() int           { return len(h) }
func (h IntHeap2) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntHeap2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap2) Push(x any) {
	v := x.([]int)
	*h = append(*h, v)
}

func (h *IntHeap2) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var pq IntHeap2
	heap.Init(&pq)
	result := [][]int{}

	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}

	visited := make(map[[2]int]bool)

	num1 := nums1[0]
	num2 := nums2[0]
	heap.Push(&pq, []int{num1 + num2, 0, 0})
	visited[[2]int{0, 0}] = true

	for k != 0 && len(pq) > 0 {
		popped := heap.Pop(&pq).([]int)

		result = append(result, []int{nums1[popped[1]], nums2[popped[2]]})
		k--

		i := popped[1]
		j := popped[2]
		if i+1 < len(nums1) && !visited[[2]int{i + 1, j}] {
			num1 = nums1[i+1]
			num2 = nums2[j]

			heap.Push(&pq, []int{num1 + num2, i + 1, j})
			visited[[2]int{i + 1, j}] = true
		}

		if j+1 < len(nums2) && !visited[[2]int{i, j + 1}] {
			num1 = nums1[i]
			num2 = nums2[j+1]

			heap.Push(&pq, []int{num1 + num2, i, j + 1})
			visited[[2]int{i, j + 1}] = true
		}
	}

	return result
}

// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/?envType=study-plan-v2&envId=top-interview-150
func findKPairSmallestSum() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}

	fmt.Println(kSmallestPairs(nums1, nums2, 9))
}
