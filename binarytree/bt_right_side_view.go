package main

import "fmt"

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nextQueue := make([]*TreeNode, 0)
	nextQueue = append(nextQueue, root)
	result := make([]int, 0)

	for len(nextQueue) > 0 {
		queue := nextQueue
		nextQueue = make([]*TreeNode, 0)
		flag := false
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			if !flag {
				result = append(result, curr.Val)
				flag = true
			}

			if curr.Right != nil {
				nextQueue = append(nextQueue, curr.Right)
			}

			if curr.Left != nil {
				nextQueue = append(nextQueue, curr.Left)
			}
		}
	}

	return result
}

// https://leetcode.com/problems/binary-tree-right-side-view/?envType=study-plan-v2&envId=top-interview-150
func btRightSideView() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(rightSideView(root))
}
