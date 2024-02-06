package main

import "fmt"

func flattend_bt_preorder(root *TreeNode) {
	curr := root

	for curr != nil {
		if curr.Left != nil {
			prev := curr.Left

			for prev.Right != nil {
				prev = prev.Right
			}

			prev.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}

		curr = curr.Right
	}
}

// program to convert tree into linked list in preorder fashion
// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/?envType=study-plan-v2&envId=top-interview-150
func flattendBTPreorderMain() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}

	flattend_bt_preorder(tree)
	fmt.Println(tree)

}
