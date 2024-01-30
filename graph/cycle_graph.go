package main

import "fmt"

func CheckCycleInGraph(edges [][]int) bool {
	colorNode := make(map[int]string)

	return dfsTraverse(0, edges, colorNode)
}

func dfsTraverse(start int, edges [][]int, colorNode map[int]string) bool {
	if colorNode[start] == "g" {
		return true
	}

	colorNode[start] = "g"

	for _, chidNode := range edges[start] {
		cycleExist := dfsTraverse(chidNode, edges, colorNode)

		if cycleExist {
			return true
		}
	}

	colorNode[start] = "b"

	return false
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/34_cycle_in_graph
func cycleGraph() {
	edges := [][]int{
		{1, 3},
		{2, 3, 4},
		{0},
		{},
		{2, 5},
		{},
	}

	fmt.Println(CheckCycleInGraph(edges))
}
