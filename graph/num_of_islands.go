package main

import "fmt"

func numIslands(grid [][]byte) int {
	totalIslands := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '0' || grid[i][j] == '2' {
				continue
			}

			grid = markIsland(i, j, grid)
			totalIslands++
		}
	}
	return totalIslands
}

func markIsland(i, j int, arr [][]byte) [][]byte {
	neighborNodes := [][]int{{i, j}}
	for len(neighborNodes) > 0 {
		arrPos := neighborNodes[0]
		i := arrPos[0]
		j := arrPos[1]

		neighborNodes = neighborNodes[1:]

		if arr[i][j] == '1' {
			arr[i][j] = '2'
			neighbors := get01Neighbors(i, j, arr)

			neighborNodes = append(neighborNodes, neighbors...)
		}
	}

	return arr
}

func get01Neighbors(i, j int, arr [][]byte) [][]int {
	nodes := [][]int{}

	if j < len(arr[0])-1 && arr[i][j+1] == '1' {
		nodes = append(nodes, []int{i, j + 1})
	}

	if i < len(arr)-1 && arr[i+1][j] == '1' {
		nodes = append(nodes, []int{i + 1, j})
	}

	if j > 0 && arr[i][j-1] == '1' {
		nodes = append(nodes, []int{i, j - 1})
	}

	if i > 0 && arr[i-1][j] == '1' {
		nodes = append(nodes, []int{i - 1, j})
	}

	return nodes
}

// https://leetcode.com/problems/number-of-islands/submissions/1170503021/?envType=study-plan-v2&envId=top-interview-150
func numOfIslandsMain() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Println(numIslands(grid))

	grid = [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}

	fmt.Println(numIslands(grid))
}
