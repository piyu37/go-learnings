package main

type BSTIterator struct {
	inorderArr []int
}

func createInorderList(root *TreeNode) []int {
	arr := []int{}
	if root == nil {
		return arr
	}

	arrLeft := createInorderList(root.Left)
	arr = append(arr, arrLeft...)
	arr = append(arr, root.Val)
	arrRight := createInorderList(root.Right)
	arr = append(arr, arrRight...)

	return arr
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		inorderArr: createInorderList(root),
	}
}

func (bsti *BSTIterator) Next() int {
	if len(bsti.inorderArr) == 0 {
		return -1
	}

	val := bsti.inorderArr[0]

	bsti.inorderArr = bsti.inorderArr[1:]

	return val
}

func (bsti *BSTIterator) HasNext() bool {
	return len(bsti.inorderArr) > 0
}

// // https://leetcode.com/problems/binary-search-tree-iterator/description/?envType=study-plan-v2&envId=top-interview-150
func bstIteratorMain() {
	arr := []interface{}{7, 3, 15, nil, nil, 9, 20}
	root := buildTreeFromArray(arr, 0)
	bi := Constructor(root)
	bi.Next()
	bi.HasNext()
}
