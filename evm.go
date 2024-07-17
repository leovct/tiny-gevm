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
	// Add the top two elements of the stack and push the result, x + y, back to the stack.
	Add() error

	// Multiply the top two elements of the stack and push the result, x * y, back to the stack.
	Mul() error

	// Subtract the top two elements of the stack and push the result, x - y, back to the stack.
	Sub() error

	// Perform the integer divison operation on the top two elements of the stack and push the result, x // y, back to the stack.
	Div() error

	// Perform the signed integer division operation (trunced) on the top two elements of the stack and push the result, x // y, back to the stack.
	SDiv() error

	// Perform the modulo remained operation on the top two elements of the stack and push the result, x % y, back to the stack.
	Mod() error

	// Perform the signed modulo remained operation on the top two elements of the stack and push the result, x % y, back to the stack.
	SMod() error

	// Perform the modulo addition operation on the top two elements of the stack and push the result, (x + y) % m, back to the stack.
	// The third top element of the stack is the integer denominator m.
	AddMod() error

	// Perform the modulo multiplication operation on the top two elements of the stack and push the result, (x * y) % m, back to the stack.
	// The third top element of the stack is the integer denominator N.
	MulMod() error

	// Perform the exponential operation on the top two elements of the stack and push the result, x ** y, back to the stack.
	Exp() error

	// Extend the length of two’s complement signed integer.
	// The first top element of the stack, b, represents the size in byte - 1 of the integer to sign extend.
	// The second top element of the stack, x, represents the integer value to sign extend.
	SignExtend() error
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
	op := func(x, y *uint256.Int) *uint256.Int {
		return new(uint256.Int).Add(x, y)
	}
	return e.performStackOperation(op)
}

func (e *EVM) Mul() error {
	op := func(x, y *uint256.Int) *uint256.Int {
		return new(uint256.Int).Mul(x, y)
	}
	return e.performStackOperation(op)
}

func (e *EVM) Sub() error {
	op := func(x, y *uint256.Int) *uint256.Int {
		return new(uint256.Int).Sub(x, y)
	}
	return e.performStackOperation(op)
}

func (e *EVM) Div() error {
	op := func(x, y *uint256.Int) *uint256.Int {
		return new(uint256.Int).Div(x, y)
	}
	return e.performStackOperation(op)
}

func (e *EVM) SDiv() error {
	op := func(x, y *uint256.Int) *uint256.Int {
		return new(uint256.Int).SDiv(x, y)
	}
	return e.performStackOperation(op)
}

func (e *EVM) Mod() error {
	op := func(x, y *uint256.Int) *uint256.Int {
		return new(uint256.Int).Mod(x, y)
	}
	return e.performStackOperation(op)
}

func (e *EVM) SMod() error {
	op := func(x, y *uint256.Int) *uint256.Int {
		return new(uint256.Int).SMod(x, y)
	}
	return e.performStackOperation(op)
}

func (e *EVM) AddMod() error {
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

	// Pop a last element from the stack.
	m, err := e.stack.Pop()
	if err != nil {
		return err
	}

	// Perform the modulo addition operation.
	result := new(uint256.Int).AddMod(x, y, m)
	_ = e.stack.Push(result)
	return nil
}

func (e *EVM) MulMod() error {
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

	// Pop a last element from the stack.
	m, err := e.stack.Pop()
	if err != nil {
		return err
	}

	// Perform the modulo multiplication operation.
	result := new(uint256.Int).MulMod(x, y, m)
	_ = e.stack.Push(result)
	return nil
}

func (e *EVM) Exp() error {
	op := func(x, y *uint256.Int) *uint256.Int {
		return new(uint256.Int).Exp(x, y)
	}
	return e.performStackOperation(op)
}

func (e *EVM) SignExtend() error {
	op := func(x, y *uint256.Int) *uint256.Int {
		return new(uint256.Int).ExtendSign(x, y)
	}
	return e.performStackOperation(op)
}

// Perform a binary operation on the top two elements on the stack.
// It pops two values from the stack, applies the operation, and pushes the result back to the stack.
func (e *EVM) performStackOperation(op func(x, y *uint256.Int) *uint256.Int) error {
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
	_ = e.stack.Push(result)
	return nil
}
