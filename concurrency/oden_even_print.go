package main

import (
	"fmt"
	"sync"
)

func main() {
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
