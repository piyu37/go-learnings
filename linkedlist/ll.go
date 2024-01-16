package main

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

type linkedList struct {
	head *node
	len  int
}

func (l *linkedList) insert(v int) {
	n := node{}
	n.val = v

	if l.head == nil {
		l.head = &n
		l.len++
		return
	}

	ptr := l.head

	for ptr.next != nil {
		ptr = ptr.next
	}

	ptr.next = &n
	l.len++
}

func (l *linkedList) traverse() {
	if l.head == nil {
		return
	}

	ptr := l.head

	for ptr != nil {
		fmt.Println(ptr.val)
		ptr = ptr.next
	}
}

func (l *linkedList) reverseTraverse() {
	if l.head == nil {
		return
	}

	var prv *node
	curr := l.head

	for curr != nil {
		tmp := curr.next
		curr.next = prv
		prv = curr
		curr = tmp
	}

	l.head = prv
	l.traverse()
}

func ll() {
	l := linkedList{}
	l.insert(1)
	l.insert(2)
	l.insert(3)

	l.traverse()
	l.reverseTraverse()
}
