package main

import (
	"fmt"
	"math"
)

func bstPreorderInsertion(arr []int) *BST {
	bst := &BST{
		Value: arr[0],
	}

	insertInBST(1, arr, math.MinInt, math.MaxInt, bst)

	return bst
}

func insertInBST(i int, arr []int, min, max int, bst *BST) int {
	if i == len(arr) || arr[i] < min || arr[i] > max {
		return i
	}

	if arr[i] < bst.Value {
		bst.Left = &BST{
			Value: arr[i],
		}

		i = insertInBST(i+1, arr, min, bst.Value-1, bst.Left)
	}

	if i == len(arr) || arr[i] < min || arr[i] > max {
		return i
	}

	bst.Right = &BST{
		Value: arr[i],
	}

	return insertInBST(i+1, arr, bst.Value, max, bst.Right)
}

func bstPreorderInsertMain() {
	arr := []int{10, 4, 2, 1, 5, 17, 19, 18}

	fmt.Println(bstPreorderInsertion(arr))

}
