// How to implement a stack in golang? using push and pop functions

package main

import "fmt"

type node struct {
	value int
	next  *node
	size  int
}

func initialize() *node {
	return &node{}
}

func (stack *node) push(value int) {
	if stack.size == 0 {
		stack.value = value
		stack.size++
		return
	}

	n := node{
		value: value,
		size:  stack.size + 1,
	}

	head := *stack
	*stack = n
	stack.next = &head
}

func (stack *node) pop() int {
	if stack.size == 0 {
		fmt.Println([]int{})
		return -1
	}

	stack.size--
	poppedValue := stack.value
	if stack.next == nil {
		fmt.Println([]int{})
	} else {
		*stack = *stack.next
		printValues(stack)
	}

	return poppedValue
}

func printValues(stack *node) {
	curr := stack
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func createStackFromLL() {
	stack := initialize()
	stack.push(5)
	stack.push(4)
	stack.push(3)
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}
