package main

import "fmt"

type LinkedListInfo struct {
	outerNodesEqual   bool
	leftNodeToCompare *LinkedList
}

func LinkedListPalindromeRec(head *LinkedList) bool {
	isPalindromeResult := isPalindrome(head, head)

	return isPalindromeResult.outerNodesEqual
}

func isPalindrome(left, right *LinkedList) LinkedListInfo {
	if right == nil {
		return LinkedListInfo{true, left}
	}

	recCallResult := isPalindrome(left, right.Next)
	leftNodeToCompare := recCallResult.leftNodeToCompare
	outerNodeEqual := recCallResult.outerNodesEqual

	recuIsEqual := outerNodeEqual && leftNodeToCompare.Value == right.Value
	nextLeftToCompare := leftNodeToCompare.Next

	return LinkedListInfo{recuIsEqual, nextLeftToCompare}
}

func llPalindromeRecursion() {
	ll := &LinkedList{
		Value: 0,
		Next: &LinkedList{
			Value: 1,
			Next: &LinkedList{
				Value: 2,
				Next: &LinkedList{
					Value: 2,
					Next: &LinkedList{
						Value: 1,
						Next: &LinkedList{
							Value: 0,
						},
					},
				},
			},
		},
	}

	fmt.Println(LinkedListPalindromeRec(ll))
}
