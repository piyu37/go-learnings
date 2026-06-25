package main

import "fmt"

func topologicalSortBFS(graph [][]int) []int {
	indegree := make([]int, len(graph))

	for _, adj := range graph {
		for _, n := range adj {
			indegree[n]++
		}
	}

	queue := make([]int, 0)
	for i := range indegree {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0)

	for len(queue) != 0 {
		node := queue[0]

		queue = queue[1:]

		result = append(result, node)

		for _, c := range graph[node] {
			indegree[c]--

			if indegree[c] == 0 {
				queue = append(queue, c)
			}
		}
	}

	return result

}

// Kahn's Algorithm
func topologicalSortBFSMain() {
	graph := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}

	fmt.Println(topologicalSortBFS(graph))
}
