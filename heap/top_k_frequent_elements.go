package main

import (
	"container/heap"
	"fmt"
)

type maxHeap [][]int

func (mh maxHeap) Len() int {
	return len(mh)
}

func (mh *maxHeap) Push(x any) {
	*mh = append(*mh, x.([]int))
}

func (mh *maxHeap) Pop() any {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]

	return val
}

func (mh maxHeap) Less(i, j int) bool {
	return mh[i][1] > mh[j][1]
}

func (mh maxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func topKFrequent(nums []int, k int) []int {
	valMap := make(map[int]int)

	for i := range nums {
		valMap[nums[i]]++
	}

	maxPQ := &maxHeap{}

	for k, v := range valMap {
		heap.Push(maxPQ, []int{k, v})
	}

	result := make([]int, 0)
	for range k {
		r := heap.Pop(maxPQ).([]int)
		result = append(result, r[0])
	}

	return result
}

// https://leetcode.com/problems/top-k-frequent-elements/description/
func topKFrequentElements() {
	nums := []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}
	k := 2

	fmt.Println(topKFrequent(nums, k))
}
