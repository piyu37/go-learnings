package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	queue := []*TreeNode{root}
	reverse := false
	for len(queue) != 0 {
		size := len(queue)
		level := make([]int, 0, size)
		for i := 1; i <= size; i++ {
			cur := queue[0]
			queue = queue[1:]

			if reverse {
				level = append([]int{cur.Val}, level...)
			} else {
				level = append(level, cur.Val)
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

		}
		reverse = !reverse
		res = append(res, level)
	}
	return res
}

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/?envType=study-plan-v2&envId=top-interview-150
func zigzagLevelOrderTraversal() {
	arr := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := buildTreeFromArray(arr, 0)

	fmt.Println(zigzagLevelOrder(root))
}
