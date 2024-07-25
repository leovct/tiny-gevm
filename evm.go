package main

import (
	"github.com/holiman/uint256"
)

// IEVM defines the methods that an Ethereum Virtual Machine implementation should have.
type IEVM interface {
	IArithmeticOps
	IComparisonAndBitwiseOps
	ISHA3Ops
	IStackOps
	IMemoryOps

	// Helper methods.
	// TODO: Get rid of these methods if possible.
	// Push an item to the stack.
	HelperPush(*uint256.Int) error
	// Pop an item of the stack.
	HelperPop() (*uint256.Int, error)
	// Write byte slice to memory at the specified offset.
	HelperStore(value []byte, offset int)
}

// EVM represents an Ethereum Virtual Machine.
type EVM struct {
	stack   IStack
	memory  IMemory
	storage IStorage
	env     ExecutionEnvironment
	state   MachineState
}

// ExecutionEnvironment represents the EVM execution environment.
type ExecutionEnvironment struct {
	// Machine code to be executed by the EVM.
	code []byte
}

// MachineState represents the EVM state.
type MachineState struct {
	// Program counter.
	pc int
}

// NewEVM creates and returns a new EVM instance.
func NewEVM(code []byte) IEVM {
	return &EVM{
		stack:   NewStack(),
		memory:  NewMemory(),
		storage: NewStorage(),
		env: ExecutionEnvironment{
			code: code,
		},
		state: MachineState{
			pc: 0,
		},
	}
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

// Perform an arithmetic or a bitwise operation on the top two elements on the stack.
// It pops two values from the stack, applies the operation, and pushes the result back to the stack.
func (e *EVM) performBinaryStackOperation(numOperands int, operation func(...*uint256.Int) *uint256.Int) error {
	// Check if there are enough elements on the stack.
	if e.stack.Size() <= numOperands {
		return ErrStackUnderflow
	}

	// Pop the required number of elements from the stack.
	operands := make([]*uint256.Int, numOperands)
	for i := 0; i < numOperands; i++ {
		var err error
		operands[i], err = e.stack.Pop()
		if err != nil {
			// Unreachable in theory.
			// This step should never fail because of an overflow.
			// Indeed, we perform a check to make sure the number of operands is always smaller or equal to the number of elements in the stack.
			return err
		}
	}

	// Perform the operation on the elements
	result := operation(operands...)

	// Push the result back to the stack.
	if err := e.stack.Push(result); err != nil {
		// Unreachable in theory.
		// This step should never fail because of an overflow.
		// Indeed, two elements are popped from the stack and only one is pushed back.
		return err
	}
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
	boolResult := op(x, y)

	// Push the result back to the stack.
	// This step should never fail because of an overflow.
	// Indeed, two elements are popped from the stack and only one is pushed back.
	var intResult *uint256.Int
	if boolResult {
		intResult = uint256.NewInt(1)
	} else {
		intResult = uint256.NewInt(0)
	}
	if err := e.stack.Push(intResult); err != nil {
		// Unreachable in theory.
		// This step should never fail because of an overflow.
		// Indeed, two elements are popped from the stack and only one is pushed back.
		return err
	}
	return nil
}
