package main

import "fmt"

func switchCase() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x is :%T", i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	case func(int):
		fmt.Printf("x is func(int)")
	case bool, string:
		fmt.Printf("x is bool or string")
	default:
		fmt.Printf("Don't Know")
	}
}
