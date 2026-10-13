package main

import (
	"container/heap"
	"fmt"
)

type minHeap1 []int

type maxHeap1 struct{ minHeap1 }

func (h maxHeap1) Less(i, j int) bool {
	return h.minHeap1[i] > h.minHeap1[j]
}

func (mh minHeap1) Len() int {
	return len(mh)
}

func (mh minHeap1) Less(i, j int) bool {
	return mh[i] < mh[j]
}

func (mh minHeap1) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minHeap1) Push(val any) {
	*mh = append(*mh, val.(int))
}

func (mh *minHeap1) Pop() any {
	top := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return top
}

type MedianFinder struct {
	maxH *maxHeap1
	minH *minHeap1
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxH: &maxHeap1{},
		minH: &minHeap1{},
	}
}

// AddNum keeps the smaller half of the numbers in maxH and the larger half in minH,
// with len(maxH) == len(minH) or len(minH)+1.
// 1. Push num onto the correct heap: maxH if num <= maxH's top, otherwise minH.
// 2. Rebalance by moving one top value if either heap has grown too big.
// Cost: 1 push, plus 1 pop + 1 push when rebalancing (1-3 heap ops, each O(log n)).

// func (this *MedianFinder) AddNum(num int) {
// 	if this.maxH.Len() == 0 || num <= (this.maxH.minHeap1)[0] {
// 		heap.Push(this.maxH, num)
// 	} else {
// 		heap.Push(this.minH, num)
// 	}

// 	// keep maxH.Len() == minH.Len() or minH.Len()+1
// 	if this.maxH.Len() > this.minH.Len()+1 {
// 		heap.Push(this.minH, heap.Pop(this.maxH))
// 	} else if this.minH.Len() > this.maxH.Len() {
// 		heap.Push(this.maxH, heap.Pop(this.minH))
// 	}
// }

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxH, num)

	heap.Push(this.minH, heap.Pop(this.maxH))

	if this.maxH.Len() < this.minH.Len() {
		heap.Push(this.maxH, heap.Pop(this.minH))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxH.Len() > this.minH.Len() {
		return float64(this.maxH.minHeap1[0])
	}

	return 0.5 * (float64(this.maxH.minHeap1[0]) + float64((*this.minH)[0]))
}

// https://leetcode.com/problems/find-median-from-data-stream/
func medianDataStream() {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	fmt.Println(mf.FindMedian()) // 1.5
	mf.AddNum(3)
	fmt.Println(mf.FindMedian()) // 2
}
