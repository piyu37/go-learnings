package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type StockTask struct {
	Symbol string
	Status string
}

func fetchPrice(id int, symbol string, done chan struct{}) (float64, error) {
	fmt.Printf("Worker %d: Fetching price for %s...\n", id, symbol)
	time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond) // Simulated fetch
	price := rand.Float64() * 1000
	fmt.Printf("Worker %d: Processed %s at price $%.2f\n", id, symbol, price)
	close(done) // Signal that the fetch is complete
	return price, nil
}

func stockWorker(id int, tasks <-chan StockTask, results chan<- StockTask, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

		done := make(chan struct{})
		var status string

		// Use fetchPrice for fetching the stock price
		go func() {
			_, err := fetchPrice(id, task.Symbol, done)
			if err == nil {
				status = "completed"
			} else {
				status = "failed"
			}
		}()

		select {
		case <-done:
			task.Status = status
		case <-ctx.Done():
			fmt.Printf("Worker %d: Timeout for %s\n", id, task.Symbol)
			task.Status = "timed out"
		}

		cancel() // Explicitly cancel the context at the end of each iteration
		results <- task
	}
}

func rateLimiter(rateLimit int, stop <-chan struct{}) <-chan struct{} {
	tokens := make(chan struct{}, rateLimit)
	go func() {
		ticker := time.NewTicker(1 * time.Second / time.Duration(rateLimit))
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				// select {
				// case tokens <- struct{}{}:
				// default:
				// }
			case <-stop:
				close(tokens)
				return
			}
		}
	}()
	return tokens
}

// Question: Real-Time Stock Price Processor
// Design a Real-Time Stock Price Processor in Go. This system should:

// Fetch stock prices from a simulated API at a fixed interval for a set of stock symbols.
// Process stock prices concurrently with a worker pool, limiting the number of concurrent processing tasks.
// Rate limit API requests so that a maximum of N requests per second are made to the API, with any additional requests being queued.
// Timeout Control: Each stock price fetch and processing task should have a timeout. If a task takes longer than the allotted time, it should be canceled and recorded as "failed."
// Graceful Shutdown: On receiving a shutdown signal, the system should complete all currently executing tasks, discard any pending tasks, and log any failed or completed stock symbol updates.

// Requirements
// Fetch Interval: The stock prices should be fetched every 5 seconds.
// Rate Limiting: Limit API requests to 10 requests per second.
// Timeout: Each task (fetch and process) should have a maximum timeout of 2 seconds.
// Worker Pool Concurrency: Use a pool of 5 workers to process stock prices concurrently.
// Task Queue: Use a buffered channel to queue tasks if the processing limit is reached.
// Graceful Shutdown: On receiving a shutdown signal, allow all ongoing tasks to complete and log the status of each stock symbol as either "completed," "timed out," or "discarded."

// Constraints
// Use time.Ticker to manage the 5-second fetch interval.
// Use a context.Context with a timeout to control each task's execution time.
// Use a buffered channel to queue API requests and prevent excessive requests beyond the limit.
// Implement the worker pool to manage concurrent task processing.
// Ensure safe access to any shared resources using appropriate synchronization.

// Example Workflow
// Every 5 seconds, the system fetches the price for a list of stock symbols (e.g., ["AAPL", "GOOGL", "AMZN"]).
// API requests are limited to 10 requests per second. If more symbols need to be fetched, they are queued until the rate limit allows more requests.
// Each price fetch task has a timeout of 2 seconds. If a task exceeds this time, it is canceled and marked as "failed."
// The fetched stock prices are processed by a pool of 5 workers, each processing one symbol at a time.
// On shutdown, the system completes all current tasks, logs their status, and discards any remaining tasks in the queue.
func stocks() {
	// Configurations
	stockSymbols := []string{"AAPL", "GOOGL", "AMZN", "MSFT", "TSLA"}
	interval := 5 * time.Second
	rateLimit := 10
	workerCount := 5

	// Channels
	taskQueue := make(chan StockTask, len(stockSymbols))
	results := make(chan StockTask, len(stockSymbols))
	stopRateLimiter := make(chan struct{})
	rateTokens := rateLimiter(rateLimit, stopRateLimiter)

	// WaitGroup to manage workers
	var wg sync.WaitGroup

	// Start worker pool
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go stockWorker(i, taskQueue, results, &wg)
	}

	// Ticker for fetching stock prices periodically
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Timer for 30-second timeout
	timeout := time.After(30 * time.Second)

	// Handle CTRL+C for manual shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to enqueue tasks
	go func() {
		for {
			select {
			case <-ticker.C:
				for _, symbol := range stockSymbols {
					select {
					case <-timeout: // Stop after timeout
						fmt.Println("\nTimeout reached! Stopping task queue.")
						close(taskQueue) // Stop accepting new tasks
						return
					case <-sig: // Stop on CTRL+C
						fmt.Println("\nInterrupt received! Stopping task queue.")
						close(taskQueue) // Stop accepting new tasks
						return
					case <-rateTokens: // Consume rate limiter tokens
						taskQueue <- StockTask{Symbol: symbol, Status: "pending"}
					}
				}
			case <-timeout: // Stop after timeout
				fmt.Println("\nTimeout reached! Stopping task queue.")
				close(taskQueue) // Stop accepting new tasks
				return
			case <-sig: // Stop on CTRL+C
				fmt.Println("\nInterrupt received! Stopping task queue.")
				close(taskQueue) // Stop accepting new tasks
				return
			}
		}
	}()

	// Wait for workers to finish processing tasks
	go func() {
		wg.Wait()
		close(results)
	}()

	// Log final statuses
	fmt.Println("\nFinal Task Summary:")
	for result := range results {
		fmt.Printf("%s: %s\n", result.Symbol, result.Status)
	}

	// Stop the rate limiter
	close(stopRateLimiter)

	fmt.Println("System shutdown complete.")
}
