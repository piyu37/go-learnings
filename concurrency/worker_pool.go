package main

import (
	"fmt"
	"time"
)

// numOfJobs = 5
// create jobChannel & resultChannel buffered
// create 3 worker pool by running 3 go routines + add buffer channel in which we need to send the info
// send values from numOfJobs to the channel
// once we receive value from job channel, send it to result channel
// once we send all values, close channel
// read all result channel values

func workersPool(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func worker_pool() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go workersPool(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
