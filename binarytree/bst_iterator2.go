package main

type BSTIterator2 struct {
	stack *[]TreeNode
}

func (bsti *BSTIterator2) leftmostInorder(root *TreeNode) {
	for root != nil {
		*bsti.stack = append(*bsti.stack, *root)
		root = root.Left
	}
}

func Constructor2(root *TreeNode) BSTIterator2 {
	bi := BSTIterator2{
		stack: &[]TreeNode{},
	}

	bi.leftmostInorder(root)

	return bi
}

func (bsti *BSTIterator2) Next() int {
	len := len(*bsti.stack)
	top := (*bsti.stack)[len-1]
	*bsti.stack = (*bsti.stack)[:len-1]

	if top.Right != nil {
		bsti.leftmostInorder(top.Right)
	}

	return top.Val
}

func (bsti *BSTIterator2) HasNext() bool {
	return len(*bsti.stack) > 0
}

// https://leetcode.com/problems/binary-search-tree-iterator/description/?envType=study-plan-v2&envId=top-interview-150
func bstIterator2Main() {
	arr := []interface{}{7, 3, 15, nil, nil, 9, 20}
	root := buildTreeFromArray(arr, 0)
	bi := Constructor2(root)
	bi.Next()
	bi.HasNext()
}
