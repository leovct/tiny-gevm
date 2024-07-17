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
	op := func(evm IEVM) error { return evm.Add() }
	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 5}
	testStackOperation(t, op, nil, initialStack, expectedStack)
}

func TestMul(t *testing.T) {
	op := func(evm IEVM) error { return evm.Mul() }
	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 6}
	testStackOperation(t, op, nil, initialStack, expectedStack)
}

func TestStackOperationOnEmptyStack(t *testing.T) {
	op := func(evm IEVM) error { return evm.Add() }
	initialStack := []uint64{}
	expectedStack := []uint64{}
	testStackOperation(t, op, ErrStackUnderflow, initialStack, expectedStack)
}

func TestStackOperationOnOneElementStack(t *testing.T) {
	op := func(evm IEVM) error { return evm.Add() }
	initialStack := []uint64{1}
	expectedStack := []uint64{}
	testStackOperation(t, op, ErrStackUnderflow, initialStack, expectedStack)
}

func TestStackOperationOnFullStack(t *testing.T) {
	op := func(evm IEVM) error { return evm.Add() }

	initialStack := make([]uint64, 1024)
	for i := range initialStack {
		initialStack[i] = 1
	}

	expectedStack := make([]uint64, 1023)
	copy(expectedStack, initialStack)
	expectedStack[1022] = 2

	testStackOperation(t, op, nil, initialStack, expectedStack)
}

// Helper function to test stack operations.
func testStackOperation(t *testing.T, op func(evm IEVM) error, expectedErr error, initialStack []uint64, expectedStack []uint64) {
	// Create a new EVM.
	evm := NewEVM()

	// Push initial elements to the stack.
	for _, v := range initialStack {
		if err := evm.Push(uint256.NewInt(v)); err != nil {
			t.Errorf("Push() returned an unexpected error: %v", err)
		}
	}

	// Perform the operation.
	if err := op(evm); err != expectedErr {
		t.Errorf("Operation returned an unexpected error: %v, wanted: %v", err, expectedErr)
	}

	// Check the stack after the operation.
	for i := len(expectedStack) - 1; i >= 0; i-- {
		popped, err := evm.Pop()
		if err != nil {
			t.Errorf("Pop() returned an unexpected error: %v", err)
		}

		expectedValue := expectedStack[i]
		if popped == nil {
			t.Errorf("Expected %v, got %v", expectedValue, popped)
		} else {
			if popped.Uint64() != expectedValue {
				t.Errorf("Expected %v, got %v", expectedValue, popped)
			}
		}
	}
}
