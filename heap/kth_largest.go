package main

import (
	"container/heap"
	"fmt"
)

type heapInt []int

func (x heapInt) Len() int           { return len(x) }
func (x heapInt) Less(i, j int) bool { return x[i] > x[j] }
func (x heapInt) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (h *heapInt) Pop() any {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}

func (h *heapInt) Push(x any) {
	*h = append(*h, x.(int))
}

func findKthLargest(nums []int, k int) int {
	numsHeap := heapInt(nums)
	heap.Init(&numsHeap)

	result := -1
	for i := 0; i < k && i < len(nums); i++ {
		numsHeapTemp := numsHeap[:len(nums)-i]
		result = heap.Pop(&numsHeapTemp).(int)
	}

	return result
}

// https://leetcode.com/problems/kth-largest-element-in-an-array/description/?envType=study-plan-v2&envId=top-interview-150
func kthLargest() {
	arr := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(arr, 4))
}
