package main

import "fmt"

func concurrency() {
	msg := make(chan string, 2)

	msg <- "Good"
	msg <- "Morning"

	fmt.Println(<-msg)
	fmt.Println(<-msg)
}
