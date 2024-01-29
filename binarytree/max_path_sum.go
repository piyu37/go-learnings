package main

import (
	"fmt"
	"math"
)

type nodeMaxPathSum struct {
	maxPathSum int
	pathTotal  int
}

func MaxPathSum(tree *BinaryTree) int {

	return calculateMaxPathSum(tree).pathTotal
}

func calculateMaxPathSum(tree *BinaryTree) *nodeMaxPathSum {
	if tree == nil {
		return &nodeMaxPathSum{0, math.MinInt32}
	}

	leftTreeMaxPathSum := calculateMaxPathSum(tree.Left)
	rightTreeMaxPathSum := calculateMaxPathSum(tree.Right)

	childMaxPathSum := max(leftTreeMaxPathSum.maxPathSum, rightTreeMaxPathSum.maxPathSum)
	finalNodeMaxPathSum := max(childMaxPathSum+tree.Value, tree.Value)
	pathTotal := max(finalNodeMaxPathSum,
		leftTreeMaxPathSum.maxPathSum+tree.Value+rightTreeMaxPathSum.maxPathSum)
	finalPathTotal := max(pathTotal, max(leftTreeMaxPathSum.pathTotal, rightTreeMaxPathSum.pathTotal))

	return &nodeMaxPathSum{finalNodeMaxPathSum, finalPathTotal}
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/08_max_path_sum
func maxPathSumMain() {
	// tree := &BinaryTree{
	// 	Value: 1,
	// 	Left: &BinaryTree{
	// 		Value: 2,
	// 		Left: &BinaryTree{
	// 			Value: 4,
	// 		},
	// 		Right: &BinaryTree{
	// 			Value: 5,
	// 		},
	// 	},
	// 	Right: &BinaryTree{
	// 		Value: 3,
	// 		Left: &BinaryTree{
	// 			Value: 6,
	// 		},
	// 		Right: &BinaryTree{
	// 			Value: 7,
	// 		},
	// 	},
	// }

	tree2 := &BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: -5,
			Left: &BinaryTree{
				Value: 6,
			},
		},
		Right: &BinaryTree{
			Value: 3,
		},
	}
	fmt.Println(MaxPathSum(tree2))
}
