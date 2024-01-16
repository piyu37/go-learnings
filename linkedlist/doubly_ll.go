package main

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) SetHead(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}

	ll.InsertBefore(ll.Head, node)
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	if ll.Tail == nil {
		ll.SetHead(node)
		return
	}

	ll.InsertAfter(ll.Tail, node)
}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	nodeToInsert.Next = node
	nodeToInsert.Prev = node.Prev
	if node != ll.Head {
		node.Prev.Next = nodeToInsert
	} else {
		ll.Head = nodeToInsert
	}

	node.Prev = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	nodeToInsert.Next = node.Next
	nodeToInsert.Prev = node

	if node != ll.Tail {
		node.Next.Prev = nodeToInsert
	} else {
		ll.Tail = nodeToInsert
	}

	node.Next = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	if position == 1 {
		ll.SetHead(nodeToInsert)
		return
	}

	currentNode := ll.Head
	count := 1

	for currentNode != nil && position != count {
		currentNode = currentNode.Next
		count++
	}

	if currentNode == nil {
		ll.SetTail(nodeToInsert)
	} else {
		ll.InsertBefore(currentNode, nodeToInsert)
	}
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	currentNode := ll.Head

	for currentNode != nil {
		if currentNode.Value == value {
			ll.Remove(currentNode)
			break
		}

		currentNode = currentNode.Next
	}
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	if node == ll.Head {
		ll.Head = node.Next
	}

	if node == ll.Tail {
		ll.Tail = node.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Next = nil
	node.Prev = nil
}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	contains := false

	currentNode := ll.Head

	for currentNode != nil {
		if currentNode.Value == value {
			contains = true
			break
		}

		currentNode = currentNode.Next
	}

	return contains
}
