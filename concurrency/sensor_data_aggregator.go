package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type sensorDetails struct {
	reading int
	id      int
}

func generateTempReading(ctx context.Context, sensorID int, ch chan<- sensorDetails, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		time.Sleep(time.Duration(rand.IntN(400)+100) * time.Millisecond)
		val := rand.IntN(50)

		select {
		case <-ctx.Done():
			fmt.Println("Shutting down sensor data aggregator: ", sensorID)
			return
		case ch <- sensorDetails{val, sensorID}:
			fmt.Printf("temp. %d is written from sensor %d\n", val, sensorID)
		}
	}
}

// The Sensor Data Aggregator (Graceful Shutdown)

// The Scenario: You have multiple IoT sensors that generate temperature readings at random intervals.
// They all send their data to a central aggregator. The system should collect data for exactly 3 seconds and then cleanly shut down.

// The Prompt for the Candidate:

// Write a Go program that simulates 3 sensors (as goroutines). Each sensor should generate a random integer and send it to a
// shared channel at random intervals (between 100ms and 500ms).
func sensorDataAggregator() {
	runningDuration := 3 * time.Second
	noOfSensors := 3

	ch := make(chan sensorDetails, 10)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), runningDuration)
	defer cancel()

	for i := range noOfSensors {
		wg.Add(1)
		go generateTempReading(ctx, i, ch, &wg)
	}
	// monitor go routine pattern: decouples "waiting for workers" from "reading data."
	go func() { //  "waiting for workers"
		wg.Wait()
		close(ch)
	}()

	for i := range ch { // "reading data"
		fmt.Printf(">> Read: temp %d from sensor %d\n", i.reading, i.id)
	}

	fmt.Println("Program finished successfully.")
}
