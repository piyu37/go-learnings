package main

import (
	"fmt"
	"sync"
	"time"
)

func consumeEvenNumber(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		v, ok := <-ch
		if !ok {
			return
		}

		fmt.Println(v)
		time.Sleep(5 * time.Second)
	}
}

func consumeOddNumber(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		v, ok := <-ch
		if !ok {
			return
		}

		fmt.Println(v)
	}
}

func producerOddEven(n, x, y int, chEven, chOdd chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(chEven)
	defer close(chOdd)

	closeTime := time.After(time.Duration(y) * time.Second)

	for i := 0; i <= n; i++ {
		select {
		case <-closeTime:
			fmt.Println("getting closed by y")
			return
		default:
			if i%2 == 0 {
				chEven <- i
			} else {
				chOdd <- i
			}
			time.Sleep(time.Duration(x) * time.Millisecond)
		}
	}

	fmt.Println("getting closed by n")
}

// - Simulate publisher subscriber model in go.
// - producer will produce integers from 0 to n with a gap of x ms.
// - 2 consumer should consume the message asynchronously.
// - consumer 1 to consume and print only odd numbers
// - consumer 2 to consume and print only even numbers.
// - output should be in sync meaning 1 2 3 4 5 ..
// - exit constraints (whichever is early):
//   - either consume all messages or
//   - time “y” seconds.
//
// - simulation should exit once all producers and consumers are closed gracefully.
func extendedOddEven() {
	chEven, chOdd := make(chan int), make(chan int)

	var wg sync.WaitGroup
	n := 50
	x := 1000
	y := 5

	wg.Add(1)
	go producerOddEven(n, x, y, chEven, chOdd, &wg)

	wg.Add(2)
	go consumeEvenNumber(chEven, &wg)
	go consumeOddNumber(chOdd, &wg)

	wg.Wait()
}
