package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	// Create a new counter
	nextCount := counter()

	// Use the counter
	fmt.Println(nextCount()) // Outputs: 1
	fmt.Println(nextCount()) // Outputs: 2
	fmt.Println(nextCount()) // Outputs: 3

	// Create another counter, independent of the first
	anotherCount := counter()
	fmt.Println(anotherCount()) // Outputs: 1
}
