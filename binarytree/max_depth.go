package main

import "fmt"

// recursion
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return 1 + max(left, right)
}

type queue struct {
	value int
	node  *TreeNode
}

func maxDepthIterative(root *TreeNode) int {
	s := make([]queue, 0)

	if root != nil {
		s = append(s, queue{1, root})
	}

	depth := 0
	for len(s) > 0 {
		currDepth, currRoot := s[0].value, s[0].node
		s = s[1:]

		if currRoot != nil {
			depth = max(depth, currDepth)
			s = append(s, queue{currDepth + 1, currRoot.Left})
			s = append(s, queue{currDepth + 1, currRoot.Right})
		}
	}

	return depth
}

// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/?envType=study-plan-v2&envId=top-interview-150
func maxDepthMain() {
	bt := &TreeNode{}
	bt.Val = 3
	bt.Left = &TreeNode{}
	bt.Left.Val = 9
	bt.Right = &TreeNode{}
	bt.Right.Val = 20
	bt.Right.Left = &TreeNode{}
	bt.Right.Left.Val = 15
	bt.Right.Right = &TreeNode{}
	bt.Right.Right.Val = 7

	fmt.Println(maxDepth(bt))
	fmt.Println(maxDepthIterative(bt))
}
