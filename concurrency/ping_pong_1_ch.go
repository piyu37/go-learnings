package main

import (
	"fmt"
	"time"
)

func ping(ch1 chan int, done chan bool) {
	ch1 <- 0
	for {
		select {
		case <-done:
			return
		default:
			<-ch1
			fmt.Println("p1")
			ch1 <- 0
		}
	}
}

func pong(ch1 chan int, done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			<-ch1
			fmt.Println("p2")
			ch1 <- 0
		}
	}
}

func pingPongWith1Channel() {
	ch := make(chan int)
	done := make(chan bool)

	go ping(ch, done)
	go pong(ch, done)

	<-time.After(1 * time.Second)

	done <- true

	close(ch)
}
