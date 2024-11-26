package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Task is a type that represents a task function with a unique ID
type Task struct {
	ID   int
	Func func(ctx context.Context)
}

type TaskManager struct {
	tasks       map[int]Task
	mu          sync.Mutex
	interval    time.Duration
	stopChannel chan struct{}
}

// NewTaskManager initializes a new TaskManager
func NewTaskManager(interval time.Duration) *TaskManager {
	return &TaskManager{
		tasks:       make(map[int]Task),
		interval:    interval,
		stopChannel: make(chan struct{}),
	}
}

// AddTask allows adding a new task to the manager
func (tm *TaskManager) AddTask(task Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.tasks[task.ID] = task
}

// RemoveTask allows removing a task from the manager by ID
func (tm *TaskManager) RemoveTask(taskID int) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	delete(tm.tasks, taskID)
}

func (tm *TaskManager) Start() {
	ticker := time.NewTicker(tm.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tm.runTasks()
		case <-tm.stopChannel:
			fmt.Println("Shutting down TaskManager...")
			return
		}
	}
}

// Stop gracefully stops the TaskManager
func (tm *TaskManager) Stop() {
	close(tm.stopChannel)
}

func (tm *TaskManager) runTasks() {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for _, t := range tm.tasks {
		go func(t Task) {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			done := make(chan struct{})

			go func() {
				t.Func(ctx)
				close(done)
			}()

			select {
			case <-ctx.Done():
				fmt.Printf("Task %d timed out.\n", t.ID)
			case <-done:
				fmt.Printf("Task %d completed successfully.\n", t.ID)
			}
		}(t)
	}
}

// Question: Scheduled Task Manager with Dynamic Cancellation and Timeout
// Design and implement a Scheduled Task Manager in Go. This manager will:

// Periodically execute a set of scheduled tasks.
// Allow for each task to run with a timeout. If a task exceeds its allocated time, it should be canceled and reported as "timed out."
// Provide the ability to dynamically add or remove tasks from the schedule without stopping the manager.
// Allow graceful shutdown of all tasks when the manager is stopped.
// Requirements
// Periodic Execution: Each task should execute at a fixed interval (e.g., every 5 seconds).
// Timeout Control: Each task should have a 2-second execution limit. If a task exceeds this time, it should be canceled, and an error should be logged indicating it was "timed out."
// Dynamic Task Management: The manager should be able to add new tasks or remove existing tasks dynamically during runtime.
// Graceful Shutdown: The manager should provide a way to gracefully shut down, allowing all tasks to complete their current execution cycle before exiting.
// Constraints
// Use time.Ticker to schedule tasks at fixed intervals.
// Use context.Context for controlling task timeouts and for dynamically canceling tasks.
// Avoid using global variables. The solution should be designed for modularity and reusability.
func taskManager() {
	// Create a TaskManager that runs tasks every 5 seconds
	manager := NewTaskManager(5 * time.Second)

	// Define some example tasks
	task1 := Task{
		ID: 1,
		Func: func(ctx context.Context) {
			time.Sleep(1 * time.Second)
			fmt.Println("Task 1 completed work.")
		},
	}

	task2 := Task{
		ID: 2,
		Func: func(ctx context.Context) {
			time.Sleep(3 * time.Second) // Intentionally longer than the 2-second limit
			fmt.Println("Task 2 completed work.")
		},
	}

	// Add tasks to the manager
	manager.AddTask(task1)
	manager.AddTask(task2)

	// Start the task manager
	go manager.Start()

	// Let the tasks run for 15 seconds
	time.Sleep(15 * time.Second)

	// Stop the task manager
	manager.Stop()
}
