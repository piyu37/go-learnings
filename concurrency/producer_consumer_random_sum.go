package main

import (
	"fmt"
	"math/rand"
	"time"
)

type sumCal struct {
	total int
}

func (sc sumCal) producer(id int, ch chan<- int, done <-chan bool) {
	for {
		select {
		case <-done:
			return
		case <-time.After(1 * time.Second):
			num := rand.Intn(100) + 1
			ch <- num
		}
	}
}

func (sc sumCal) consumer(ch <-chan int) {
	for i := range ch {
		sc.total += i
		fmt.Println(sc.total)
	}
}

// Question: Write a Go program that simulates a real-time data processing system.
// The program should have multiple data producers and a single data consumer.
// Each data producer generates a random number and sends it to the data consumer
// through a channel. The data consumer receives the numbers from different producers,
// calculates their sum, and prints the current sum whenever a new number is received.

// Implement a function producer that takes a producer ID and a channel as input.
// The function should generate random numbers between 1 and 100 at regular
// intervals (e.g., every 1 second) and send them to the channel.

// Implement a function consumer that takes a slice of channels as input. The function should
// continuously listen to the channels and calculate the sum of the received numbers from all
// producers. Whenever a new number is received from any producer, the function should print the
// current sum.

// Your task is to write the complete program, including the producer and consumer functions,
// and any other necessary functions or variables.
func producerConsumerRandomSum() {
	fmt.Println("---------producerConsumerRandomSum----------")
	ch := make(chan int)
	done := make(chan bool)
	sc := sumCal{}

	for i := 0; i < 3; i++ {
		go sc.producer(i, ch, done)
	}

	go sc.consumer(ch)

	<-time.After(5 * time.Second)

	for i := 0; i < 3; i++ {
		done <- true
	}

	close(ch)
}
