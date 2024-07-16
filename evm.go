package main

import "github.com/holiman/uint256"

// IEVM defines the methods that an Ethereum Virtual Machine implementation should have.
type IEVM interface{}

// EVM represents an Ethereum Virtual Machine.
type EVM struct {
	stack   IStack
	memory  IMemory
	storage IStorage
}

// NewEVM creates and returns a new EVM instance.
func NewEVM() *EVM {
	return &EVM{
		stack:   NewStack(),
		memory:  NewMemory(),
		storage: NewStorage(),
	}
}

// Add performs addition of the top two elements on the stack.
// It pops two values from the stack, adds them, and pushes the result back to the stack.
func (e *EVM) Add() error {
	a, err := e.stack.Pop()
	if err != nil {
		return err
	}

	b, err := e.stack.Pop()
	if err != nil {
		return err
	}

	result := new(uint256.Int).Add(a, b)
	// This step should never fail because of an overflow.
	// Indeed, two elements are popped from the stack and only one is pushed back.
	_ = e.stack.Push(result)
	return nil
}
