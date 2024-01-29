package main

// find mid
// print mid
// call left(0, mid)
// call right(mid+1, len)

func createMinHeightBST(arr []int) *BST {
	start := 0
	end := len(arr)
	mid := (start + (end - start)) / 2

	bst := &BST{
		Value: arr[mid],
	}

	if start < mid {
		bst.Left = &BST{}
		insert(bst.Left, arr, start, mid)
	}
	if mid+1 < end {
		bst.Right = &BST{}
		insert(bst.Right, arr, mid+1, end)
	}

	return bst
}

func insert(bst *BST, arr []int, start, end int) {
	mid := start + ((end - start) / 2)
	bst.Value = arr[mid]

	if start < mid {
		bst.Left = &BST{}
		insert(bst.Left, arr, start, mid)
	}
	if mid+1 < end {
		bst.Right = &BST{}
		insert(bst.Right, arr, mid+1, end)
	}
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/11_min_height_bst
func minHeightBST() {
	arr := []int{1, 2, 5, 7, 10, 13, 14, 15, 22}

	createMinHeightBST(arr)
}
