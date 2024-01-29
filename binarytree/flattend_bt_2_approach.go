package main

import "fmt"

func (tree *BinaryTree) dfs() {
	if tree == nil {
		return
	}

	tree.Left.dfs()
	if tree.Left != nil {
		rightMost := getRightmostChild(tree.Left)
		rightMost.Right = tree
		tree.Left = rightMost
	}

	tree.Right.dfs()
	if tree.Right != nil {
		leftMost := getLeftmostChild(tree.Right)
		tree.Right = leftMost
		leftMost.Left = tree
	}
}

func getRightmostChild(node *BinaryTree) *BinaryTree {
	var currentNode = node
	for currentNode.Right != nil {
		currentNode = currentNode.Right
	}
	return currentNode
}

func getLeftmostChild(node *BinaryTree) *BinaryTree {
	var currentNode = node
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/08_flatten_binary_tree
func flattendBT2Approach() {
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
			Left: &BinaryTree{
				Value: 6,
			},
		},
	}

	tree.dfs()
	fmt.Println(tree)
}
