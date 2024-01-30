package main

import "fmt"

func MinPassMatrix(matrix [][]int) int {

	queue := make([][]int, 0)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] > 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	queue = append(queue, []int{-1, -1})

	pass := 0

	for len(queue) > 0 {
		pos := queue[0]

		i := pos[0]
		j := pos[1]

		if i == -1 && j == -1 {
			if len(queue) == 1 && i == -1 && j == -1 {
				break
			}

			pass++
			queue = queue[1:]
			queue = append(queue, []int{-1, -1})
			continue
		}

		queue = queue[1:]

		negNeighbors := findNeighbors(i, j, matrix)

		queue = append(queue, negNeighbors...)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] < 0 {
				return -1
			}
		}
	}

	return pass
}

func findNeighbors(i, j int, arr [][]int) [][]int {
	nodes := [][]int{}

	if j < len(arr[0])-1 && arr[i][j+1] < 0 {
		nodes = append(nodes, []int{i, j + 1})
		arr[i][j+1] *= -1
	}

	if i < len(arr)-1 && arr[i+1][j] < 0 {
		nodes = append(nodes, []int{i + 1, j})
		arr[i+1][j] *= -1
	}

	if j > 0 && arr[i][j-1] < 0 {
		nodes = append(nodes, []int{i, j - 1})
		arr[i][j-1] *= -1
	}

	if i > 0 && arr[i-1][j] < 0 {
		nodes = append(nodes, []int{i - 1, j})
		arr[i-1][j] *= -1
	}

	return nodes
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/54_minimum_passes_of_matrix
func minPassMatrixMain() {
	arr := [][]int{
		{-2, -3, -4, -1, -9},
		{-4, -3, -4, -1, -2},
		{-6, -7, -2, -1, -1},
		{0, 0, 0, 0, -3},
		{1, -2, -3, -6, -1},
	}

	fmt.Println(MinPassMatrix(arr))
}
