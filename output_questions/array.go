package main

import (
	"container/list"
	"fmt"
)

// When you slice an existing slice using the syntax slice[start:end], the new slice shares the same underlying array. The formula Go uses for the new capacity is:
// New Capacity = Original Capacity - Start Index

func slice1() {
	// 0233
	a := [...]int{0, 1, 2, 3} // 0123 4 4
	x := a[:1]                // 0 1 4
	y := a[2:]                // 23 2 2
	fmt.Println(x, cap(x), y, cap(y))
	x = append(x, y...) // 023 3 4
	x = append(x, y...) // 02333 5 8
	fmt.Println(a, x)
}

// 91, 62, 25, 88
func containerTypes() {
	l := list.New()
	e4 := l.PushBack(88)
	e1 := l.PushFront(91)
	l.InsertBefore(25, e4)
	l.InsertAfter(62, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
