package main

import "fmt"

type weightedNode struct {
	weight   float64
	neighbor string
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	wnMap := make(map[string][]*weightedNode)
	for i := range equations {
		equation := equations[i]
		wnMap[equation[0]] = append(wnMap[equation[0]], &weightedNode{values[i], equation[1]})
		wnMap[equation[1]] = append(wnMap[equation[1]], &weightedNode{1 / values[i], equation[0]})
	}

	result := make([]float64, 0)
	for i := range queries {
		query := queries[i]
		ans := weightedBFS(query[0], query[1], wnMap)
		result = append(result, ans)

	}

	return result
}

func weightedBFS(src, target string, wnMap map[string][]*weightedNode) float64 {
	_, srcExist := wnMap[src]
	_, targetExist := wnMap[target]
	if !srcExist || !targetExist {
		return -1
	}

	queue := make([]*weightedNode, 0)
	visited := make(map[string]struct{})
	visited[src] = struct{}{}

	queue = append(queue, &weightedNode{1, src})

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n.neighbor == target {
			return n.weight
		}

		adjList := wnMap[n.neighbor]
		for i := range adjList {
			wi := adjList[i].weight
			nei := adjList[i].neighbor
			if _, ok := visited[nei]; !ok {
				queue = append(queue, &weightedNode{n.weight * wi, nei})
				visited[nei] = struct{}{}
			}
		}
	}

	return -1
}

// https://leetcode.com/problems/evaluate-division/description/?envType=study-plan-v2&envId=top-interview-150
func evaluateDivision() {
	equations := [][]string{
		{"a", "b"},
		{"b", "c"},
	}

	values := []float64{2.0, 3.0}

	queries := [][]string{
		{"a", "c"},
		{"b", "a"},
		{"a", "e"},
		{"a", "a"},
		{"x", "x"},
	}

	fmt.Println(calcEquation(equations, values, queries))

	equations = [][]string{
		{"x1", "x2"},
		{"x2", "x3"},
		{"x3", "x4"},
		{"x4", "x5"},
	}

	values = []float64{3.0, 4.0, 5.0, 6.0}

	queries = [][]string{
		{"x1", "x5"},
		{"x5", "x2"},
		{"x2", "x4"},
		{"x2", "x2"},
		{"x2", "x9"},
		{"x9", "x9"},
	}

	fmt.Println(calcEquation(equations, values, queries))
}
