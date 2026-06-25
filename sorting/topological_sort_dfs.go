package main

import "fmt"

func topologicalSortDFS(graph [][]int) []int {
	visited := make([]bool, len(graph))
	result := make([]int, 0)

	for n := range graph {
		topoDFS(n, graph, visited, &result)
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result

}

func topoDFS(node int, graph [][]int, visited []bool, result *[]int) {
	if visited[node] {
		return
	}

	visited[node] = true

	for _, c := range graph[node] {
		topoDFS(c, graph, visited, result)
	}

	*result = append(*result, node)
}

func topologicalSortDFSMain() {
	graph := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}

	fmt.Println(topologicalSortDFS(graph))
}
