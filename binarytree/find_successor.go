package main

import "fmt"

// This is an input class. Do not edit.
type BinaryTreeWithParent struct {
	Value int

	Left   *BinaryTreeWithParent
	Right  *BinaryTreeWithParent
	Parent *BinaryTreeWithParent
}

func FindSuccessor(tree *BinaryTreeWithParent, node *BinaryTreeWithParent) *BinaryTreeWithParent {
	if tree == nil || node == nil {
		return nil
	}

	if node.Right != nil {
		return findLeftMostNode(node.Right)
	}

	if node.Parent == nil {
		return nil
	}

	if node == node.Parent.Left {
		return node.Parent
	}

	return node.Parent.Parent
}

func findLeftMostNode(tree *BinaryTreeWithParent) *BinaryTreeWithParent {
	if tree.Left != nil {
		return findLeftMostNode(tree.Left)
	}

	return tree
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/36_find_successor
func findSuccessorMain() {
	tree := &BinaryTreeWithParent{
		Value: 1,
		Left: &BinaryTreeWithParent{
			Value: 2,
			Left: &BinaryTreeWithParent{
				Value: 4,
			},
			Right: &BinaryTreeWithParent{
				Value: 5,
			},
		},
		Right: &BinaryTreeWithParent{
			Value: 3,
			Left: &BinaryTreeWithParent{
				Value: 6,
			},
			Right: &BinaryTreeWithParent{
				Value: 7,
			},
		},
	}

	tree.Left.Parent = tree
	tree.Left.Left.Parent = tree.Left
	tree.Left.Right.Parent = tree.Left
	tree.Right.Parent = tree
	tree.Right.Left.Parent = tree.Right
	tree.Right.Right.Parent = tree.Right

	fmt.Println(FindSuccessor(tree, tree.Right.Left))

}
