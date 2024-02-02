package main

import "fmt"

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
}

func consumer1(ch <-chan int, done chan bool) {
	for i := range ch {
		fmt.Println(i)
	}

	done <- true
}

func generatorReceiverWorkerPool() {
	ch := make(chan int)

	done := make(chan bool)

	go producer(ch)

	for i := 0; i < 2; i++ {
		go consumer1(ch, done)
	}

	for i := 0; i < 2; i++ {
		<-done
	}

	close(done)
}
