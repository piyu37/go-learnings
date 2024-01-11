package main

import "fmt"

type MinMaxStack struct {
	stack []int
	entry []entry
}

type entry struct {
	min int
	max int
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{}
}

func (stack *MinMaxStack) Peek() int {
	stackLen := len(stack.stack)

	if stackLen > 0 {
		return stack.stack[stackLen-1]
	}

	return -1
}

func (stack *MinMaxStack) Pop() int {
	stackLen := len(stack.stack)

	if stackLen > 0 {
		last := stack.stack[stackLen-1]
		stack.stack = stack.stack[:stackLen-1]
		stack.entry = stack.entry[:stackLen-1]

		return last
	}

	return -1
}

func (stack *MinMaxStack) Push(number int) {
	if len(stack.stack) == 0 {
		stack.stack = append(stack.stack, number)
		stack.entry = append(stack.entry, entry{number, number})

		return
	}

	lastEntry := stack.entry[len(stack.entry)-1]
	min := lastEntry.min
	if number < min {
		min = number
	}
	max := lastEntry.max
	if number > max {
		max = number
	}

	stack.stack = append(stack.stack, number)
	stack.entry = append(stack.entry, entry{min, max})
}

func (stack *MinMaxStack) GetMin() int {
	if len(stack.entry) > 0 {
		return stack.entry[len(stack.entry)-1].min
	}

	return -1
}

func (stack *MinMaxStack) GetMax() int {
	if len(stack.entry) > 0 {
		return stack.entry[len(stack.entry)-1].max
	}

	return -1
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/27_min_max_stack
func minMaxStackMain() {
	minMaxStack := NewMinMaxStack()

	minMaxStack.Push(5)
	fmt.Println(minMaxStack.GetMin())
	fmt.Println(minMaxStack.GetMax())
	fmt.Println(minMaxStack.Peek())
	minMaxStack.Push(7)
	fmt.Println(minMaxStack.GetMin())
	fmt.Println(minMaxStack.GetMax())
	fmt.Println(minMaxStack.Peek())
	minMaxStack.Push(2)
	fmt.Println(minMaxStack.GetMin())
	fmt.Println(minMaxStack.GetMax())
	fmt.Println(minMaxStack.Pop())
	fmt.Println(minMaxStack.GetMin())
	fmt.Println(minMaxStack.GetMax())
	fmt.Println(minMaxStack.Pop())
	fmt.Println(minMaxStack.GetMin())
	fmt.Println(minMaxStack.GetMax())
}
