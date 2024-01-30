package main

import "fmt"

func RemoveIslands(arr [][]int) [][]int {
	i := 0
	for j := range arr[i] {
		if arr[i][j] == 2 || arr[i][j] == 0 {
			continue
		}

		arr = markNonIsland(i, j, arr)
	}
	i = len(arr) - 1
	for j := range arr[i] {
		if arr[i][j] == 2 || arr[i][j] == 0 {
			continue
		}

		arr = markNonIsland(i, j, arr)
	}

	j := 0
	for i := range arr {
		if arr[i][j] == 2 || arr[i][j] == 0 {
			continue
		}

		arr = markNonIsland(i, j, arr)
	}

	j = len(arr[0]) - 1
	for i := range arr {
		if arr[i][j] == 2 || arr[i][j] == 0 {
			continue
		}

		arr = markNonIsland(i, j, arr)
	}

	for i := range arr {
		for j := range arr[0] {
			if arr[i][j] == 1 {
				arr[i][j] = 0
			}

			if arr[i][j] == 2 {
				arr[i][j] = 1
			}
		}
	}

	return arr
}

func markNonIsland(i, j int, arr [][]int) [][]int {
	neighborNodes := [][]int{{i, j}}
	for len(neighborNodes) > 0 {
		arrPos := neighborNodes[0]
		i := arrPos[0]
		j := arrPos[1]

		neighborNodes = neighborNodes[1:]

		if arr[i][j] == 1 {
			arr[i][j] = 2
			neighbors := getNeighbors(i, j, arr)

			neighborNodes = append(neighborNodes, neighbors...)
		}
	}

	return arr
}

func getNeighbors(i, j int, arr [][]int) [][]int {
	nodes := [][]int{}

	if j < len(arr[0])-1 && arr[i][j+1] == 1 {
		nodes = append(nodes, []int{i, j + 1})
	}

	if i < len(arr)-1 && arr[i+1][j] == 1 {
		nodes = append(nodes, []int{i + 1, j})
	}

	if j > 0 && arr[i][j-1] == 1 {
		nodes = append(nodes, []int{i, j - 1})
	}

	if i > 0 && arr[i-1][j] == 1 {
		nodes = append(nodes, []int{i - 1, j})
	}

	return nodes
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/43_remove_islands
func removeIslandsMain() {
	arr := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 0, 1, 0, 1, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 1},
	}

	fmt.Println(RemoveIslands(arr))
}
