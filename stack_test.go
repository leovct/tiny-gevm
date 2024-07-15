package main

import (
	"bytes"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	if s == nil {
		t.Error("NewStack() returned nil")
	} else {
		if len(s.data) != 0 {
			t.Errorf("NewStack() created a stack with %d elements, want 0", len(s.data))
		}
	}
}

func TestPush(t *testing.T) {
	// Create an empty stack.
	s := NewStack()

	// Push an element to the stack.
	data := [32]byte{0x01}
	err := s.Push(data)
	if err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}
	if len(s.data) != 1 {
		t.Errorf("Push() resulted in a stack with %d elements, want 1", len(s.data))
	}
	if !bytes.Equal(s.data[0][:], data[:]) {
		t.Errorf("Push() stored %v, want %v", s.data[0], data)
	}
}

func TestPushFull(t *testing.T) {
	// Create an empty stack.
	s := NewStack()

	// Push 1024 elements to the stack.
	data := [32]byte{0x01}
	for i := 0; i < STACK_MAX_SIZE; i++ {
		err := s.Push(data)
		if err != nil {
			t.Errorf("Push() returned an unexpected error on iteration %d: %v", i, err)
		}
	}

	// Try to push another element to the stack. It should fail.
	err := s.Push(data)
	if err != ErrStackFull {
		t.Errorf("Push() on full stack returned %v, want %v", err, ErrStackFull)
	}
}

func TestPop(t *testing.T) {
	// Create an empty stack.
	s := NewStack()

	// Push 3 elements to the stack.
	data1 := [32]byte{0x01}
	data2 := [32]byte{0x02}
	data3 := [32]byte{0x03}
	_ = s.Push(data1)
	_ = s.Push(data2)
	_ = s.Push(data3)

	// Pop an element from the stack.
	value, err := s.Pop()
	if err != nil {
		t.Errorf("Pop() returned an unexpected error: %v", err)
	}
	if value != data3 {
		t.Errorf("Pop() returned %v, want %v", value, [32]byte{0x03})
	}
	if len(s.data) != 2 {
		t.Errorf("Pop() resulted in a stack with %d elements, want 2", len(s.data))
	}
	if !bytes.Equal(s.data[0][:], data1[:]) {
		t.Errorf("Pop() stored %v, want %v", s.data[0], data1)
	}
	if !bytes.Equal(s.data[1][:], data2[:]) {
		t.Errorf("Pop() stored %v, want %v", s.data[0], data2)
	}
}

func TestPopEmpty(t *testing.T) {
	// Create an empty stack.
	s := NewStack()

	// Try to pop an element from the empty stack. It should fail.
	value, err := s.Pop()
	if err != ErrStackEmpty {
		t.Errorf("Pop() on full stack returned %v, want %v", err, ErrStackEmpty)
	}
	if value != [32]byte{} {
		t.Errorf("Pop() returned %v, want %v", value, [32]byte{})
	}
}
