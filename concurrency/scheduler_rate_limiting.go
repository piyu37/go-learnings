package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func schedule(fn chan func(), wg *sync.WaitGroup) {
	for i := range fn {
		i()
	}

	wg.Done()
}

// Question: Concurrent Task Scheduler with Rate Limiting
// Write a Go program to implement a concurrent task scheduler with the following requirements:

// The program has a list of tasks represented as functions that print their ID and sleep for a random duration (between 1 and 3 seconds).
// Each task should execute concurrently, but the scheduler must limit the number of concurrent tasks to N (a configurable value).
// The scheduler should ensure that no task is lost and all tasks are executed.
// Once all tasks are completed, the program should print "All tasks completed."
// Constraints:
// You must use goroutines, channels, and synchronization primitives to solve the problem.
// Avoid using any third-party libraries.
// The solution should demonstrate knowledge of worker pools, rate limiting, and graceful shutdown.
// Example:
// Given 10 tasks and a concurrency limit of 3, the program should:

// Start executing up to 3 tasks concurrently.
// As tasks finish, new tasks should be scheduled until all 10 tasks are executed.
// Ensure no deadlocks or race conditions occur.
func schedulerRateLimiting() {
	workers := 3
	ch := make(chan func(), workers)
	var wg sync.WaitGroup

	input := getInput()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go schedule(ch, &wg)
	}

	go func() {
		for _, fn := range input {
			ch <- fn
		}

		close(ch)
	}()

	wg.Wait()

	fmt.Println("All tasks completed.")

}

func getInput() []func() {
	tasks := []func(){
		func() {
			id := 1
			fmt.Printf("Task %d is starting\n", id)
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			fmt.Printf("Task %d is completed\n", id)
		},
		func() {
			id := 2
			fmt.Printf("Task %d is starting\n", id)
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			fmt.Printf("Task %d is completed\n", id)
		},
		func() {
			id := 3
			fmt.Printf("Task %d is starting\n", id)
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			fmt.Printf("Task %d is completed\n", id)
		},
		func() {
			id := 4
			fmt.Printf("Task %d is starting\n", id)
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			fmt.Printf("Task %d is completed\n", id)
		},
		func() {
			id := 5
			fmt.Printf("Task %d is starting\n", id)
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			fmt.Printf("Task %d is completed\n", id)
		},
		func() {
			id := 6
			fmt.Printf("Task %d is starting\n", id)
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			fmt.Printf("Task %d is completed\n", id)
		},
		func() {
			id := 7
			fmt.Printf("Task %d is starting\n", id)
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			fmt.Printf("Task %d is completed\n", id)
		},
		func() {
			id := 8
			fmt.Printf("Task %d is starting\n", id)
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			fmt.Printf("Task %d is completed\n", id)
		},
		func() {
			id := 9
			fmt.Printf("Task %d is starting\n", id)
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			fmt.Printf("Task %d is completed\n", id)
		},
		func() {
			id := 10
			fmt.Printf("Task %d is starting\n", id)
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			fmt.Printf("Task %d is completed\n", id)
		},
	}

	return tasks
}
