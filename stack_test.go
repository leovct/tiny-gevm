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
