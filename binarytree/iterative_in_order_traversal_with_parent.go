package main

import "fmt"

func (tree *BinaryTreeWithParent) IterativeInOrderTraversal(callback func(int)) {
	if tree == nil {
		return
	}

	currentNode := tree
	var nextNode *BinaryTreeWithParent

	prevNode := tree.Parent

	for currentNode != nil {
		if currentNode.Parent == prevNode {
			if currentNode.Left != nil {
				nextNode = currentNode.Left
				prevNode = currentNode
				currentNode = nextNode

				continue
			} else {
				callback(currentNode.Value)
				if currentNode.Right != nil {
					nextNode = currentNode.Right
					prevNode = currentNode
					currentNode = nextNode
				} else {
					nextNode = prevNode
					prevNode = currentNode
					currentNode = nextNode
				}
			}
		} else {
			if prevNode == currentNode.Left {
				callback(currentNode.Value)

				if currentNode.Right != nil {
					nextNode = currentNode.Right
					prevNode = currentNode
					currentNode = nextNode
				} else {
					nextNode = currentNode.Parent
					prevNode = currentNode
					currentNode = nextNode
				}
			} else {
				nextNode = currentNode.Parent
				prevNode = currentNode
				currentNode = nextNode
			}
		}
	}
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/07_iterative_in_order_traversal
func IterativeInOrderTraversalWithParent() {
	tree := &BinaryTreeWithParent{
		Value: 1,
		Left: &BinaryTreeWithParent{
			Value: 2,
			Left: &BinaryTreeWithParent{
				Value: 4,
				Right: &BinaryTreeWithParent{
					Value: 9,
				},
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
	tree.Left.Left.Right.Parent = tree.Left.Left
	tree.Right.Parent = tree
	tree.Right.Left.Parent = tree.Right
	tree.Right.Right.Parent = tree.Right

	function := func(v int) {
		fmt.Println(v)
	}

	tree.IterativeInOrderTraversal(function)
}
