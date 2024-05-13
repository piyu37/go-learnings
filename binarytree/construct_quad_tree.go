package main

import "fmt"

type QuadNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadNode
	TopRight    *QuadNode
	BottomLeft  *QuadNode
	BottomRight *QuadNode
}

func construct(grid [][]int) *QuadNode {
	isLeaf := isSameValueInGrid(grid, 0, len(grid), 0, len(grid[0]))

	if isLeaf {
		return &QuadNode{
			Val:    grid[0][0] == 1,
			IsLeaf: true,
		}
	}

	n := &QuadNode{
		Val:    true,
		IsLeaf: false,
	}

	return constructTree(grid, n, 0, len(grid), 0, len(grid[0]))
}

func isSameValueInGrid(grid [][]int, startRow, endRow, startCol, endCol int) bool {
	val := grid[startRow][startCol]
	for i := startRow; i < endRow; i++ {
		for j := startCol; j < endCol; j++ {
			if grid[i][j] != val {
				return false
			}
		}
	}

	return true
}

func constructTree(grid [][]int, n *QuadNode, startRow, endRow, startCol, endCol int) *QuadNode {
	isTopLeftLeaf := isSameValueInGrid(grid, startRow, (startRow+endRow)/2, startCol, (startCol+endCol)/2)
	isTopRightLeaf := isSameValueInGrid(grid, startRow, (startRow+endRow)/2, (startCol+endCol)/2, endCol)
	isBottomLeftLeaf := isSameValueInGrid(grid, (startRow+endRow)/2, endRow, startCol, (startCol+endCol)/2)
	isBottomRightLeaf := isSameValueInGrid(grid, (startRow+endRow)/2, endRow, (startCol+endCol)/2, endCol)

	if isTopLeftLeaf {
		n.TopLeft = &QuadNode{
			Val:    grid[startRow][startCol] == 1,
			IsLeaf: true,
		}
	} else {
		n.TopLeft = constructTree(grid, &QuadNode{Val: true, IsLeaf: false},
			startRow, endRow/2, startCol, endCol/2)
	}

	if isTopRightLeaf {
		n.TopRight = &QuadNode{
			Val:    grid[startRow][endCol/2] == 1,
			IsLeaf: true,
		}
	} else {
		n.TopRight = constructTree(grid, &QuadNode{Val: true, IsLeaf: false},
			startRow, endRow/2, endCol/2, endCol)
	}

	if isBottomLeftLeaf {
		n.BottomLeft = &QuadNode{
			Val:    grid[endRow-1][startCol] == 1,
			IsLeaf: true,
		}
	} else {
		n.BottomLeft = constructTree(grid, &QuadNode{Val: true, IsLeaf: false},
			startRow/2, endRow, startCol, endCol/2)
	}

	if isBottomRightLeaf {
		n.BottomRight = &QuadNode{
			Val:    grid[endRow-1][endCol-1] == 1,
			IsLeaf: true,
		}
	} else {
		n.BottomRight = constructTree(grid, &QuadNode{Val: true, IsLeaf: false},
			startRow/2, endRow, endCol/2, endCol)
	}

	return n
}

func constructOptimized(grid [][]int) *QuadNode {
	return dfs(grid, len(grid), 0, 0)
}

func dfs(grid [][]int, n, r, c int) *QuadNode {
	allSame := true
	val := grid[r][c]
	for i := r; i < r+n; i++ {
		for j := c; j < c+n; j++ {
			if grid[i][j] != val {
				allSame = false
				break
			}
		}
	}

	if allSame {
		return &QuadNode{Val: grid[r][c] == 1, IsLeaf: true}
	}

	n = n / 2

	topLeft := dfs(grid, n, r, c)
	topRight := dfs(grid, n, r, c+n)
	bottomLeft := dfs(grid, n, r+n, c)
	bottomRight := dfs(grid, n, r+n, c+n)

	return &QuadNode{Val: true, IsLeaf: false, TopLeft: topLeft, TopRight: topRight,
		BottomLeft: bottomLeft, BottomRight: bottomRight}
}

// https://leetcode.com/problems/construct-quad-tree/description/?envType=study-plan-v2&envId=top-interview-150
func constructQuadTree() {
	grid := [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}

	fmt.Println(construct(grid))
	fmt.Println(constructOptimized(grid))

	grid = [][]int{
		{0, 1},
		{1, 0},
	}

	fmt.Println(construct(grid))
	fmt.Println(constructOptimized(grid))
}
