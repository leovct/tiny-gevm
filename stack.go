package main

import "errors"

// MAX_STACK_SIZE defines the maximum number of elements the stack can hold.
const MAX_STACK_SIZE int = 1024

var (
	// ErrStackUnderflow is returned when trying to pop from an empty stack.
	ErrStackUnderflow = errors.New("stack underflow")
	// ErrStackOverflow is returned when trying to push to a full stack.
	ErrStackOverflow = errors.New("stack overflow")
)

// IStack defines the methods that a stack implementation should have.
type IStack interface {
	Push([32]byte) error
	Pop() ([32]byte, error)
}

// Stack represents a last-in-first-out (LIFO) stack of 32-byte arrays.
type Stack struct {
	data [][32]byte
}

// NewStack creates and returns a new, empty Stack.
func NewStack() IStack {
	return &Stack{data: make([][32]byte, 0)}
}

// Push adds a new element to the top of the stack.
// It returns an error if the stack is full.
func (s *Stack) Push(element [32]byte) error {
	if len(s.data) >= MAX_STACK_SIZE {
		return ErrStackOverflow
	}
	s.data = append(s.data, element)
	return nil
}

// Pop removes and returns the top element from the stack.
// If the stack is empty, it returns a zero-value 32-byte array and an error.
func (s *Stack) Pop() ([32]byte, error) {
	if len(s.data) == 0 {
		return [32]byte{}, ErrStackUnderflow
	}
	index := len(s.data) - 1
	element := s.data[index]
	s.data = s.data[:index]
	return element, nil
}
