package main

import "fmt"

type treeBalanceInfo struct {
	isBalanced bool
	height     int
}

func HeightBalancedBinaryTree(tree *BinaryTree) bool {

	return isBalance(tree).isBalanced
}

func isBalance(tree *BinaryTree) *treeBalanceInfo {
	if tree == nil {
		return &treeBalanceInfo{
			true, 0,
		}
	}

	leftTreeInfo := isBalance(tree.Left)
	rightTreeInfo := isBalance(tree.Right)

	currentHeight := 1 + max(leftTreeInfo.height, rightTreeInfo.height)

	if !leftTreeInfo.isBalanced || !rightTreeInfo.isBalanced || abs(leftTreeInfo.height-rightTreeInfo.height) > 1 {
		return &treeBalanceInfo{
			false,
			currentHeight,
		}
	}

	return &treeBalanceInfo{
		true,
		currentHeight,
	}
}

func abs(v int) int {
	if v > 0 {
		return v
	}

	return -v
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/38_height_balanced_binary_tree
func balancedBT() {
	tree := &BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
			},
			Right: &BinaryTree{
				Value: 5,
				Left: &BinaryTree{
					Value: 7,
				},
				Right: &BinaryTree{
					Value: 8,
				},
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Right: &BinaryTree{
				Value: 6,
			},
		},
	}

	fmt.Println(HeightBalancedBinaryTree(tree))
}
