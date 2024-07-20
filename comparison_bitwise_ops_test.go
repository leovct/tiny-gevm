package main

import (
	"testing"
)

func TestLt(t *testing.T) {
	op := func(evm IEVM) error { return evm.Lt() }

	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 0}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)

	initialStack = []uint64{1, 3, 2}
	expectedStack = []uint64{1, 1}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}

func TestGt(t *testing.T) {
	op := func(evm IEVM) error { return evm.Gt() }

	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 1}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)

	initialStack = []uint64{1, 3, 2}
	expectedStack = []uint64{1, 0}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}

func TestSLt(t *testing.T) {
	op := func(evm IEVM) error { return evm.SLt() }

	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 0}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)

	initialStack = []uint64{1, 3, 2}
	expectedStack = []uint64{1, 1}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}

func TestSGt(t *testing.T) {
	op := func(evm IEVM) error { return evm.SGt() }

	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 1}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)

	initialStack = []uint64{1, 3, 2}
	expectedStack = []uint64{1, 0}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}

func TestEq(t *testing.T) {
	op := func(evm IEVM) error { return evm.Eq() }

	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 0}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)

	initialStack = []uint64{1, 3, 3}
	expectedStack = []uint64{1, 1}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}

func TestIsZero(t *testing.T) {
	op := func(evm IEVM) error { return evm.IsZero() }

	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 0}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)

	initialStack = []uint64{1, 2, 0}
	expectedStack = []uint64{1, 1}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}

func TestAnd(t *testing.T) {
	op := func(evm IEVM) error { return evm.And() }
	initialStack := []uint64{1, 2, 3}
	// Apply the operation: 3 | 2 or 0b11 & 0b10 = 0b10
	// Stack: [1, 2, 3] -> [1, 2]
	expectedStack := []uint64{1, 2}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}

func TestOr(t *testing.T) {
	op := func(evm IEVM) error { return evm.Or() }
	initialStack := []uint64{1, 2, 3}
	// Apply the operation: 3 | 2 or 0b11 | 0b10 = 0b11
	// Stack: [1, 2, 3] -> [1, 3]
	expectedStack := []uint64{1, 3}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}

func TestXor(t *testing.T) {
	op := func(evm IEVM) error { return evm.Xor() }
	initialStack := []uint64{1, 2, 3}
	// Apply the operation: 3 ^ 2 or 0b11 ^ 0b10 = 0b1
	// Stack: [1, 2, 3] -> [1, 1]
	expectedStack := []uint64{1, 1}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}

func TestNot(t *testing.T) {
	op := func(evm IEVM) error { return evm.Not() }
	initialStack := []uint64{1, 2, 0x0000000000000003}
	// The binary representation of 3 over 8 bytes is the following:
	// 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000011
	// 7 bytes full of zeros and a final byte, 00000011.
	// The result of the NOT operation is equal to:
	// 11111111 11111111 11111111 11111111 11111111 11111111 11111111 11111100
	// 7 bytes full of ones and a final byte, 11111100.
	// The representation of this number is 0xFFFFFFFFFFFFFFFC in hexadecimal.
	// Stack: [1, 2, 0x0000000000000003] -> [1, 2, 0xFFFFFFFFFFFFFFFC]
	expectedStack := []uint64{1, 2, 0xFFFFFFFFFFFFFFFC}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}

func TestByte(t *testing.T) {
	op := func(evm IEVM) error { return evm.Byte() }

	// For these examples, we use the number 0x11223344.
	// The 31st byte is 0x44
	// The 30th byte is 0x33
	// The 29th byte is 0x22
	// The 28th byte is 0x11
	// The other bytes are equal to 0x0

	initialStack := []uint64{1, 2, 3, 0x11223344, 31}
	// Stack: [1, 2, 3, 0x11223344, 31] -> [1, 2, 3, 0x44]
	expectedStack := []uint64{1, 2, 3, 0x44}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)

	initialStack = []uint64{1, 2, 3, 0x11223344, 28}
	// Stack: [1, 2, 3, 0x11223344, 28] -> [1, 2, 3, 0x11]
	expectedStack = []uint64{1, 2, 3, 0x11}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)

	initialStack = []uint64{1, 2, 3, 0x11223344, 27}
	// Stack: [1, 2, 3, 0x11223344, 27] -> [1, 2, 3, 0x0]
	expectedStack = []uint64{1, 2, 3, 0x0}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)

	// A last test just to check something.
	initialStack = []uint64{1, 2, 3, 31}
	// Stack: [1, 2, 3, 31] -> [1, 2, 3]
	expectedStack = []uint64{1, 2, 3}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}

func TestShl(t *testing.T) {
	op := func(evm IEVM) error { return evm.Shl() }
	initialStack := []uint64{1, 2, 3}
	// Apply the operation: 3 << 2 or 0b11 << 0b10 = 0b1100
	// Stack: [1, 2, 3] -> [1, 12]
	expectedStack := []uint64{1, 12}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}

func TestShr(t *testing.T) {
	op := func(evm IEVM) error { return evm.Shr() }
	initialStack := []uint64{1, 2, 3}
	// Apply the operation: 3 >> 2 or 0b11 >> 0b10 = 0b0
	// Stack: [1, 2, 3] -> [1, 0]
	expectedStack := []uint64{1, 0}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}

func TestSar(t *testing.T) {
	op := func(evm IEVM) error { return evm.Sar() }
	initialStack := []uint64{1, 2, 3}
	// Apply the operation: 3 >> 2 or 0b11 >> 0b10 = 0b0
	// Stack: [1, 2, 3] -> [1, 0]
	expectedStack := []uint64{1, 0}
	testStackOperation(t, op, nil, initialStack, expectedStack, nil)
}
