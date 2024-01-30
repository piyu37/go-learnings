package main

import "fmt"

func RiverSizes(matrix [][]int) []int {
	visited := make([][]bool, len(matrix))
	riverSize := make([]int, 0)

	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				continue
			}

			if visited[i][j] {
				continue
			}

			riverSize = traverseNode(i, j, matrix, visited, riverSize)

		}
	}
	return riverSize
}

func traverseNode(i, j int, matrix [][]int, visited [][]bool, riverSize []int) []int {
	nodeToExplore := [][]int{{i, j}}
	currentRiverSize := 0

	for len(nodeToExplore) > 0 {
		node := nodeToExplore[0]

		i := node[0]
		j := node[1]

		nodeToExplore = nodeToExplore[1:]

		if visited[i][j] {
			continue
		}

		visited[i][j] = true

		currentRiverSize++

		neighborNodes := getNeighborNodees(i, j, matrix, visited)

		nodeToExplore = append(nodeToExplore, neighborNodes...)
	}

	if currentRiverSize > 0 {
		riverSize = append(riverSize, currentRiverSize)
	}

	return riverSize
}

func getNeighborNodees(i, j int, matrix [][]int, visited [][]bool) [][]int {
	nodes := make([][]int, 0)

	if j < len(matrix[0])-1 && matrix[i][j+1] == 1 && !visited[i][j+1] {
		nodes = append(nodes, []int{i, j + 1})
	}

	if i < len(matrix)-1 && matrix[i+1][j] == 1 && !visited[i+1][j] {
		nodes = append(nodes, []int{i + 1, j})
	}

	if j > 0 && matrix[i][j-1] == 1 && !visited[i][j-1] {
		nodes = append(nodes, []int{i, j - 1})
	}

	if i > 0 && matrix[i-1][j] == 1 && !visited[i-1][j] {
		nodes = append(nodes, []int{i - 1, j})
	}

	return nodes
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/01_river_sizes
func riverSizeMain() {
	arr := [][]int{
		{1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0},
		{1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0},
		{0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0},
		{1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 1},
	}

	fmt.Println(RiverSizes(arr))
}
