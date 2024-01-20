package main

import "fmt"

// Do not edit the class below except for the buildHeap,
// siftDown, siftUp, peek, remove, and insert methods.
// Feel free to add new properties and methods to the class.
type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	// Do not edit the lines below.
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) BuildHeap(array []int) {
	firstParent := (len(array) - 2) / 2

	for i := firstParent; i >= 0; i-- {
		h.siftDown(i, len(array)-1)
	}
}

func (h *MinHeap) siftDown(currentIndex, endIndex int) {
	for currentIndex < endIndex {
		childIdx1 := 2*currentIndex + 1
		childIdx2 := 2*currentIndex + 2

		var minChildIdx int

		if childIdx1 <= endIndex && childIdx2 <= endIndex {
			minChildIdx = childIdx1
			if (*h)[childIdx2] < (*h)[childIdx1] {
				minChildIdx = childIdx2
			}
		} else if childIdx1 > endIndex && childIdx2 > endIndex {
			return
		} else if childIdx1 > endIndex {
			minChildIdx = childIdx2
		} else {
			minChildIdx = childIdx1
		}

		if (*h)[currentIndex] > (*h)[minChildIdx] {
			h.swap(currentIndex, minChildIdx)
			currentIndex = minChildIdx
		} else {
			return
		}
	}
}

func (h *MinHeap) siftUp() {
	hLen := len(*h)

	currentIdx := hLen - 1
	parentIdx := (currentIdx - 1) / 2

	for currentIdx > 0 && (*h)[parentIdx] > (*h)[currentIdx] {
		h.swap(currentIdx, parentIdx)

		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}

func (h MinHeap) Peek() int {
	// Write your code here.
	return h[0]
}

func (h *MinHeap) Remove() int {
	// Write your code here.

	hLen := len(*h)

	h.swap(0, hLen-1)
	peeked := (*h)[hLen-1]

	*h = (*h)[:hLen-1]

	h.siftDown(0, hLen-2)

	return peeked
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)

	h.siftUp()
}

func (h MinHeap) swap(idx1, idx2 int) {
	h[idx1], h[idx2] = h[idx2], h[idx1]
}

// code to create heap from array and without array + remove element from heap
func heapMain() {
	arr := []int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41}

	fmt.Println(NewMinHeap(arr))

	arr = make([]int, 0)

	h := MinHeap(arr)

	h.Insert(5)
	h.Insert(8)
	h.Insert(9)
	h.Insert(4)
	h.Remove()
}
