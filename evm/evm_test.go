package evm

import (
	"bytes"
	"testing"

	"github.com/holiman/uint256"
)

// ExtendedEVM defines the methods that an Ethereum Virtual Machine implementation should have,
// including helper methods for testing purposes.
type ExtendedEVM interface {
	IEVM

	// Helper methods.
	// Push an item to the stack.
	HelperPush(*uint256.Int) error
	// Pop an item of the stack.
	HelperPop() (*uint256.Int, error)
	// Write byte slice to memory at the specified offset.
	HelperStore(value []byte, offset int)
	// Load a chunck of the memory.
	HelperLoad(size int) []byte
}

func (e *EVM) HelperPop() (*uint256.Int, error) {
	return e.stack.Pop()
}

func (e *EVM) HelperPush(value *uint256.Int) error {
	return e.stack.Push(value)
}

func (e *EVM) HelperStore(value []byte, offset int) {
	e.memory.Store(value, offset)
}

func (e *EVM) HelperLoad(size int) []byte {
	return e.memory.Load(0, size)
}

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
	testStackOperationWithNewEVM(t, addOp, ErrStackUnderflow, emptyStack, emptyStack, nil, nil, nil)
	testStackOperationWithNewEVM(t, addOp, ErrStackUnderflow, oneElementStack, emptyStack, nil, nil, nil)

	// 3 operands arithmetic operation.
	addModOp := func(evm IEVM) error { return evm.AddMod() }
	testStackOperationWithNewEVM(t, addModOp, ErrStackUnderflow, emptyStack, emptyStack, nil, nil, nil)
	testStackOperationWithNewEVM(t, addModOp, ErrStackUnderflow, oneElementStack, emptyStack, nil, nil, nil)
	testStackOperationWithNewEVM(t, addModOp, ErrStackUnderflow, twoElementsStack, emptyStack, nil, nil, nil)

	// 2 operands comparison operation.
	eqOp := func(evm IEVM) error { return evm.Eq() }
	testStackOperationWithNewEVM(t, eqOp, ErrStackUnderflow, emptyStack, emptyStack, nil, nil, nil)
	testStackOperationWithNewEVM(t, eqOp, ErrStackUnderflow, oneElementStack, emptyStack, nil, nil, nil)
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

	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil, nil)
}

// Helper function to test stack operations with a fresh new EVM.
func testStackOperationWithNewEVM(t *testing.T, op func(evm IEVM) error, expectedErr error, initialStack, expectedStack []uint64, initialMemory, expectedMemory, code []byte) {
	evm := NewEVM(code)
	testStackOperationWithExistingEVM(t, evm, op, expectedErr, initialStack, expectedStack, initialMemory, expectedMemory)
}

// Helper function to test stack operations with an existing EVM.
func testStackOperationWithExistingEVM(t *testing.T, evm IEVM, op func(evm IEVM) error, expectedErr error, initialStack, expectedStack []uint64, initialMemory, expectedMemory []byte) {
	// Extend the capabilities of the EVM using the internal EVM which defines helper methods to access the states of the stack and the memory.
	testEvm, ok := evm.(ExtendedEVM)
	if !ok {
		t.Fatal("IEVM does not implement internalEVM")
	}

	// Push initial elements to the stack.
	for i, v := range initialStack {
		if err := testEvm.HelperPush(uint256.NewInt(v)); err != nil {
			t.Errorf("Push() returned an unexpected error at iteration %d: %v", i, err)
		}
	}

	// Load initial elements to the memory.
	testEvm.HelperStore(initialMemory, 0)

	// Perform the operation.
	if err := op(evm); err != expectedErr {
		t.Errorf("Operation returned an unexpected error: %v, wanted: %v", err, expectedErr)
	}

	// Check the stack after the operation.
	for i := len(expectedStack) - 1; i >= 0; i-- {
		popped, err := testEvm.HelperPop()
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

	// Check the memory after the operation.
	actualMemory := testEvm.HelperLoad(len(expectedMemory))
	if !bytes.Equal(actualMemory, expectedMemory) {
		t.Errorf("Memory mismatch. Expected: %v, got: %v", expectedMemory, actualMemory)
	}
}
