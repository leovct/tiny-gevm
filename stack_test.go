package main

import (
	"bytes"
	"testing"
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
	err := s.Push([32]byte{0x1})
	if err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	// The stack should be equal to [0x1, 0x2].
	err = s.Push([32]byte{0x2})
	if err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	// Check the elements by popping them off.
	// The stack should be equal to [0x1].
	popped, err := s.Pop()
	expectedValue := [32]byte{0x2}
	if err != nil || !bytes.Equal(popped[:], expectedValue[:]) {
		t.Errorf("Expected %v, got %v", expectedValue, popped)
	}

	// The stack should be equal to [].
	popped, err = s.Pop()
	expectedValue = [32]byte{0x1}
	if err != nil || !bytes.Equal(popped[:], expectedValue[:]) {
		t.Errorf("Expected %v, got %v", expectedValue, popped)
	}
}

func TestPushFull(t *testing.T) {
	// Create an empty stack.
	s := NewStack()

	// Push 1024 elements to the stack.
	// The stack should contain 1024 0x1 elements.
	data := [32]byte{0x1}
	for i := 0; i < MAX_STACK_SIZE; i++ {
		err := s.Push(data)
		if err != nil {
			t.Errorf("Push() returned an unexpected error on iteration %d: %v", i, err)
		}
	}

	// Try to push another element to the stack. It should fail.
	err := s.Push(data)
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
	if value != [32]byte{} {
		t.Errorf("Pop() returned %v, want %v", value, [32]byte{})
	}
}
