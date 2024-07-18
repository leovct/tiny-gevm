package main

import "github.com/holiman/uint256"

// IArithmeticOps defines arithmetic operations for the EVM.
// All operations pop their operands from the stack and push the result back.
// All methods return an error if there are not enough elements on the stack.
type IArithmeticOps interface {
	// Addition operation
	// Stack: [x, y, ...] -> [x + y, ...]
	Add() error

	// Multiplication operation
	// Stack: [x, y, ...] -> [x * y, ...]
	Mul() error

	// Subtraction operation
	// Stack: [x, y, ...] -> [x - y, ...]
	Sub() error

	// Integer division operation
	// Stack: [x, y, ...] -> [x // y, ...]
	Div() error

	// Signed interger division operation (truncated)
	// Stack: [x, y, ...] -> [x // y, ...]
	SDiv() error

	// Modulo remainder operation
	// Stack: [x, y, ...] -> [x % y, ...]
	Mod() error

	// Signed modulo remainder operation
	// Stack: [x, y, ...] -> [x % y, ...]
	SMod() error

	// Modulo addition operation
	// Stack: [x, y, m, ...] -> [(x + y) % m, ...]
	AddMod() error

	// Modulo multiplication operation
	// Stack: [x, y, m, ...] -> [(x * y) % m, ...]
	MulMod() error

	// Exponential operation
	// Stack: [x, y, ...] -> [x ** y, ...]
	Exp() error

	// Extend length of two’s complement signed integer
	// Stack: [x, b, ...] -> [signextend(x, b), ...]
	SignExtend() error
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
	x, err := e.stack.Pop()
	if err != nil {
		return err
	}

	y, err := e.stack.Pop()
	if err != nil {
		return err
	}

	m, err := e.stack.Pop()
	if err != nil {
		return err
	}

	result := new(uint256.Int).AddMod(x, y, m)
	_ = e.stack.Push(result)
	return nil
}

func (e *EVM) MulMod() error {
	x, err := e.stack.Pop()
	if err != nil {
		return err
	}

	y, err := e.stack.Pop()
	if err != nil {
		return err
	}

	m, err := e.stack.Pop()
	if err != nil {
		return err
	}

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
