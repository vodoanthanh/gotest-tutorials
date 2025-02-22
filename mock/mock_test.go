package mock

import (
	"sync"
	"testing"
)

// Test function
func TestDoubleNumber(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	var mu sync.Mutex
	var calledWith int
	var returnValue int

	// Mock double function
	originalDouble := executeFunc
	executeFunc = func(val int) int {
		mu.Lock()
		defer mu.Unlock()
		calledWith = val
		returnValue = val * 2
		wg.Done() // Notify that the mock function was called
		return returnValue
	}
	defer func() { executeFunc = originalDouble }() // Restore original function

	// Call doubleNumber asynchronously
	go doubleNumber(5)

	// Wait for doubleNumber to complete
	wg.Wait()

	// Assertions
	mu.Lock()
	defer mu.Unlock()

	if calledWith != 5 {
		t.Errorf("Expected double to be called with 5, but got %d", calledWith)
	}

	if returnValue != 10 {
		t.Errorf("Expected double to return 10, but got %d", returnValue)
	}
}
