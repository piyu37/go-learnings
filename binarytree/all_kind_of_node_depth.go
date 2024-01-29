package main

import "fmt"

func calculateSumOfSubtrees(bt *BinaryTree, depthSum, depth int) int {
	if bt == nil {
		return 0
	}

	depthSum += depth

	left := calculateSumOfSubtrees(bt.Left, depthSum, depth+1)
	right := calculateSumOfSubtrees(bt.Right, depthSum, depth+1)

	return depthSum + left + right
}

// https://github.com/lee-hen/Algoexpert/tree/master/easy/06_node_depths
func allKindOfNodeDepth() {
	bt := &BinaryTree{
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

	fmt.Println(calculateSumOfSubtrees(bt, 0, 0))
}
