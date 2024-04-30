package main

import "fmt"

type bim map[int][2]int

func snakesAndLadders(board [][]int) int {
	boardLen := len(board)
	queue := make([][2]int, 0)
	visited := make([]int, boardLen*boardLen+1)

	queue = append(queue, [2]int{1, 0})

	var boardIndexMap bim = make(map[int][2]int)
	boardIndexMap.initNumToBoardCordinates(board)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		square, moves := curr[0], curr[1]

		for i := 1; i <= 6; i++ {
			nextSquare := square + i

			cordinates := boardIndexMap.getNumToBoardCordinates(nextSquare)
			r, c := cordinates[0], cordinates[1]

			if board[r][c] != -1 {
				nextSquare = board[r][c]
			}

			if nextSquare == boardLen*boardLen {
				return moves + 1
			}

			if visited[nextSquare] == 0 {
				visited[nextSquare] = 1
				queue = append(queue, [2]int{nextSquare, moves + 1})
			}
		}
	}

	return -1
}

func (boardIndexMap bim) initNumToBoardCordinates(board [][]int) {
	n := 1
	evenFlag := true
	for i := len(board) - 1; i >= 0; i-- {
		for j := 0; j < len(board[0]); j++ {
			boardIndexMap[n] = [2]int{i, j}

			if evenFlag {
				n++
			} else {
				n--
			}
		}

		if evenFlag {
			n--
		} else {
			n++
		}

		n += len(board)
		evenFlag = !evenFlag
	}
}

func (boardIndexMap bim) getNumToBoardCordinates(num int) [2]int {
	return boardIndexMap[num]
}

// https://leetcode.com/problems/snakes-and-ladders/?envType=study-plan-v2&envId=top-interview-150
func snakesLadderMain() {
	board := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}

	fmt.Println(snakesAndLadders(board))

	board = [][]int{
		{-1, -1, 19, 10, -1},
		{2, -1, -1, 6, -1},
		{-1, 17, -1, 19, -1},
		{25, -1, 20, -1, -1},
		{-1, -1, -1, -1, 15},
	}

	fmt.Println(snakesAndLadders(board))
}
