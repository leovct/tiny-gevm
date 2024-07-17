package main

import (
	"testing"

	"github.com/holiman/uint256"
)

func TestNewEVM(t *testing.T) {
	evm := NewEVM()
	if evm == nil {
		t.Error("NewEVM() returned nil")
	}
}

func TestAdd(t *testing.T) {
	// Create a new EVM.
	evm := NewEVM()

	// Push some elements to the stack.
	// The stack should be equal to [0x1, 0x2, 0x3].
	err := evm.Push(uint256.NewInt(1))
	if err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	err = evm.Push(uint256.NewInt(2))
	if err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	err = evm.Push(uint256.NewInt(3))
	if err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	// Add the two elements at the top of the stack.
	// The stack should be equal to [0x1, 0x5].
	err = evm.Add()
	if err != nil {
		t.Errorf("Add() returned an unexpected error: %v", err)
	}

	// Pop the element from the stack to check the result.
	result, err := evm.Pop()
	expectedResult := uint256.NewInt(5)
	if err != nil {
		t.Errorf("Pop() returned an unexpected error: %v", err)
	}
	if !result.Eq(expectedResult) {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestAddOnEmptyStack(t *testing.T) {
	// Create a new EVM.
	evm := NewEVM()

	// Try to add the two elements at the top of the stack.
	err := evm.Add()
	if err == nil {
		t.Error("Add() should return an error because there are no elements in the stack")
	}
}

func TestAddOnOneElementStack(t *testing.T) {
	// Create a new EVM.
	evm := NewEVM()

	// Add one element to the stack.
	// The stack should be equal to [0x1].
	err := evm.Push(uint256.NewInt(1))
	if err != nil {
		t.Errorf("Push() returned an unexpected error: %v", err)
	}

	// Try to add the two elements at the top of the stack.
	err = evm.Add()
	if err == nil {
		t.Error("Add() should return an error because there is only one element in the stack")
	}
}

func TestAddOnFullStack(t *testing.T) {
	// Create a new EVM.
	evm := NewEVM()

	// Push 1024 elements to the stack.
	// The stack should contain 1024 0x1 elements.
	for i := 0; i < MAX_STACK_SIZE; i++ {
		err := evm.Push(uint256.NewInt(1))
		if err != nil {
			t.Errorf("Push() returned an unexpected error on iteration %d: %v", i, err)
		}
	}

	// Try to add the two elements at the top of the stack.
	err := evm.Add()
	if err != nil {
		t.Error("Add() should return an error because the stack is full")
	}
}
