package main

import (
	"errors"

	"github.com/holiman/uint256"
)

// MAX_STACK_SIZE defines the maximum number of elements the stack can hold.
const MAX_STACK_SIZE int = 1024

var (
	// ErrStackUnderflow is returned when trying to pop from an empty stack.
	ErrStackUnderflow = errors.New("stack underflow")
	// ErrStackOverflow is returned when trying to push to a full stack.
	ErrStackOverflow = errors.New("stack overflow")

	// ErrEmptyStack is returned when trying to access an element from an empty stack.
	ErrEmptyStack = errors.New("empty stack")
	// ErrStackIndexOutOfRange is returned when trying to access an element from a stack that doesn't exist.
	ErrStackIndexOutOfRange = errors.New("stack index out of range")
)

// IStack defines the methods that a stack implementation should have.
type IStack interface {
	// Push adds a new element to the top of the stack.
	// It returns an error if the stack is full.
	Push(*uint256.Int) error

	// Pop removes and returns the top element from the stack.
	// If the stack is empty, it returns a zero-value 32-byte array and an error.
	Pop() (*uint256.Int, error)

	// Get returns the i-th element from the stack without poping it.
	// The index is 1-based, where 1 refers to the top of the stack (last element).
	// For example, Get(1) returns the top element, Get(2) returns the second from the top, and so on.
	// If the stack is empty, it returns a zero-value 32-byte array and an error.
	Get(i int) (*uint256.Int, error)

	// Size returns the number of elements currently on the stack.
	Size() int
}

// Stack represents a last-in-first-out (LIFO) stack of 32-byte arrays.
type Stack struct {
	data []uint256.Int
}

// NewStack creates and returns a new, empty Stack instance.
func NewStack() IStack {
	return &Stack{data: make([]uint256.Int, 0)}
}

func (s *Stack) Push(value *uint256.Int) error {
	if len(s.data) >= MAX_STACK_SIZE {
		return ErrStackOverflow
	}
	s.data = append(s.data, *value)
	return nil
}

func (s *Stack) Pop() (*uint256.Int, error) {
	if len(s.data) == 0 {
		return nil, ErrStackUnderflow
	}
	index := len(s.data) - 1
	element := s.data[index]
	s.data = s.data[:index]
	return &element, nil
}

func (s *Stack) Get(i int) (*uint256.Int, error) {
	if len(s.data) == 0 {
		return nil, ErrEmptyStack
	}
	if len(s.data) < i {
		return nil, ErrStackIndexOutOfRange
	}
	index := len(s.data) - i
	element := s.data[index]
	return &element, nil
}

func (s *Stack) Size() int {
	return len(s.data)
}
