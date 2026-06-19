package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(M * N * 3^L)
 * ----------------------------
 * Where M is the number of rows in the board, N is the number of columns,
 * and L is the length of the string `word`.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(M * N * 3^L)
 * - The algorithm iterates through every cell in the M x N grid to find a
 * potential starting character. This gives us the initial M * N factor.
 * - From each valid starting cell, the algorithm triggers a DFS backtracking
 * search. The recursion tree for this search goes at most L levels deep.
 * - At the very first character, the algorithm can branch in up to 4 directions.
 * However, for every subsequent character, it branches in at most 3 directions
 * (because the cell it just came from is marked with '#' and is ignored).
 * - This gives an upper bound of 3^L states explored during a single DFS run.
 *
 * 2. Work Per State -> O(1)
 * - Inside the recursive `checkWordExist` function, the algorithm performs
 * constant-time operations: character comparisons, in-place board mutations
 * (`board[...][...] = '#'`), and calling `getNeighbors`.
 * - The `getNeighbors` function performs exactly 4 boundary checks and allocates
 * a tiny slice of maximum size 4. This is bounded, O(1) constant work.
 *
 * Total Time = O(M * N) starting points * O(3^L) DFS states * O(1) work = O(M * N * 3^L)
 *
 * SPACE COMPLEXITY: O(L)
 * ----------------------------
 * Where L is the length of the string `word`.
 *
 * 1. Recursion Stack: In the worst-case scenario (where a long matching prefix
 * is found), the recursion successfully matches characters one by one. The call
 * stack will go exactly L levels deep before returning. This consumes O(L) space.
 * 2. Auxiliary Space: A common trap in this problem is allocating a separate
 * M x N "visited" matrix, which would take O(M * N) space. However, this
 * algorithm cleverly mutates the original `board` in-place using a `#` marker.
 * Because no large data structures are created (only the small, 4-element
 * slices inside `getNeighbors`), the auxiliary memory usage is O(1).
 *
 * Total Space = Stack O(L) + Auxiliary O(1) = O(L)
 */
func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := 0; j < len(board[0]); j++ {
			char := board[i][j]

			if char == word[0] {
				board[i][j] = '#'
				if checkWordExist(board, i, j, 1, word) {
					return true
				}

				board[i][j] = word[0]
			}
		}
	}

	return false
}

func checkWordExist(board [][]byte, i, j, wordIdx int, word string) bool {
	if wordIdx == len(word) {
		return true
	}

	neighbours := getNeighbors(i, j, board)

	for idx := range neighbours {
		neighbour := neighbours[idx]

		if board[neighbour[0]][neighbour[1]] == word[wordIdx] {
			board[neighbour[0]][neighbour[1]] = '#'
			if checkWordExist(board, neighbour[0], neighbour[1], wordIdx+1, word) {
				return true
			}

			board[neighbour[0]][neighbour[1]] = word[wordIdx]
		}
	}

	return false
}

func getNeighbors(i, j int, arr [][]byte) [][]int {
	nodes := [][]int{}

	if j < len(arr[0])-1 && arr[i][j+1] != '#' {
		nodes = append(nodes, []int{i, j + 1})
	}

	if i < len(arr)-1 && arr[i+1][j] != '#' {
		nodes = append(nodes, []int{i + 1, j})
	}

	if j > 0 && arr[i][j-1] != '#' {
		nodes = append(nodes, []int{i, j - 1})
	}

	if i > 0 && arr[i-1][j] != '#' {
		nodes = append(nodes, []int{i - 1, j})
	}

	return nodes
}

// https://leetcode.com/problems/word-search/?envType=study-plan-v2&envId=top-interview-150
func wordSearch() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word := "ABCCED"

	fmt.Println(exist(board, word))

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word = "SEE"

	fmt.Println(exist(board, word))

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word = "ABCB"

	fmt.Println(exist(board, word))
}
