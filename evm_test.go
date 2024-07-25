package main

import (
	"testing"

	"github.com/holiman/uint256"
)

func TestNewEVM(t *testing.T) {
	evm := NewEVM(nil)
	if evm == nil {
		t.Error("NewEVM() returned nil")
	}
}

func TestStackOperationUnderflows(t *testing.T) {
	var emptyStack []uint64
	oneElementStack := []uint64{1}
	twoElementsStack := []uint64{1, 2}

	// 2 operands arithmetic operation.
	addOp := func(evm IEVM) error { return evm.Add() }
	testStackOperationWithNewEVM(t, addOp, ErrStackUnderflow, emptyStack, emptyStack, nil, nil)
	testStackOperationWithNewEVM(t, addOp, ErrStackUnderflow, oneElementStack, emptyStack, nil, nil)

	// 3 operands arithmetic operation.
	addModOp := func(evm IEVM) error { return evm.AddMod() }
	testStackOperationWithNewEVM(t, addModOp, ErrStackUnderflow, emptyStack, emptyStack, nil, nil)
	testStackOperationWithNewEVM(t, addModOp, ErrStackUnderflow, oneElementStack, emptyStack, nil, nil)
	testStackOperationWithNewEVM(t, addModOp, ErrStackUnderflow, twoElementsStack, emptyStack, nil, nil)

	// 2 operands comparison operation.
	eqOp := func(evm IEVM) error { return evm.Eq() }
	testStackOperationWithNewEVM(t, eqOp, ErrStackUnderflow, emptyStack, emptyStack, nil, nil)
	testStackOperationWithNewEVM(t, eqOp, ErrStackUnderflow, oneElementStack, emptyStack, nil, nil)
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

	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil)
}

// Helper function to test stack operations with a fresh new EVM.
func testStackOperationWithNewEVM(t *testing.T, op func(evm IEVM) error, expectedErr error, initialStack []uint64, expectedStack []uint64, memory, code []byte) {
	evm := NewEVM(code)
	testStackOperationWithExistingEVM(t, evm, op, expectedErr, initialStack, expectedStack, memory)
}

// Helper function to test stack operations with an existing EVM.
func testStackOperationWithExistingEVM(t *testing.T, evm IEVM, op func(evm IEVM) error, expectedErr error, initialStack []uint64, expectedStack []uint64, memory []byte) {
	// Push initial elements to the stack.
	for i, v := range initialStack {
		if err := evm.HelperPush(uint256.NewInt(v)); err != nil {
			t.Errorf("Push() returned an unexpected error at iteration %d: %v", i, err)
		}
	}

	// Load initial elements to the memory.
	evm.HelperStore(memory, 0)

	// Perform the operation.
	if err := op(evm); err != expectedErr {
		t.Errorf("Operation returned an unexpected error: %v, wanted: %v", err, expectedErr)
	}

	// Check the stack after the operation.
	for i := len(expectedStack) - 1; i >= 0; i-- {
		popped, err := evm.HelperPop()
		if err != nil {
			t.Errorf("Pop() returned an unexpected error: %v", err)
		}

		expectedValue := expectedStack[i]
		if popped == nil {
			t.Errorf("Expected %v, got %v", expectedValue, popped.Uint64())
		} else {
			if popped.Uint64() != expectedValue {
				t.Errorf("Expected %v, got %v", expectedValue, popped.Uint64())
			}
		}
	}
}
