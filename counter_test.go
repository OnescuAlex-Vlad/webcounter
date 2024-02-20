package main

import (
	"testing"
	
)

func TestIncrementValue(t *testing.T) {
	counter := NewCounter()

	// Test case 1: Incrementing the value by 1
	counter.IncrementCounter(1)
	if counter.value != 1 {
		t.Errorf("Expected value to be 1, got %d", counter.value)
	}

	// Test case 2: Incrementing the value by 5
	counter.IncrementCounter(5)
	if counter.value != 6 {
		t.Errorf("Expected value to be 6, got %d", counter.value)
	}

	// Test case 3: Incrementing the value by 0
	counter.IncrementCounter(0)
	if counter.value != 6 {
		t.Errorf("Expected value to be 6, got %d", counter.value)
	}
}

func TestResetValue(t *testing.T) {
	counter := NewCounter()

	// Test case 1: Resetting the value to 0
	counter.IncrementCounter(1)
	counter.ResetValue()
	if counter.value != 0 {
		t.Errorf("Expected value to be 0, got %d", counter.value)
	}

}