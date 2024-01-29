package main

import "fmt"

func NodeDepths(root *BinaryTree) int {
	nodeDepthsTotalCount := 0

	calculateNodeDepthsTotal(root, 0, &nodeDepthsTotalCount)

	return nodeDepthsTotalCount
}

func calculateNodeDepthsTotal(root *BinaryTree, depthCount int, nodeDepthsTotalCount *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*nodeDepthsTotalCount += depthCount

		return
	}

	*nodeDepthsTotalCount += depthCount
	calculateNodeDepthsTotal(root.Left, depthCount+1, nodeDepthsTotalCount)
	calculateNodeDepthsTotal(root.Right, depthCount+1, nodeDepthsTotalCount)
}

// // https://github.com/lee-hen/Algoexpert/tree/master/easy/06_node_depths
func nodeDepthsMain() {
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

	fmt.Println(NodeDepths(tree))
}
