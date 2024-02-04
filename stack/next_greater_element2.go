package main

import "fmt"

// Stack type using a slice
type Stack []int

// Create a new stack
func createStack() Stack {
	return make(Stack, 0)
}

// Check if the stack is empty
func (s Stack) isEmpty() bool {
	return len(s) == 0
}

// Push an element onto the stack
func (s *Stack) push(x int) {
	*s = append(*s, x)
}

// Pop an element from the stack
func (s *Stack) pop() int {
	if s.isEmpty() {
		fmt.Println("Error: stack underflow")
		return -1 // Return a sentinel value to indicate an error
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

// Function to print the next greater element for each element in arr
func printNGE(arr []int) {
	stack := createStack()
	element := 0
	next := 0

	// Push the first element to the stack
	stack.push(arr[0])

	output := make([]int, 0)

	// Iterate over the rest of the elements
	for i := 1; i < len(arr); i++ {
		next = arr[i]

		if !stack.isEmpty() {
			// If the stack is not empty, pop an element from the stack
			element = stack.pop()

			// If the popped element is smaller than the next, print the pair
			for element < next {
				output = append(output, next)
				fmt.Printf("%d -- %d\n", element, next)
				if stack.isEmpty() {
					break
				}
				element = stack.pop()
			}

			// If the element is greater than the next, push it back
			if element > next {
				stack.push(element)
			}
		}

		// Push the next element to the stack so we can find the next greater for it
		stack.push(next)
	}

	// After the loop, the remaining elements in the stack do not have a next greater element
	for !stack.isEmpty() {
		element = stack.pop()
		next = -1
		output = append(output, next)
		fmt.Printf("%d -- %d\n", element, next)
	}

	fmt.Println(output)
}

func nextGreaterElement2() {
	arr := []int{4, 1, 2, 6}
	printNGE(arr)
}
