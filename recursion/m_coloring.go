package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(V * M^V)
 * ----------------------------
 * Where:
 * - V = Total number of vertices in the graph.
 * - M = Total number of available colors.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(M^V)
 *    - For every vertex (out of V total vertices), the algorithm can attempt to
 *      assign up to M different colors.
 *    - This creates a decision tree with a branching factor of M and a maximum
 *      depth of V, resulting in at most M^V possible states in the worst case.
 *
 * 2. Work Per State -> O(V)
 *    - At each recursive step, `isValidColor()` checks the adjacency list `edges[n]`
 *      to ensure no adjacent vertex has the same color.
 *    - In the worst case (a dense graph where a vertex is connected to all other V-1
 *      vertices), this loop takes O(V) time.
 *
 * Total Time = O(M^V) * O(V) = O(V * M^V)
 *
 * SPACE COMPLEXITY: O(V)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes at most V levels deep (one for each vertex).
 * 2. Auxiliary Space: The `coloredV` array holds exactly V integers to track colors.
 *
 * Total Space = O(V) + O(V) = O(V)
 */
func graphColoring(v int, edges [][]int, m int) bool {
	return isColoringPossible(v, edges, m, 0, make([]int, v))
}

func isColoringPossible(v int, edges [][]int, m int, n int, coloredV []int) bool {
	if n == v {
		return true
	}

	for color := 1; color <= m; color++ {
		if isValidColor(color, edges, n, coloredV) {
			coloredV[n] = color
			if isColoringPossible(v, edges, m, n+1, coloredV) {
				return true
			}

			coloredV[n] = 0
		}
	}

	return false
}

func isValidColor(color int, edges [][]int, n int, coloredV []int) bool {
	for _, vertex := range edges[n] {
		if coloredV[vertex] == color {
			return false
		}
	}

	return true
}

// https://www.geeksforgeeks.org/problems/m-coloring-problem-1587115620/1
func mColoring() {
	v := 4
	edges := [][]int{{1, 2, 3}, {0, 3}, {0, 3}, {0, 1, 2}}
	m := 3

	fmt.Println(graphColoring(v, edges, m))

	v = 3
	edges = [][]int{{1, 2}, {0, 2}, {0, 1}}
	m = 2
	fmt.Println(graphColoring(v, edges, m))
}
