package evm

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
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y := operands[0], operands[1]
		return new(uint256.Int).Add(x, y)
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) Mul() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y := operands[0], operands[1]
		return new(uint256.Int).Mul(x, y)
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) Sub() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y := operands[0], operands[1]
		return new(uint256.Int).Sub(x, y)
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) Div() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y := operands[0], operands[1]
		return new(uint256.Int).Div(x, y)
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) SDiv() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y := operands[0], operands[1]
		return new(uint256.Int).SDiv(x, y)
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) Mod() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y := operands[0], operands[1]
		return new(uint256.Int).Mod(x, y)
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) SMod() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y := operands[0], operands[1]
		return new(uint256.Int).SMod(x, y)
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) AddMod() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y, m := operands[0], operands[1], operands[2]
		return new(uint256.Int).AddMod(x, y, m)
	}
	return e.performBinaryStackOperation(3, op)
}

func (e *EVM) MulMod() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y, m := operands[0], operands[1], operands[2]
		return new(uint256.Int).MulMod(x, y, m)
	}
	return e.performBinaryStackOperation(3, op)
}

func (e *EVM) Exp() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y := operands[0], operands[1]
		return new(uint256.Int).Exp(x, y)
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) SignExtend() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y := operands[0], operands[1]
		return new(uint256.Int).ExtendSign(x, y)
	}
	return e.performBinaryStackOperation(2, op)
}
