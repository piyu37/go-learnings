package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N^2 * 4^(N^2))
 * ----------------------------
 * Where N is the dimension of the N x N maze.
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(4^(N^2))
 *    - At every cell in the maze, the algorithm can potentially explore up to
 *      4 directions (Down, Left, Right, Up).
 *    - Because the `visited` map prevents cycles, the absolute longest possible
 *      path would visit every single cell in the grid exactly once.
 *    - The grid has N * N = N^2 total cells.
 *    - A decision tree with a branching factor of 4 and a maximum depth of N^2
 *      yields a theoretical worst-case bound of 4^(N^2) states.
 *
 * 2. Work Per State -> O(N^2)
 *    - Boundary checks and map lookups/insertions in Go are O(1) average time.
 *    - However, doing `path+"D"` (or L/R/U) creates a completely new string in
 *      memory. Because the path string can grow up to N^2 characters long, this
 *      string concatenation takes up to O(N^2) time at every recursive step.
 *
 * Total Time = O(4^(N^2)) * O(N^2) = O(N^2 * 4^(N^2))
 *
 * SPACE COMPLEXITY: O(N^2)
 * ----------------------------
 * 1. Recursion Stack: The longest path requires the call stack to go at most
 *    N^2 levels deep.
 * 2. Auxiliary Space: The `visited` map stores coordinate pairs and can hold
 *    at most N^2 entries.
 *
 * Total Space = O(N^2) + O(N^2) = O(N^2)
 */
func ratInMaze(maze [][]int) []string {
	result := make([]string, 0)
	visited := make(map[[2]int]bool)

	var findPath func(maze [][]int, row, col int, path string)

	findPath = func(maze [][]int, row, col int, path string) {
		if row == len(maze)-1 && col == len(maze)-1 {
			result = append(result, path)
			return
		}

		if isValidPath(maze, row+1, col, visited) {
			visited[[2]int{row + 1, col}] = true
			findPath(maze, row+1, col, path+"D")
			visited[[2]int{row + 1, col}] = false
		}

		if isValidPath(maze, row, col-1, visited) {
			visited[[2]int{row, col - 1}] = true
			findPath(maze, row, col-1, path+"L")
			visited[[2]int{row, col - 1}] = false
		}

		if isValidPath(maze, row, col+1, visited) {
			visited[[2]int{row, col + 1}] = true
			findPath(maze, row, col+1, path+"R")
			visited[[2]int{row, col + 1}] = false
		}

		if isValidPath(maze, row-1, col, visited) {
			visited[[2]int{row - 1, col}] = true
			findPath(maze, row-1, col, path+"U")
			visited[[2]int{row - 1, col}] = false
		}
	}

	findPath(maze, 0, 0, "")

	return result
}

func isValidPath(maze [][]int, row, col int, visited map[[2]int]bool) bool {
	if row < 0 || row == len(maze) || col < 0 || col == len(maze) {
		return false
	}

	if visited[[2]int{row, col}] || maze[row][col] == 0 {
		return false
	}

	return true
}

// https://www.geeksforgeeks.org/problems/rat-in-a-maze-problem/1
func ratInMazeMain() {
	maze := [][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
	}

	fmt.Println(ratInMaze(maze))
}
