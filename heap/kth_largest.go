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

func findKthLargestOptimal(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}

	minVal, maxVal := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < minVal {
			minVal = nums[i]
		}

		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}

	arrSize := maxVal - minVal + 1
	newArr := make([]int, arrSize)

	for i := range nums {
		newArr[nums[i]-minVal]++
	}

	remain := k

	for i := arrSize - 1; i >= 0; i-- {
		if newArr[i] > 0 {
			remain -= newArr[i]

			if remain <= 0 {
				return i + minVal
			}
		}
	}

	return -1
}

// https://leetcode.com/problems/kth-largest-element-in-an-array/description/?envType=study-plan-v2&envId=top-interview-150
func kthLargest() {
	arr := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(arr, 4))
	fmt.Println(findKthLargestOptimal(arr, 4))
}
