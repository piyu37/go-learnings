package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTreeFromArray(arr []interface{}, index int) *TreeNode {
	if index >= len(arr) || arr[index] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[index].(int)}
	root.Left = buildTreeFromArray(arr, 2*index+1)
	root.Right = buildTreeFromArray(arr, 2*index+2)

	return root
}

func getMinimumDifference(root *TreeNode) int {
	var prev = math.MaxInt
	var minAbsoluteDiff = math.MaxInt

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)

		if prev != math.MaxInt {
			diff := abs(root.Val - prev)

			if diff < minAbsoluteDiff {
				minAbsoluteDiff = diff
			}
		}

		prev = root.Val

		traverse(root.Right)
	}

	traverse(root)

	return minAbsoluteDiff
}

func abs(val int) int {
	if val < 0 {
		return -val
	}

	return val
}

// https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/
func minAbsoluteDiffMain() {
	arr := []interface{}{1, nil, 5, nil, nil, 3}

	t := buildTreeFromArray(arr, 0)

	fmt.Println(getMinimumDifference(t))
}
