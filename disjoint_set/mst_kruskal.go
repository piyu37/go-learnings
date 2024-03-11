package main

import (
	"fmt"
	"sort"
)

type node struct {
	weight       int
	name, parent int
}

func mstKruskal(graph [][][]int) ([][]int, int) {
	mstArray := make([][]int, 0)
	mstWeight := 0

	nodes := make([]node, 0)

	for i := range graph {
		edges := graph[i]

		for j := range edges {
			nodes = append(nodes, node{edges[j][0], i, edges[j][1]})
		}
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].weight < nodes[j].weight
	})

	ds := initDisjointSet(len(graph))

	for i := range nodes {
		wei := nodes[i].weight
		u := nodes[i].name
		v := nodes[i].parent

		if ds.findUPar(u) != ds.findUPar(v) {
			mstWeight += wei
			mstArray = append(mstArray, []int{u, v})
			ds.unionBySize(u, v)
		}
	}

	return mstArray, mstWeight
}

func mstKruskalMain() {
	graph := [][][]int{
		{{2, 1}, {1, 3}, {4, 4}},
		{{2, 0}, {3, 2}, {3, 3}, {7, 5}},
		{{3, 1}, {5, 3}, {8, 5}},
		{{1, 0}, {3, 1}, {5, 2}, {9, 4}},
		{{4, 0}, {9, 3}},
		{{7, 1}, {8, 2}},
	}

	fmt.Println(mstKruskal(graph))
}
