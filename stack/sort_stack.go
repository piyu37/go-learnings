package main

import "fmt"

func SortStack(stack []int) []int {
	if len(stack) == 0 {
		return stack
	}

	top := stack[len(stack)-1]

	stack = stack[:len(stack)-1]

	SortStack(stack)

	insert(&stack, top)

	return stack
}

func insert(stack *[]int, top int) {
	if len(*stack) == 0 || top >= (*stack)[len(*stack)-1] {
		*stack = append(*stack, top)

		return
	}

	topTemp := (*stack)[len(*stack)-1]

	*stack = (*stack)[:len(*stack)-1]

	insert(stack, top)

	*stack = append(*stack, topTemp)
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/52_sort_stack
func sortStackMain() {
	arr := []int{-5, 2, -2, 4, 3, 1}
	fmt.Println(SortStack(arr))
}
