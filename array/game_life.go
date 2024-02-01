package main

import "fmt"

// 0 => -10
// 1 => 10
func gameOfLife(board [][]int) {
	colLen := len(board[0])
	for i := 0; i < len(board); i++ {
		for j := 0; j < colLen; j++ {
			liveNeighbours := 0
			if i-1 >= 0 {
				if j-1 >= 0 && board[i-1][j-1] > 0 {
					liveNeighbours++
				}

				if j+1 < colLen && board[i-1][j+1] > 0 {
					liveNeighbours++
				}

				if board[i-1][j] > 0 {
					liveNeighbours++
				}
			}

			if j-1 >= 0 && board[i][j-1] > 0 {
				liveNeighbours++
			}

			if j+1 < colLen && board[i][j+1] > 0 {
				liveNeighbours++
			}

			if i+1 < len(board) {
				if j-1 >= 0 && board[i+1][j-1] > 0 {
					liveNeighbours++
				}

				if j+1 < colLen && board[i+1][j+1] > 0 {
					liveNeighbours++
				}

				if board[i+1][j] > 0 {
					liveNeighbours++
				}
			}

			switch board[i][j] {
			case 0:
				if liveNeighbours == 3 {
					board[i][j] = -10
				}

			case 1:
				if liveNeighbours < 2 {
					board[i][j] = 10
				} else if liveNeighbours > 3 {
					board[i][j] = 10
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < colLen; j++ {
			if board[i][j] == -10 {
				board[i][j] = 1
			} else if board[i][j] == 10 {
				board[i][j] = 0
			}
		}
	}
}

// https://leetcode.com/problems/game-of-life/description/?envType=study-plan-v2&envId=top-interview-150
func gameLife() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
	fmt.Println(board)
}
