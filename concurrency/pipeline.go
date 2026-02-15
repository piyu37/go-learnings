package main

import (
	"fmt"
	"sync"
)

// generate - convertes a list of integers to a channel
func generate(done <-chan struct{}, nums ...int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, num := range nums {
			select {
			case ch <- num:
			case <-done:
				return
			}
		}
	}()

	return ch
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(done <-chan struct{}, inCh <-chan int) (<-chan int, <-chan int) {
	ch := make(chan int)
	ch2 := make(chan int)
	go func() {
		defer close(ch)
		for i := range inCh {
			select {
			case ch <- i * i:
			case <-done:
				return
			}
		}
	}()

	go func() {
		defer close(ch2)
		for i := range inCh {
			select {
			case ch2 <- i * i:
			case <-done:
				return
			}
		}
	}()

	return ch, ch2
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(ch <-chan int) {
		defer wg.Done()
		for i := range ch {
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// TODO: Build a Pipeline
// generator() -> square() -> print
func pipeline() {
	done := make(chan struct{})
	genCh := generate(done, 2, 3, 4, 5, 6)
	out1, out2 := square(done, genCh)
	resultCh := merge(done, out1, out2)

	for i := range resultCh {
		fmt.Println(i)
	}

	// fmt.Println(<-resultCh)
	close(done)

}
