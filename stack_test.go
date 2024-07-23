package main

import (
	"testing"

	"github.com/holiman/uint256"
)

func TestNewStack(t *testing.T) {
	// Create an empty stack.
	s := NewStack()
	if s == nil {
		t.Error("NewStack() returned nil")
	}
}

func TestPushAndPop(t *testing.T) {
	// Create an empty stack.
	s := NewStack()

	// Push two elements to the stack.
	// The stack should be equal to [0x1].
	err := s.Push(uint256.NewInt(1))
	if err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	// The stack should be equal to [0x1, 0x2].
	err = s.Push(uint256.NewInt(2))
	if err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	// Check the elements by popping them off.
	// The stack should be equal to [0x1].
	popped1, err := s.Pop()
	expectedValue := uint256.NewInt(2)
	if err != nil || !popped1.Eq(expectedValue) {
		t.Errorf("Expected %v, got %v", expectedValue, popped1)
	}

	// The stack should be equal to [].
	popped2, err := s.Pop()
	expectedValue = uint256.NewInt(1)
	if err != nil || !popped2.Eq(expectedValue) {
		t.Errorf("Expected %v, got %v", expectedValue, popped2)
	}
}

func TestPushFull(t *testing.T) {
	// Create an empty stack.
	s := NewStack()

	// Push 1024 elements to the stack.
	// The stack should contain 1024 0x1 elements.
	for i := 0; i < MAX_STACK_SIZE; i++ {
		err := s.Push(uint256.NewInt(1))
		if err != nil {
			t.Errorf("Push() returned an unexpected error on iteration %d: %v", i, err)
		}
	}

	// Try to push another element to the stack. It should fail.
	err := s.Push(uint256.NewInt(1))
	if err != ErrStackOverflow {
		t.Errorf("Push() on full stack returned %v, want %v", err, ErrStackOverflow)
	}
}

func TestPopEmpty(t *testing.T) {
	// Create an empty stack.
	s := NewStack()

	// Try to pop an element from the empty stack. It should fail.
	value, err := s.Pop()
	if err != ErrStackUnderflow {
		t.Errorf("Pop() on full stack returned %v, want %v", err, ErrStackUnderflow)
	}
	if value != nil {
		t.Errorf("Pop() returned %v, want %v", value, nil)
	}
}

func TestSwap(t *testing.T) {
	// Create an empty stack.
	s := NewStack()

	// Try to swap two elements of an empty stack.
	if err := s.Swap(1); err == nil {
		t.Errorf("Swap() returned an unexpected error: %v, wanted: %v", err, ErrEmptyStack)
	}

	// Push two elements to the stack.
	// The stack should be equal to [0x1].
	if err := s.Push(uint256.NewInt(1)); err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	// The stack should be equal to [0x1, 0x2].
	if err := s.Push(uint256.NewInt(2)); err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	// Swap the 1st and 2nd elements.
	if err := s.Swap(2); err != nil {
		t.Errorf("Swap() returned an unexpected error: %v, wanted: %v", err, nil)
	}

	// Swap an element that doesn't exist.
	if err := s.Swap(3); err == nil {
		t.Errorf("Swap() returned an unexpected error: %v, wanted: %v", err, ErrStackIndexOutOfRange)
	}
}

func TestGet(t *testing.T) {
	// Create an empty stack.
	s := NewStack()

	// Try to get an element from the empty stack.
	value0, err := s.Get(1)
	if err == nil {
		t.Errorf("Get() returned an unexpected error: %v, wanted: %v", err, ErrEmptyStack)
	}
	if value0 != nil {
		t.Errorf("Get() returned %v, want %v", value0, nil)
	}

	// Push two elements to the stack.
	// The stack should be equal to [0x1].
	if err = s.Push(uint256.NewInt(1)); err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	// The stack should be equal to [0x1, 0x2].
	if err = s.Push(uint256.NewInt(2)); err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	// Get the first element.
	value1, err := s.Get(1)
	if err != nil {
		t.Errorf("Get() returned an unexpected error: %v", err)
	}
	if value1.Uint64() != 2 {
		t.Errorf("Get() returned %v, want %v", value1.Uint64(), 2)
	}

	// Get the second element.
	value2, err := s.Get(2)
	if err != nil {
		t.Errorf("Get() returned an unexpected error: %v", err)
	}
	if value2.Uint64() != 1 {
		t.Errorf("Get() returned %v, want %v", value2.Uint64(), 1)
	}

	// Get an index out of range.
	value3, err := s.Get(3)
	if err == nil {
		t.Errorf("Get() returned an unexpected error: %v, wanted: %v", err, ErrStackIndexOutOfRange)
	}
	if value3 != nil {
		t.Errorf("Get() returned %v, want %v", value3, nil)
	}
}

func TestSize(t *testing.T) {
	// Create an empty stack.
	s := NewStack()

	// Push two elements to the stack.
	// The stack should be equal to [0x1].
	err := s.Push(uint256.NewInt(1))
	if err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	// The stack should be equal to [0x1, 0x2].
	err = s.Push(uint256.NewInt(2))
	if err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	// Check the size of the stack.
	size := s.Size()
	expectedSize := 2
	if size != expectedSize {
		t.Errorf("Expected %v, got %v", expectedSize, size)
	}
}
