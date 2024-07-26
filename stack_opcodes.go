package main

import (
	"errors"

	"github.com/holiman/uint256"
)

var (
	// ErrPushSize is returned when the push size is outside the valid range of 1 to 32.
	ErrInvalidPushSize = errors.New("invalid push size")
	// ErrPushSizeExceedsCodeSize is returned when trying the push size exceeds the code size.
	ErrPushSizeExceedsCodeSize = errors.New("push size exceeds code size")

	// ErrDupSize is returned when the dup size is outside the valid range of 1 to 16.
	ErrInvalidDupSize = errors.New("invalid dup size")

	// ErrInvalidSwapSize is returned when the swap size is outside the valid range of 1 to 16.
	ErrInvalidSwapSize = errors.New("invalid swap size")
)

// IStackOps defines operations on the EVM stack.
// All operations take their values from code and push the result to the stack.
// All methods return an error if there are not enough elements in code or two many elements in the stack.
type IStackOps interface {
	// Pop an item from the stack.
	Pop() error

	IPushOps
	IDupOps
	ISwapOps
}

func (e *EVM) Pop() error {
	_, err := e.stack.Pop()
	return err
}

// IPushOps defines push operations on the EVM stack.
type IPushOps interface {
	// Push0 places the value zero on the stack.
	Push0() error

	// Push1 places a 1-byte item on the stack.
	Push1() error

	// Push2 places a 2-byte item on the stack.
	Push2() error

	// Push3 places a 3-byte item on the stack.
	Push3() error

	// Push4 places a 4-byte item on the stack.
	Push4() error

	// Push5 places a 5-byte item on the stack.
	Push5() error

	// Push6 places a 6-byte item on the stack.
	Push6() error

	// Push7 places a 7-byte item on the stack.
	Push7() error

	// Push8 places an 8-byte item on the stack.
	Push8() error

	// Push9 places a 9-byte item on the stack.
	Push9() error

	// Push10 places a 10-byte item on the stack.
	Push10() error

	// Push11 places an 11-byte item on the stack.
	Push11() error

	// Push12 places a 12-byte item on the stack.
	Push12() error

	// Push13 places a 13-byte item on the stack.
	Push13() error

	// Push14 places a 14-byte item on the stack.
	Push14() error

	// Push15 places a 15-byte item on the stack.
	Push15() error

	// Push16 places a 16-byte item on the stack.
	Push16() error

	// Push17 places a 17-byte item on the stack.
	Push17() error

	// Push18 places an 18-byte item on the stack.
	Push18() error

	// Push19 places a 19-byte item on the stack.
	Push19() error

	// Push20 places a 20-byte item on the stack.
	Push20() error

	// Push21 places a 21-byte item on the stack.
	Push21() error

	// Push22 places a 22-byte item on the stack.
	Push22() error

	// Push23 places a 23-byte item on the stack.
	Push23() error

	// Push24 places a 24-byte item on the stack.
	Push24() error

	// Push25 places a 25-byte item on the stack.
	Push25() error

	// Push26 places a 26-byte item on the stack.
	Push26() error

	// Push27 places a 27-byte item on the stack.
	Push27() error

	// Push28 places a 28-byte item on the stack.
	Push28() error

	// Push29 places a 29-byte item on the stack.
	Push29() error

	// Push30 places a 30-byte item on the stack.
	Push30() error

	// Push31 places a 31-byte item on the stack.
	Push31() error

	// Push32 places a 32-byte item on the stack.
	Push32() error
}

func (e *EVM) Push0() error {
	return e.stack.Push(uint256.NewInt(0))
}

func (e *EVM) Push1() error {
	return e.pushN(1)
}

func (e *EVM) Push2() error {
	return e.pushN(2)
}

func (e *EVM) Push3() error {
	return e.pushN(3)
}

func (e *EVM) Push4() error {
	return e.pushN(4)
}

func (e *EVM) Push5() error {
	return e.pushN(5)
}

func (e *EVM) Push6() error {
	return e.pushN(6)
}

func (e *EVM) Push7() error {
	return e.pushN(7)
}

func (e *EVM) Push8() error {
	return e.pushN(8)
}

func (e *EVM) Push9() error {
	return e.pushN(9)
}

func (e *EVM) Push10() error {
	return e.pushN(10)
}

func (e *EVM) Push11() error {
	return e.pushN(11)
}

func (e *EVM) Push12() error {
	return e.pushN(12)
}

func (e *EVM) Push13() error {
	return e.pushN(13)
}

func (e *EVM) Push14() error {
	return e.pushN(14)
}

func (e *EVM) Push15() error {
	return e.pushN(15)
}

func (e *EVM) Push16() error {
	return e.pushN(16)
}

func (e *EVM) Push17() error {
	return e.pushN(17)
}

func (e *EVM) Push18() error {
	return e.pushN(18)
}

func (e *EVM) Push19() error {
	return e.pushN(19)
}

func (e *EVM) Push20() error {
	return e.pushN(20)
}

func (e *EVM) Push21() error {
	return e.pushN(21)
}

func (e *EVM) Push22() error {
	return e.pushN(22)
}

func (e *EVM) Push23() error {
	return e.pushN(23)
}

func (e *EVM) Push24() error {
	return e.pushN(24)
}

func (e *EVM) Push25() error {
	return e.pushN(25)
}

func (e *EVM) Push26() error {
	return e.pushN(26)
}

func (e *EVM) Push27() error {
	return e.pushN(27)
}

func (e *EVM) Push28() error {
	return e.pushN(28)
}

func (e *EVM) Push29() error {
	return e.pushN(29)
}

func (e *EVM) Push30() error {
	return e.pushN(30)
}

func (e *EVM) Push31() error {
	return e.pushN(31)
}

