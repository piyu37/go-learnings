package main

import (
	"container/heap"
	"fmt"
)

type node struct {
	weight       int
	name, parent int
}

type minHeap []node

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].weight < h[j].weight
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(v any) {
	*h = append(*h, v.(node))
}

func (h *minHeap) Pop() any {
	arr := *h
	val := arr[len(arr)-1]
	*h = arr[:len(arr)-1]

	return val
}

func mstPrimms(graph [][][]int) ([][]int, int) {
	var pq minHeap

	heap.Init(&pq)

	minSpanningArr := make([][]int, 0)
	minSpanningWeight := 0
	visited := make([]int, len(graph))
	heap.Push(&pq, node{0, 0, -1})

	for len(pq) > 0 {
		curr := heap.Pop(&pq).(node)

		if visited[curr.name] == 1 {
			continue
		}

		visited[curr.name] = 1
		if curr.parent != -1 {
			minSpanningArr = append(minSpanningArr, []int{curr.parent, curr.name})
			minSpanningWeight += curr.weight
		}

		neighbours := graph[curr.name]

		for i := range neighbours {
			neighbour := neighbours[i]
			wei := neighbour[0]
			adjNode := neighbour[1]
			if visited[adjNode] == 0 {
				heap.Push(&pq, node{wei, adjNode, curr.name})
			}
		}
	}

	return minSpanningArr, minSpanningWeight
}

// understand what's fibbonaci heap
func mstPrimmsMain() {
	graph := [][][]int{
		{{2, 1}, {1, 3}, {4, 4}},
		{{2, 0}, {3, 2}, {3, 3}, {7, 5}},
		{{3, 1}, {5, 3}, {8, 5}},
		{{1, 0}, {3, 1}, {5, 2}, {9, 4}},
		{{4, 0}, {9, 3}},
		{{7, 1}, {8, 2}},
	}

	fmt.Println(mstPrimms(graph))
}
