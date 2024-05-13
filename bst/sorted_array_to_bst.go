package main

import "fmt"

func sortedArrayToBST(nums []int) *TreeNode {
	return convertArrayToBST(nums, 0, len(nums)-1)
}

func convertArrayToBST(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2

	node := &TreeNode{
		Val: nums[mid],
	}

	node.Left = convertArrayToBST(nums, left, mid-1)
	node.Right = convertArrayToBST(nums, mid+1, right)

	return node
}

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/?envType=study-plan-v2&envId=top-interview-150
func sortedArrayToBSTMain() {
	arr := []int{-10, -3, 0, 5, 9}
	fmt.Println(sortedArrayToBST(arr))
}
