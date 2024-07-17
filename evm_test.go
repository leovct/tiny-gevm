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
	expectedStack := []uint64{1, 5} // 5 = 3+2
	testStackOperation(t, op, nil, initialStack, expectedStack)
}

func TestMul(t *testing.T) {
	op := func(evm IEVM) error { return evm.Mul() }
	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 6} // 6 = 3*2
	testStackOperation(t, op, nil, initialStack, expectedStack)
}

func TestSub(t *testing.T) {
	op := func(evm IEVM) error { return evm.Sub() }
	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 1} // 1 = 3-2
	testStackOperation(t, op, nil, initialStack, expectedStack)
}

func TestDiv(t *testing.T) {
	op := func(evm IEVM) error { return evm.Div() }
	initialStack := []uint64{1, 2, 4}
	expectedStack := []uint64{1, 2} // 2 = 4/2
	testStackOperation(t, op, nil, initialStack, expectedStack)
}

func TestSDiv(t *testing.T) {
	op := func(evm IEVM) error { return evm.SDiv() }
	initialStack := []uint64{1, 2, 4}
	expectedStack := []uint64{1, 2} // 2 = 4/2
	testStackOperation(t, op, nil, initialStack, expectedStack)
}

func TestMod(t *testing.T) {
	op := func(evm IEVM) error { return evm.Mod() }
	initialStack := []uint64{1, 5, 12}
	expectedStack := []uint64{1, 2} // 2 = 12%5
	testStackOperation(t, op, nil, initialStack, expectedStack)
}

func TestSMod(t *testing.T) {
	op := func(evm IEVM) error { return evm.SMod() }
	initialStack := []uint64{1, 5, 12}
	expectedStack := []uint64{1, 2} // 2 = 12%5
	testStackOperation(t, op, nil, initialStack, expectedStack)
}

func TestAddMod(t *testing.T) {
	op := func(evm IEVM) error { return evm.AddMod() }
	initialStack := []uint64{1, 7, 2, 15}
	expectedStack := []uint64{1, 3} // 3 = (15+2)%7
	testStackOperation(t, op, nil, initialStack, expectedStack)
}

func TestMulMod(t *testing.T) {
	op := func(evm IEVM) error { return evm.MulMod() }
	initialStack := []uint64{1, 7, 2, 15}
	expectedStack := []uint64{1, 2} // 2 = (15*2)%7
	testStackOperation(t, op, nil, initialStack, expectedStack)
}

func TestExp(t *testing.T) {
	op := func(evm IEVM) error { return evm.Exp() }
	initialStack := []uint64{1, 2, 4}
	expectedStack := []uint64{1, 16} // 16 = 2**4
	testStackOperation(t, op, nil, initialStack, expectedStack)
}

func TestSignExtend(t *testing.T) {
	op := func(evm IEVM) error { return evm.SignExtend() }
	initialStack := []uint64{1, 0, 4}
	expectedStack := []uint64{1, 4}
	testStackOperation(t, op, nil, initialStack, expectedStack)
}

func TestStackOperationUnderflows(t *testing.T) {
	var emptyStack []uint64
	oneElementStack := []uint64{1}
	twoElementsStack := []uint64{1, 2}

	// Any operation similar to Add
	addOp := func(evm IEVM) error { return evm.Add() }
	testStackOperation(t, addOp, ErrStackUnderflow, emptyStack, emptyStack)
	testStackOperation(t, addOp, ErrStackUnderflow, oneElementStack, emptyStack)

	// AddMod
	addModOp := func(evm IEVM) error { return evm.AddMod() }
	testStackOperation(t, addModOp, ErrStackUnderflow, emptyStack, emptyStack)
	testStackOperation(t, addModOp, ErrStackUnderflow, oneElementStack, emptyStack)
	testStackOperation(t, addModOp, ErrStackUnderflow, twoElementsStack, emptyStack)

	// MulMod
	mulModOp := func(evm IEVM) error { return evm.MulMod() }
	testStackOperation(t, mulModOp, ErrStackUnderflow, emptyStack, emptyStack)
	testStackOperation(t, mulModOp, ErrStackUnderflow, oneElementStack, emptyStack)
	testStackOperation(t, mulModOp, ErrStackUnderflow, twoElementsStack, emptyStack)
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
