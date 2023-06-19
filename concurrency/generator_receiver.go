package main

import "fmt"

// Send and receive integers to a channel using go routines.
// In the main function call a function "generator" which will add 10 integers to a channel; return the channel from the functions.
// In the main function call a function "receiver"; pass the channel to this function and print all the 10 integers.

func generator(ch chan int) {
	for i := 0; i <= 10; i++ {
		ch <- i
	}

	close(ch)
}

func receiver(ch chan int, done chan bool) {
	for i := range ch {
		fmt.Println(i)
	}

	done <- true
}

func main() {
	ch := make(chan int)

	done := make(chan bool)

	go generator(ch)
	go receiver(ch, done)

	<-done

	close(done)
}
