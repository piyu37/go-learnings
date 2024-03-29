package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func odd_even_print() {
	syncChannel := make(chan bool) // unbuffered channel.

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		// prints even numbers.
		defer wg.Done()

		for i := 0; i < 50; i += 2 {
			<-syncChannel

			fmt.Printf("%d ", i)

			syncChannel <- true
		}
		// atomic.StoreInt32(&done, 1)
	}()

	syncChannel <- true

	go func() {
		// prints odd numbers.
		defer wg.Done()

		for i := 1; i < 50; i += 2 {
			<-syncChannel

			fmt.Printf("%d ", i)

			// if atomic.LoadInt32(&done) != 0 {
			// 	return
			// }

			if i == 49 {
				continue
			}

			syncChannel <- true
		}
	}()

	wg.Wait()
	close(syncChannel)
}

// using atomic
func odd_even_print2() {
	var done int32
	syncChannel := make(chan bool) // unbuffered channel.

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		// prints even numbers.
		defer wg.Done()

		for i := 0; i < 50; i += 2 {
			<-syncChannel

			fmt.Printf("%d ", i)

			syncChannel <- true
		}
		atomic.StoreInt32(&done, 1)
	}()

	syncChannel <- true

	go func() {
		// prints odd numbers.
		defer wg.Done()

		for i := 1; i < 50; i += 2 {
			<-syncChannel

			fmt.Printf("%d ", i)

			if atomic.LoadInt32(&done) != 0 {
				return
			}

			syncChannel <- true
		}
	}()

	wg.Wait()
	close(syncChannel)
}

func oddEvenPrintMain() {
	odd_even_print()
	odd_even_print2()
}
