package main

import "fmt"

func dfsNeighboursXO(board [][]byte, i, j int) {
	board[i][j] = 'T'

	if j < len(board[0])-1 && board[i][j+1] == 'O' {
		dfsNeighboursXO(board, i, j+1)
	}
	if i < len(board)-1 && board[i+1][j] == 'O' {
		dfsNeighboursXO(board, i+1, j)
	}
	if j > 0 && board[i][j-1] == 'O' {
		dfsNeighboursXO(board, i, j-1)
	}
	if i > 0 && board[i-1][j] == 'O' {
		dfsNeighboursXO(board, i-1, j)
	}
}

func solve(board [][]byte) {
	i := 0
	j := 0
	for ; j < len(board[0]); j++ {
		if board[i][j] == 'O' {
			dfsNeighboursXO(board, i, j)
		}
	}

	j--
	i++
	for ; i < len(board); i++ {
		if board[i][j] == 'O' {
			dfsNeighboursXO(board, i, j)
		}
	}

	i--
	j--
	for ; j >= 0; j-- {
		if board[i][j] == 'O' {
			dfsNeighboursXO(board, i, j)
		}
	}

	j++
	i--
	for ; i > 0; i-- {
		if board[i][j] == 'O' {
			dfsNeighboursXO(board, i, j)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}

// https://leetcode.com/problems/surrounded-regions/description/?envType=study-plan-v2&envId=top-interview-150
func surroundedRegions() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	solve(board)

	fmt.Println(board)
	board = [][]byte{
		{'O', 'O'},
		{'O', 'O'},
	}

	solve(board)

	fmt.Println(board)

}
