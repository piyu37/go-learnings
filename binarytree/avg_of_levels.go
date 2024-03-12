package main

import "fmt"

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	nextQueue := make([]*TreeNode, 0)
	nextQueue = append(nextQueue, root)
	result := make([]float64, 0)

	for len(nextQueue) > 0 {
		total := 0.0
		noOfNodes := len(nextQueue)
		queue := nextQueue
		nextQueue = make([]*TreeNode, 0)
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			total += float64(curr.Val)

			if curr.Left != nil {
				nextQueue = append(nextQueue, curr.Left)
			}

			if curr.Right != nil {
				nextQueue = append(nextQueue, curr.Right)
			}
		}

		result = append(result, total/float64(noOfNodes))
	}

	return result
}

// https://leetcode.com/problems/average-of-levels-in-binary-tree/description/?envType=study-plan-v2&envId=top-interview-150
func avgOfLevelsMain() {
	root := &TreeNode{
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
				Val: 7,
			},
		},
	}

	fmt.Println(averageOfLevels(root))
}
