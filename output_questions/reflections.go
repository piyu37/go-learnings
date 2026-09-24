package main

import "fmt"

func observe(i interface{}) {
	fmt.Printf("The type passed is: %T\n", i)
	fmt.Printf("The value passed is: %#v\n", i)
	fmt.Println()
}

// The type passed is: float64
// The value passed is: 35

// The type passed is: string
// The value passed is: "Washing Powder"
func reflections() {
	var value float64 = 35
	value2 := "Washing Powder"

	observe(value)
	observe(value2)
}
