package main

import "github.com/holiman/uint256"

// IEVM defines the methods that an Ethereum Virtual Machine implementation should have.
type IEVM interface {
	//// Stack operations
	// Push an item to the stack.
	Push(*uint256.Int) error

	// Pop an item from the stack.
	Pop() (*uint256.Int, error)

	//// Math operations
	// Add the top two elements of the stack and push the result back to the stack.
	Add() error

	// Multiply the top two elements of the stack and push the result back to the stack.
	Mul() error

	// Subtract the top two elements of the stack and push the result back to the stack.
	Sub() error

	// Perform the integer divison operation on the top two elements of the stack and push the result back to the stack.
	Div() error

	// Perform the signed integer division operation (trunced) on the top two elements of the stack and push the result back to the stack.
	SDiv() error
}

// EVM represents an Ethereum Virtual Machine.
type EVM struct {
	stack   IStack
	memory  IMemory
	storage IStorage
}

// NewEVM creates and returns a new EVM instance.
func NewEVM() IEVM {
	return &EVM{
		stack:   NewStack(),
		memory:  NewMemory(),
		storage: NewStorage(),
	}
}

func (e *EVM) Push(value *uint256.Int) error {
	return e.stack.Push(value)
}

func (e *EVM) Pop() (*uint256.Int, error) {
	return e.stack.Pop()
}

func (e *EVM) Add() error {
	op := func(a, b *uint256.Int) *uint256.Int {
		return new(uint256.Int).Add(a, b)
	}
	return e.performStackOperation(op)
}

func (e *EVM) Mul() error {
	op := func(a, b *uint256.Int) *uint256.Int {
		return new(uint256.Int).Mul(a, b)
	}
	return e.performStackOperation(op)
}

func (e *EVM) Sub() error {
	op := func(a, b *uint256.Int) *uint256.Int {
		return new(uint256.Int).Sub(a, b)
	}
	return e.performStackOperation(op)
}

func (e *EVM) Div() error {
	op := func(a, b *uint256.Int) *uint256.Int {
		return new(uint256.Int).Div(a, b)
	}
	return e.performStackOperation(op)
}

func (e *EVM) SDiv() error {
	op := func(a, b *uint256.Int) *uint256.Int {
		return new(uint256.Int).SDiv(a, b)
	}
	return e.performStackOperation(op)
}

// Perform a binary operation on the top two elements on the stack.
// It pops two values from the stack, applies the operation, and pushes the result back to the stack.
func (e *EVM) performStackOperation(op func(a, b *uint256.Int) *uint256.Int) error {
	// Pop an element from the stack.
	a, err := e.stack.Pop()
	if err != nil {
		return err
	}

	// Pop another element from the stack.
	b, err := e.stack.Pop()
	if err != nil {
		return err
	}

	// Perform the operation on the two elements.
	result := op(a, b)

	// Push the result back to the stack.
	// This step should never fail because of an overflow.
	// Indeed, two elements are popped from the stack and only one is pushed back.
	_ = e.stack.Push(result)
	return nil
}
