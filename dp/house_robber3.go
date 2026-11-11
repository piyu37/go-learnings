package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Same idea as rob2, but on a tree: for every node we track two totals
// for its subtree, and a parent only needs these two numbers from each child.
// Time: O(n), Space: O(h) for the recursion stack (h = tree height)
func robTree(root *TreeNode) int {
	robbed, skipped := robSubtree(root)
	return max(robbed, skipped)
}

// robSubtree returns the best total for this subtree when
//   robbed  -> this node is robbed
//   skipped -> this node is not robbed
func robSubtree(node *TreeNode) (robbed, skipped int) {
	if node == nil {
		return 0, 0
	}

	leftRobbed, leftSkipped := robSubtree(node.Left)
	rightRobbed, rightSkipped := robSubtree(node.Right)

	// Rob this node: its children must be skipped.
	robbed = node.Val + leftSkipped + rightSkipped

	// Skip this node: each child is free to be robbed or not, take the better.
	skipped = max(leftRobbed, leftSkipped) + max(rightRobbed, rightSkipped)

	return robbed, skipped
}

// https://leetcode.com/problems/house-robber-iii/description/
func houseRobber3() {
	// [3,2,3,null,3,null,1]
	root1 := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}},
	}
	fmt.Println(robTree(root1)) // 7

	// [3,4,5,1,3,null,1]
	root2 := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 1}},
	}
	fmt.Println(robTree(root2)) // 9
}
