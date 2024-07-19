package main

import "github.com/holiman/uint256"

// IComparisonAndBitwiseOps defines comparison and bitwise logic operations for the EVM.
// All operations pop their operands from the stack and push the result back.
// All methods return an error if there are not enough elements on the stack.
type IComparisonAndBitwiseOps interface {
	// Less-than comparison.
	// Stack: [x, y, ...] -> [x < y, ...]
	Lt() error

	// Greater-than comparison.
	// Stack: [x, y, ...] -> [x > y, ...]
	Gt() error

	// Signed less-than comparison.
	// Stack: [x, y, ...] -> [x < y, ...]
	SLt() error

	// Signed greater-than comparison.
	// Stack: [x, y, ...] -> [x > y, ...]
	SGt() error

	// Equality comparison.
	// Stack: [x, y, ...] -> [x == y, ...]
	Eq() error

	// Is-zero comparison.
	// Stack: [x, ...] -> [x == 0, ...]
	IsZero() error

	// Bitwise AND operation.
	// Stack: [x, y, ...] -> [x & y, ...]
	And() error

	// Bitwise OR operation.
	// Stack: [x, y, ...] -> [x | y, ...]
	Or() error

	// Bitwise XOR operation.
	// Stack: [x, y, ...] -> [x ^ y, ...]
	Xor() error

	// Bitwise NOT operation.
	// Stack: [x, ...] -> [~x, ...]
	Not() error

	// Retrieve single byte from a 32-byte word.
	// Stack: [i, x, ...] -> [y, ...]
	// Where y is the i'th byte of x, counting from the least significant byte.
	// If i >= 32, the result is 0.
	Byte() error

	// Left shift operation.
	// Shift the bits towards the most significant one.
	// The bits moved after the 256th one are discarded, the new bits are set to 0.
	// It is equivalent to multiplying value by 2 ** shift.
	// Stack: [shift, value, ...] -> [value << shift, ...]
	Shl() error

	// Right shift operation.
	// Shift the bits towards the least significant one.
	// The bits moved before the first one are discarded, the new bits are set to 0.
	// Stack: [shift, value, ...] -> [value >> shift, ...]
	// It is equivalent to dividing value by 2 ** shift, throwing out any remainders.
	Shr() error

	// Arithmetic (signed) right shift operation
	// Shift the bits towards the least significant one.
	// The bits moved before the first one are discarded, the new bits are set to 0 if the previous most significant bit was 0, otherwise the new bits are set to 1.
	// Stack: [shift, value, ...] -> [value >> shift, ...]
	Sar() error
}

func (e *EVM) Lt() error {
	op := func(x, y *uint256.Int) bool {
		return x.Lt(y)
	}
	return e.performComparisonStackOperation(op)
}

func (e *EVM) Gt() error {
	op := func(x, y *uint256.Int) bool {
		return x.Gt(y)
	}
	return e.performComparisonStackOperation(op)
}

func (e *EVM) SLt() error {
	op := func(x, y *uint256.Int) bool {
		return x.Slt(y)
	}
	return e.performComparisonStackOperation(op)
}

func (e *EVM) SGt() error {
	op := func(x, y *uint256.Int) bool {
		return x.Sgt(y)
	}
	return e.performComparisonStackOperation(op)
}

func (e *EVM) Eq() error {
	op := func(x, y *uint256.Int) bool {
		return x.Eq(y)
	}
	return e.performComparisonStackOperation(op)
}

func (e *EVM) IsZero() error {
	op := func(x, y *uint256.Int) bool {
		return x.IsZero()
	}
	return e.performComparisonStackOperation(op)
}

func (e *EVM) And() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y := operands[0], operands[1]
		return new(uint256.Int).And(x, y)
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) Or() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y := operands[0], operands[1]
		return new(uint256.Int).Or(x, y)
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) Xor() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x, y := operands[0], operands[1]
		return new(uint256.Int).Xor(x, y)
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) Not() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		x := operands[0]
		return new(uint256.Int).Not(x)
	}
	return e.performBinaryStackOperation(1, op)
}

func (e *EVM) Byte() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		i, x := operands[0], operands[1]
		return x.Byte(i)
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) Shl() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		value, shift := operands[0], operands[1]
		return new(uint256.Int).Lsh(value, uint(shift.Uint64()))
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) Shr() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		value, shift := operands[0], operands[1]
		return new(uint256.Int).Rsh(value, uint(shift.Uint64()))
	}
	return e.performBinaryStackOperation(2, op)
}

func (e *EVM) Sar() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		value, shift := operands[0], operands[1]
		return new(uint256.Int).SRsh(value, uint(shift.Uint64()))
	}
	return e.performBinaryStackOperation(2, op)
}
