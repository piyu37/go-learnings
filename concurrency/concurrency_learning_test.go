package main

import (
	"bytes"
	"os"
	"sync"
	"testing"
)

// Test updateMessage
func TestUpdateMessage(t *testing.T) {
	// Reset the global msg variable
	msg = ""

	// Call updateMessage to set a new message
	updateMessage("Test message")

	// Check if the message was updated correctly
	if msg != "Test message" {
		t.Errorf("Expected msg to be 'Test message', got '%s'", msg)
	}
}

// Test printMessage
func TestPrintMessage(t *testing.T) {
	// Set the global msg variable
	msg = "Hello, test world!"

	// Create a buffer to capture output
	var buf bytes.Buffer

	// Redirect stdout to the buffer
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	// Run printMessage in a goroutine
	go func() {
		defer w.Close() // Ensure the writer is closed after the message is printed
		printMessage(&wg)
	}()

	// Wait for printMessage to finish
	wg.Wait()

	// Restore the original stdout
	os.Stdout = originalStdout

	// Read the captured output
	buf.ReadFrom(r)
	output := buf.String()

	// Check if the captured output matches the expected message
	expectedOutput := "Hello, test world!\n"
	if output != expectedOutput {
		t.Errorf("Expected output '%s', got '%s'", expectedOutput, output)
	}
}

// Test concurrencyLearning
func TestConcurrencyLearning(t *testing.T) {
	// Run concurrencyLearning and ensure no race conditions or errors
	concurrencyLearning()

	// At this point, the global msg variable would contain the last updated value
	expectedMessages := []string{"Hello, universe!", "Hello, cosmos!", "Hello, world!"}
	found := false

	for _, expected := range expectedMessages {
		if msg == expected {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Final msg value '%s' is not one of the expected messages: %v", msg, expectedMessages)
	}
}
