package main

import (
	"github.com/holiman/uint256"
)

// IEVM defines the methods that an Ethereum Virtual Machine implementation should have.
type IEVM interface {
	// Stack operations.
	// Push an item to the stack.
	Push(*uint256.Int) error
	// Pop an item from the stack.
	Pop() (*uint256.Int, error)

	// Arithmetic operations.
	IArithmeticOps

	// Comparison and bitwise logic operations.
	IComparisonAndBitwiseOps
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

// Perform an arithmetic or a bitwise operation on the top two elements on the stack.
// It pops two values from the stack, applies the operation, and pushes the result back to the stack.
func (e *EVM) performBinaryStackOperation(numOperands int, operation func(...*uint256.Int) *uint256.Int) error {
	// Check if there are enough elements on the stack.
	if e.stack.Size() < numOperands {
		return ErrStackUnderflow
	}

	// Pop the required number of elements from the stack.
	operands := make([]*uint256.Int, numOperands)
	for i := 0; i < numOperands; i++ {
		var err error
		operands[i], err = e.stack.Pop()
		if err != nil {
			return err
		}
	}

	// Perform the operation on the elements
	result := operation(operands...)

	// Push the result back to the stack.
	// This step should never fail because of an overflow.
	// Indeed, two elements are popped from the stack and only one is pushed back.
	_ = e.stack.Push(result)
	return nil
}

// Perform a comparison operation on the top two elements on the stack.
// It pops two values from the stack, applies the operation, and pushes the result back to the stack.
// If the result is true, push one to the stack, else push zero.
func (e *EVM) performComparisonStackOperation(op func(x, y *uint256.Int) bool) error {
	// Pop an element from the stack.
	x, err := e.stack.Pop()
	if err != nil {
		return err
	}

	// Pop another element from the stack.
	y, err := e.stack.Pop()
	if err != nil {
		return err
	}

	// Perform the operation on the two elements.
	result := op(x, y)

	// Push the result back to the stack.
	// This step should never fail because of an overflow.
	// Indeed, two elements are popped from the stack and only one is pushed back.
	if result {
		_ = e.stack.Push(uint256.NewInt(1))
	} else {
		_ = e.stack.Push(uint256.NewInt(0))
	}
	return nil
}
