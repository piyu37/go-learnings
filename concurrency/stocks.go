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

func fetchPrice(ctx context.Context, id int, symbol string) (float64, error) {
	fmt.Printf("Worker %d: Fetching %s...\n", id, symbol)
	delay := time.Duration(rand.Intn(2500)) * time.Millisecond
	select {
	case <-time.After(delay):
		price := rand.Float64() * 1000
		fmt.Printf("Worker %d: Processed %s at $%.2f\n", id, symbol, price)
		return price, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func stockWorker(id int, tasks <-chan StockTask, results chan<- StockTask, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		_, err := fetchPrice(ctx, id, task.Symbol)
		cancel()

		switch {
		case err == context.DeadlineExceeded || err == context.Canceled:
			task.Status = "timed out"
		case err != nil:
			task.Status = "failed"
		default:
			task.Status = "completed"
		}
		results <- task
	}
}

func rateLimiter(ctx context.Context, rateLimit int) <-chan struct{} {
	// Without buffering:
	// If workers are busy for 150 ms, they miss some ticks because the tokens <- struct{}{} will block.
	// That effectively reduces throughput below 10/sec.

	// With buffering:
	// Tokens accumulate (up to rps in the buffer).
	// When workers catch up, they can “burst” up to that buffer size in one go.
	tokens := make(chan struct{}, rateLimit)
	ticker := time.NewTicker(1 * time.Second / time.Duration(rateLimit))
	go func() {
		defer ticker.Stop()
		defer close(tokens)
		for {
			select {
			case <-ticker.C:
				// This is a non-blocking send.
				// Without it: if the channel is full, the goroutine blocks on tokens <- struct{}{} → deadlock or leak if nobody’s consuming.
				// With it: if the buffer is full, we just drop the token (so max burst size = buffer size).
				// hence, if buffer is full, we will just run default case;
				// That prevents the rate-limiter goroutine from ever stalling.
				select {
				case tokens <- struct{}{}:
				default: // buffer full → drop token (keeps goroutine from blocking)
				}
			case <-ctx.Done():
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

	// Root context + signal handling (single broadcast for shutdown)
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle CTRL+C for manual shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Println("\nShutdown signal received — finishing in-flight tasks, discarding queued ones...")
		cancel()
	}()

	// Channels
	// producer: every 5s enqueue all symbols into requestQueue (no token here)
	requestQueue := make(chan StockTask, 32)         // tasks waiting for rate-limit
	taskQueue := make(chan StockTask, workerCount*2) // tasks ready for workers
	results := make(chan StockTask, 256)

	rateTokens := rateLimiter(rootCtx, rateLimit)

	// WaitGroup to manage workers
	var wg sync.WaitGroup

	// Start worker pool
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go stockWorker(i, taskQueue, results, &wg)
	}

	// Wait for workers to finish processing tasks
	go func() {
		wg.Wait()
		close(results)
	}()

	// Ticker for fetching stock prices periodically
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Timer for 30-second timeout
	timeout := time.After(30 * time.Second)

	// Producer: enqueue one task per symbol every 5s (no rate-limit here)
	go func() {
		// Stop accepting new tasks in request queue
		defer close(requestQueue)

		for {
			select {
			case <-ticker.C:
				for _, symbol := range stockSymbols {
					select {
					case <-rootCtx.Done():
						return
					case requestQueue <- StockTask{Symbol: symbol, Status: "pending"}:
					}
				}
			case <-timeout: // Stop after timeout
				fmt.Println("\nRun timer reached — closing request queue (no new tasks).")
				return
			case <-rootCtx.Done():
				return
			}
		}
	}()

	wg.Add(1)
	// Dispatcher: for each token, forward at most one queued task to workers.
	// On shutdown (rootCtx.Done), drain requestQueue as "discarded", then close taskQueue.
	go func() {
		defer wg.Done()
		defer close(taskQueue)
	loop:
		for {
			select {
			case <-rootCtx.Done():
				break loop
			case _, ok := <-rateTokens:
				if !ok {
					break loop
				}

				// Got a token, now try to get a task.
				select {
				case task, ok := <-requestQueue:
					if !ok {
						// Producer is done, no more tasks will ever arrive.
						break loop
					}
					taskQueue <- task
				case <-rootCtx.Done():
					// Shutdown signal received while waiting for a task.
					break loop
				}
			}
		}

		// GUARANTEED CLEANUP: This code now runs AFTER the loop has exited.
		fmt.Println("Dispatcher shutting down. Draining remaining requests...")
		for task := range requestQueue {
			task.Status = "discarded"
			results <- task
		}
	}()

	// Log final statuses
	fmt.Println("\nFinal Task Summary:")
	for result := range results {
		fmt.Printf("%s: %s\n", result.Symbol, result.Status)
	}

	fmt.Println("System shutdown complete.")
}
