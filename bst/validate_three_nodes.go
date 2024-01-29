package main

import "fmt"

func ValidateThreeNodes(nodeOne *BST, nodeTwo *BST, nodeThree *BST) bool {
	if isAncestor(nodeOne, nodeTwo, nodeThree) {
		return isAncestor(nodeTwo, nodeThree, nil)
	}

	if isAncestor(nodeThree, nodeTwo, nodeOne) {
		return isAncestor(nodeTwo, nodeOne, nil)
	}

	return false
}

func isAncestor(node1, node2, node3 *BST) bool {
	curr := node1

	for curr != nil {
		if node3 != nil && curr.Value == node3.Value {
			return false
		}

		if node2.Value < curr.Value {
			curr = curr.Left
		} else if node2.Value > curr.Value {
			curr = curr.Right
		} else {
			return true
		}
	}

	return false
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/07_validate_three_nodes
func validateThreeNodesMain() {
	node2 := &BST{
		Value: 3,
	}

	node3 := &BST{
		Value: 2,
		Left: &BST{
			Value: 1,
			Left: &BST{
				Value: 0,
			},
		},
		Right: &BST{
			Value: 4,
			Left:  node2,
		},
	}
	node1 := &BST{
		Value: 5,
		Left:  node3,
		Right: &BST{
			Value: 7,
			Left: &BST{
				Value: 6,
			},
			Right: &BST{
				Value: 8,
			},
		},
	}

	fmt.Println(ValidateThreeNodes(node1, node2, node3))
}
