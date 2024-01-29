package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	traversedNode := make([]*TreeNode, 0)

	if root == nil {
		return result
	}

	traversedNode = append(traversedNode, root)
	result = append(result, []int{root.Val})

	for {
		levelElements := make([]int, 0)
		levelElementCount := len(traversedNode)
		for levelElementCount > 0 {
			top := traversedNode[0]
			traversedNode = traversedNode[1:]

			if top.Left != nil {
				traversedNode = append(traversedNode, top.Left)
				levelElements = append(levelElements, top.Left.Val)
			}

			if top.Right != nil {
				traversedNode = append(traversedNode, top.Right)
				levelElements = append(levelElements, top.Right.Val)
			}

			levelElementCount--
		}

		if len(levelElements) > 0 {
			result = append(result, levelElements)
		} else {
			break
		}
	}

	return result
}

// program to iterate tree in level order and append the output in array according to levels
// e.g below code output => [[3], [9, 20], [15, 17]]
func levelOrderIteration() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 17,
			},
		},
	}

	fmt.Println(levelOrder(tree))
}
