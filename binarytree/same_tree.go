package main

import "fmt"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return true
}

type dequeue struct {
	node1, node2 *TreeNode
}

func isSameTreeIterative(p *TreeNode, q *TreeNode) bool {
	deq := make([]dequeue, 0)
	deq = append(deq, dequeue{p, q})

	for len(deq) > 0 {
		p, q = deq[0].node1, deq[0].node2
		deq = deq[1:]

		if !check(p, q) {
			return false
		}

		if p != nil {
			deq = append(deq, dequeue{p.Left, q.Left})
			deq = append(deq, dequeue{p.Right, q.Right})
		}
	}

	return true
}

// https://leetcode.com/problems/same-tree/?envType=study-plan-v2&envId=top-interview-150
func isSameTreeMain() {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(isSameTree(p, q))
	fmt.Println(isSameTreeIterative(p, q))
}