func (e *EVM) Push32() error {
	return e.pushN(32)
}

func (e *EVM) pushN(n int) error {
	if n < 1 || n > 32 {
		// Unreachable in theory.
		// This step should never fail because the EVM should only expose Push1() to Push32().
		// Note that the EVM exposes Push0() but the logic is different and does not rely on pushN().
		return ErrInvalidPushSize
	}
	if len(e.env.code) <= e.state.pc+n {
		return ErrPushSizeExceedsCodeSize
	}

	code := e.env.code[e.state.pc+1 : e.state.pc+1+n]
	value := new(uint256.Int).SetBytes(code)
	if err := e.stack.Push(value); err != nil {
		return err
	}
	e.state.pc += n + 1
	return nil
}

// IDupOps defines duplication operation on the EVM stack.
type IDupOps interface {
	// Duplicate 1st stack item.
	Dup1() error

	// Duplicate 2nd stack item.
	Dup2() error

	// Duplicate 3rd stack item.
	Dup3() error

	// Duplicate 4th stack item.
	Dup4() error

	// Duplicate 5th stack item.
	Dup5() error

	// Duplicate 6th stack item.
	Dup6() error

	// Duplicate 7th stack item.
	Dup7() error

	// Duplicate 8th stack item.
	Dup8() error

	// Duplicate 9th stack item.
	Dup9() error

	// Duplicate 10th stack item.
	Dup10() error

	// Duplicate 11th stack item.
	Dup11() error

	// Duplicate 12th stack item.
	Dup12() error

	// Duplicate 13th stack item.
	Dup13() error

	// Duplicate 14th stack item.
	Dup14() error

	// Duplicate 15th stack item.
	Dup15() error

	// Duplicate 16th stack item.
	Dup16() error
}

func (e *EVM) Dup1() error {
	return e.dupN(1)
}

func (e *EVM) Dup2() error {
	return e.dupN(2)
}

func (e *EVM) Dup3() error {
	return e.dupN(3)
}

func (e *EVM) Dup4() error {
	return e.dupN(4)
}

func (e *EVM) Dup5() error {
	return e.dupN(5)
}

func (e *EVM) Dup6() error {
	return e.dupN(6)
}

func (e *EVM) Dup7() error {
	return e.dupN(7)
}

func (e *EVM) Dup8() error {
	return e.dupN(8)
}

func (e *EVM) Dup9() error {
	return e.dupN(9)
}

func (e *EVM) Dup10() error {
	return e.dupN(10)
}

func (e *EVM) Dup11() error {
	return e.dupN(11)
}

func (e *EVM) Dup12() error {
	return e.dupN(12)
}

func (e *EVM) Dup13() error {
	return e.dupN(13)
}

func (e *EVM) Dup14() error {
	return e.dupN(14)
}

func (e *EVM) Dup15() error {
	return e.dupN(15)
}

func (e *EVM) Dup16() error {
	return e.dupN(16)
}

func (e *EVM) dupN(n int) error {
	if n < 1 || n > 16 {
		// Unreachable in theory.
		// This step should never fail because the EVM should only expose Dup1() to Dup16().
		return ErrInvalidDupSize
	}

	value, err := e.stack.Get(n)
	if err != nil {
		return err
	}
	if err := e.stack.Push(value); err != nil {
		return err
	}
	return nil
}

// ISwapOps defines swap operation on the EVM stack.
type ISwapOps interface {
	// Swaplicate 1st stack item.
	Swap1() error

	// Swaplicate 2nd stack item.
	Swap2() error

	// Swaplicate 3rd stack item.
	Swap3() error

	// Swaplicate 4th stack item.
	Swap4() error

	// Swaplicate 5th stack item.
	Swap5() error

	// Swaplicate 6th stack item.
	Swap6() error

	// Swaplicate 7th stack item.
	Swap7() error

	// Swaplicate 8th stack item.
	Swap8() error

	// Swaplicate 9th stack item.
	Swap9() error

	// Swaplicate 10th stack item.
	Swap10() error

	// Swaplicate 11th stack item.
	Swap11() error

	// Swaplicate 12th stack item.
	Swap12() error

	// Swaplicate 13th stack item.
	Swap13() error

	// Swaplicate 14th stack item.
	Swap14() error

	// Swaplicate 15th stack item.
	Swap15() error

	// Swaplicate 16th stack item.
	Swap16() error
}

func (e *EVM) Swap1() error {
	return e.swapN(2)
}

func (e *EVM) Swap2() error {
	return e.swapN(3)
}

func (e *EVM) Swap3() error {
	return e.swapN(4)
}

func (e *EVM) Swap4() error {
	return e.swapN(5)
}

func (e *EVM) Swap5() error {
	return e.swapN(6)
}

func (e *EVM) Swap6() error {
	return e.swapN(7)
}

func (e *EVM) Swap7() error {
	return e.swapN(8)
}

func (e *EVM) Swap8() error {
	return e.swapN(9)
}

func (e *EVM) Swap9() error {
	return e.swapN(10)
}

func (e *EVM) Swap10() error {
	return e.swapN(11)
}

func (e *EVM) Swap11() error {
	return e.swapN(12)
}

func (e *EVM) Swap12() error {
	return e.swapN(13)
}

func (e *EVM) Swap13() error {
	return e.swapN(14)
}

func (e *EVM) Swap14() error {
	return e.swapN(15)
}

func (e *EVM) Swap15() error {
	return e.swapN(16)
}

func (e *EVM) Swap16() error {
	return e.swapN(17)
}

func (e *EVM) swapN(n int) error {
	if n < 2 || n > 17 {
		// Unreachable in theory.
		// This step should never fail because the EVM should only expose Swap1() to Swap16().
		return ErrInvalidSwapSize
	}
	return e.stack.Swap(n)
}
