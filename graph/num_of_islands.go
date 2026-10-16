package main

import "fmt"

func numIslands(grid [][]byte) int {
	totalIslands := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				markIsland(i, j, grid)
				totalIslands++
			}
		}
	}
	return totalIslands
}

var directions = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func markIsland(i, j int, grid [][]byte) {
	grid[i][j] = '2'
	queue := [][2]int{{i, j}}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, d := range directions {
			r, c := cell[0]+d[0], cell[1]+d[1]
			if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == '1' {
				grid[r][c] = '2'
				queue = append(queue, [2]int{r, c})
			}
		}
	}
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
