package main

import "fmt"

type BSTWithLeftSubtreeSize struct {
	value, leftSubTreeSize int
	left, right            *BSTWithLeftSubtreeSize
}

func rightSmallerThan(arr []int) []int {
	rightSmallerThanCounts := make([]int, len(arr))

	bst := &BSTWithLeftSubtreeSize{
		value: arr[len(arr)-1],
	}

	for i := len(arr) - 2; i >= 0; i-- {
		bst.insert(arr[i], i, rightSmallerThanCounts, 0)
	}

	return rightSmallerThanCounts
}

func (bst *BSTWithLeftSubtreeSize) insert(value, i int, rightSmallerThanCounts []int, numSmallerAtInsertTime int) {
	if value < bst.value {
		bst.leftSubTreeSize++
		if bst.left == nil {
			bst.left = &BSTWithLeftSubtreeSize{
				value: value,
			}
			rightSmallerThanCounts[i] = numSmallerAtInsertTime
		} else {
			bst.left.insert(value, i, rightSmallerThanCounts, numSmallerAtInsertTime)
		}

		return
	}

	numSmallerAtInsertTime += bst.leftSubTreeSize

	if value > bst.value {
		numSmallerAtInsertTime++
	}

	if bst.right == nil {
		bst.right = &BSTWithLeftSubtreeSize{
			value: value,
		}
		rightSmallerThanCounts[i] = numSmallerAtInsertTime
	} else {
		bst.right.insert(value, i, rightSmallerThanCounts, numSmallerAtInsertTime)
	}
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/06_right_smaller_than
func rightSmallerThanMain() {
	arr := []int{8, 5, 11, -1, 3, 4, 2}
	fmt.Println(rightSmallerThan(arr))
}
