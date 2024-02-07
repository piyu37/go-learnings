package main

import "fmt"

func inorder(tree *TreeNode) []int {
	if tree == nil {
		return nil
	}

	inorderArr := []int{}

	curr := tree

	for curr != nil {
		if curr.Left != nil {
			next := curr.Left

			for next.Right != nil && next.Right != curr {
				next = next.Right
			}

			if next.Right == curr {
				next.Right = nil
				inorderArr = append(inorderArr, curr.Val)
				curr = curr.Right
			} else {
				next.Right = curr
				curr = curr.Left
			}
		} else {
			inorderArr = append(inorderArr, curr.Val)
			curr = curr.Right
		}
	}

	return inorderArr
}

func preorder(tree *TreeNode) []int {
	if tree == nil {
		return nil
	}

	result := []int{}

	curr := tree

	for curr != nil {
		if curr.Left != nil {
			next := curr.Left

			for next.Right != nil && next.Right != curr {
				next = next.Right
			}

			if next.Right == curr {
				next.Right = nil
				curr = curr.Right
			} else {
				next.Right = curr
				result = append(result, curr.Val)
				curr = curr.Left
			}
		} else {
			result = append(result, curr.Val)
			curr = curr.Right
		}
	}

	return result
}

func morrisTraversal() {
	arr := []interface{}{1, 2, 3, 4, 5, nil, nil, nil, nil, nil, 6}

	t := buildTreeFromArray(arr, 0)

	fmt.Println(inorder(t))
	fmt.Println(preorder(t))
}
