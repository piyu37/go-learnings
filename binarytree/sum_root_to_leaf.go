package main

import "fmt"

func sumNumbers(root *TreeNode) int {
	return calculateSumTillLeaf(root, 0)
}

func calculateSumTillLeaf(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	sum = (sum * 10) + root.Val

	leftSum := calculateSumTillLeaf(root.Left, sum)
	rightSum := calculateSumTillLeaf(root.Right, sum)

	if leftSum+rightSum > 0 {
		return leftSum + rightSum
	}

	return sum
}

func calculateSumTillLeafIterMorris(tree *TreeNode) int {
	if tree == nil {
		return 0
	}

	curr := tree
	currNum := 0
	sum := 0
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			steps := 1

			for next.Right != nil && next.Right != curr {
				next = next.Right
				steps++
			}

			if next.Right == nil {
				next.Right = curr
				currNum = currNum*10 + curr.Val
				curr = curr.Left
			} else {
				if next.Left == nil {
					sum += currNum
				}

				for i := 0; i < steps; i++ {
					currNum /= 10
				}

				next.Right = nil
				curr = curr.Right
			}
		} else {
			currNum = currNum*10 + curr.Val
			if curr.Right == nil {
				sum += currNum
			}

			curr = curr.Right
		}
	}

	return sum
}

// https://leetcode.com/problems/sum-root-to-leaf-numbers/description/?envType=study-plan-v2&envId=top-interview-150
func sumRootToLeaf() {
	arr := []interface{}{4, 9, 0, 5, 1}

	root := buildTreeFromArray(arr, 0)

	fmt.Println(sumNumbers(root))

	fmt.Println(calculateSumTillLeafIterMorris(root))
}
