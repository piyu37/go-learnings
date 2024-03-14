package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doubleProducer(done <-chan bool) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for {
			select {
			case ch <- rand.Intn(10) * 2:
			case <-done:
				return
			default:
				time.Sleep(2 * time.Second)
			}
		}
	}()

	return ch
}

func doubleConsumer(ch <-chan int) {
	for i := range ch {
		fmt.Printf("double of %d: %d\n", i/2, i)
	}
}

func numberDoubleWorkerPool() {
	workers := 3

	done := make(chan bool)

	chs := make([]<-chan int, workers)

	cCh := make(chan int, workers)

	for i := 0; i < workers; i++ {
		chs[i] = doubleProducer(done)
	}

	go doubleConsumer(cCh)

	// This is the one way to handle multiple producer channels using single consumer
	for i := 0; i < workers; i++ {
		go func(i int) {
			for i := range chs[i] {
				cCh <- i
			}
		}(i)
	}

	// // This is the second way to consume multiple producer channels using multiple workers
	// for i := range chs {
	// 	go doubleConsumer(chs[i])
	// }

	time.Sleep(10 * time.Second)

	close(done)
}
