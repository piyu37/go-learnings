package main

import "fmt"

// type LinkedList struct {
// 	Value int
// 	Next  *LinkedList
// }

func RearrangeLinkedList(head *LinkedList, k int) *LinkedList {
	node := head
	for node != nil && node.Value < k && node.Next != nil && node.Next.Value < k {
		node = node.Next
	}

	updateHead := false
	if node == head && node.Value >= k {
		updateHead = true
	}

	ptr := node
	for node != nil && node.Next != nil {
		if node.Next.Value < k {
			if updateHead {
				temp := node.Next
				node.Next = node.Next.Next
				temp2 := head
				head = temp
				head.Next = temp2
				ptr = head
				updateHead = false
			} else {
				temp := node.Next
				node.Next = node.Next.Next
				temp2 := ptr.Next
				ptr.Next = temp
				temp.Next = temp2
				ptr = ptr.Next
			}
		} else {
			node = node.Next
		}
	}

	node = head
	for node != nil && node.Next != nil {
		if node.Value == k && node == head {
			node = node.Next
		} else if node.Next.Value == k {
			if node == ptr {
				ptr = ptr.Next
				node = node.Next
			} else if !updateHead {
				temp := node.Next
				node.Next = node.Next.Next
				temp2 := ptr.Next
				ptr.Next = temp
				temp.Next = temp2
				ptr = ptr.Next
			} else {
				temp := node.Next
				node.Next = node.Next.Next
				temp2 := head
				head = temp
				head.Next = temp2
				ptr = head
				updateHead = false
			}
		} else {
			node = node.Next
		}
	}

	return head
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/25_rearrange_linked_list
func rearrangeLL() {
	ll := &LinkedList{
		Value: 0,
		Next: &LinkedList{
			Value: 3,
			Next: &LinkedList{
				Value: 2,
				Next: &LinkedList{
					Value: 1,
					Next: &LinkedList{
						Value: 4,
						Next: &LinkedList{
							Value: 5,
							Next: &LinkedList{
								Value: 3,
								Next: &LinkedList{
									Value: -1,
									Next: &LinkedList{
										Value: -2,
										Next: &LinkedList{
											Value: 3,
											Next: &LinkedList{
												Value: 6,
												Next: &LinkedList{
													Value: 7,
													Next: &LinkedList{
														Value: 3,
														Next: &LinkedList{
															Value: 2,
															Next: &LinkedList{
																Value: -9000,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(RearrangeLinkedList(ll, 3))
}
