package main

import (
	"fmt"
	"sync"
)

var msg string
var lock sync.RWMutex

func updateMessage(s string) {
	lock.Lock()
	msg = s
	lock.Unlock()
}

func printMessage(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(msg)
}

// challenge: modify this code so that the calls to updateMessage() on lines
// 28, 30, and 33 run as goroutines, and implement wait groups so that
// the program runs properly, and prints out three different messages.
// Then, write a test for all three functions in this program: updateMessage(),
// printMessage(), and main().
func concurrencyLearning() {
	msg = "Hello, world!"
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		updateMessage("Hello, universe!")
		printMessage(&wg)
	}()

	wg.Add(1)
	go func() {
		updateMessage("Hello, cosmos!")
		printMessage(&wg)
	}()

	wg.Add(1)
	go func() {
		updateMessage("Hello, world!")
		printMessage(&wg)
	}()

	wg.Wait()
}
