package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
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
