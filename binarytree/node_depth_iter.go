package main

import "fmt"

type level struct {
	node  *BinaryTree
	depth int
}

func NodeDepthsIter(root *BinaryTree) int {
	totalDepth := 0

	stack := []level{
		{
			node:  root,
			depth: 0,
		},
	}

	var top level

	for len(stack) > 0 {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if top.node == nil {
			continue
		}

		totalDepth += top.depth

		stack = append(stack, level{node: top.node.Left, depth: top.depth + 1})
		stack = append(stack, level{node: top.node.Right, depth: top.depth + 1})
	}

	return totalDepth
}

// https://github.com/lee-hen/Algoexpert/tree/master/easy/06_node_depths
func nodeDepthsIterMain() {
	tree := &BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
				Left: &BinaryTree{
					Value: 8,
				},
				Right: &BinaryTree{
					Value: 9,
				},
			},
			Right: &BinaryTree{
				Value: 5,
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 6,
			},
			Right: &BinaryTree{
				Value: 7,
			},
		},
	}

	fmt.Println(NodeDepthsIter(tree))
}
