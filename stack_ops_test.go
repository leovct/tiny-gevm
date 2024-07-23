package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/holiman/uint256"
)

func TestPush0(t *testing.T) {
	op := func(evm IEVM) error { return evm.Push0() }
	initialStack := []uint64{1, 2, 3}
	expectedStack := []uint64{1, 2, 3, 0}
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil)
}

func TestPushN(t *testing.T) {
	initialStack := []uint64{1, 2, 3}
	for i := 1; i <= 32; i++ {
		t.Run(fmt.Sprintf("Push%d", i), func(t *testing.T) {
			testPushOperation(t, i, initialStack)
		})
	}
}

func TestPushSizeExceedsCodeSize(t *testing.T) {
	// Create a new EVM.
	evm := NewEVM(nil)

	// Try to push the first element from code (which is empty).
	if err := evm.Push1(); err == nil {
		t.Errorf("Operation returned an unexpected error: %v, wanted: %v", err, ErrPushSizeExceedsCodeSize)
	}
}

func TestPushOnFullStack(t *testing.T) {
	// Create a new EVM.
	code := []byte{0x60, 0x11}
	evm := NewEVM(code)

	// Push elements to the stack until its full.
	for i := 0; i < 1024; i++ {
		if err := evm.Push(uint256.NewInt(uint64(i))); err != nil {
			t.Errorf("Push() returned an unexpected error at iteration %d: %v", i, err)
			break
		}
	}

	// Try to push the first element (0x11) from code.
	if err := evm.Push1(); err == nil {
		t.Errorf("Operation returned an unexpected error: %v, wanted: %v", err, ErrStackOverflow)
	}
}

func testPushOperation(t *testing.T, pushSize int, initialStack []uint64) {
	// Generate value1 and value2 based on pushSize.
	value1 := make([]byte, pushSize)
	value2 := make([]byte, pushSize)
	for i := 0; i < pushSize; i++ {
		value1[i] = 0x11
		value2[i] = 0xAA
	}

	// Create a new EVM.
	code := append([]byte{0x60}, append(value1, append([]byte{0x60}, value2...)...)...)
	evm := NewEVM(code)

	// Generate the PUSH opcode function dynamically.
	op := func(evm IEVM) error {
		methodName := fmt.Sprintf("Push%d", pushSize)
		method := reflect.ValueOf(evm).MethodByName(methodName)
		if !method.IsValid() {
			return fmt.Errorf("method %s not found", methodName)
		}
		results := method.Call(nil)
		if len(results) > 0 && !results[0].IsNil() {
			return results[0].Interface().(error)
		}
		return nil
	}

	// Push a first value to the stack.
	value1ToUint64 := new(uint256.Int).SetBytes(value1).Uint64()
	expectedStack := append(initialStack, value1ToUint64)
	testStackOperationWithExistingEVM(t, evm, op, nil, initialStack, expectedStack, nil)

	// Push a second value to the stack.
	initialStack = append(initialStack, value1ToUint64)
	value2ToUint64 := new(uint256.Int).SetBytes(value2).Uint64()
	expectedStack = append(initialStack, value2ToUint64)
	testStackOperationWithExistingEVM(t, evm, op, nil, initialStack, expectedStack, nil)
}

func TestDupN(t *testing.T) {
	initialStack := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	for i := 1; i <= 16; i++ {
		t.Run(fmt.Sprintf("Dup%d", i), func(t *testing.T) {
			testDupOperation(t, i, initialStack)
		})
	}
}

func TestDupOnEmptyStack(t *testing.T) {
	op := func(evm IEVM) error { return evm.Dup1() }
	initialStack := []uint64{}
	expectedStack := []uint64{}
	testStackOperationWithNewEVM(t, op, ErrEmptyStack, initialStack, expectedStack, nil, nil)
}

func TestDupOnFullStack(t *testing.T) {
	// Create a new EVM.
	evm := NewEVM(nil)

	// Push elements to the stack until its full.
	for i := 0; i < 1024; i++ {
		if err := evm.Push(uint256.NewInt(uint64(i))); err != nil {
			t.Errorf("Push() returned an unexpected error at iteration %d: %v", i, err)
			break
		}
	}

	// Try to duplicate the first element.
	if err := evm.Dup1(); err == nil {
		t.Errorf("Operation returned an unexpected error: %v, wanted: %v", err, ErrStackOverflow)
	}
}

func testDupOperation(t *testing.T, dupSize int, initialStack []uint64) {
	// Generate the DUP opcode function dynamically.
	op := func(evm IEVM) error {
		methodName := fmt.Sprintf("Dup%d", dupSize)
		method := reflect.ValueOf(evm).MethodByName(methodName)
		if !method.IsValid() {
			return fmt.Errorf("method %s not found", methodName)
		}
		results := method.Call(nil)
		if len(results) > 0 && !results[0].IsNil() {
			return results[0].Interface().(error)
		}
		return nil
	}

	// Get the element to duplicate.
	index := len(initialStack) - dupSize
	value := initialStack[index]

	// Compute the expected stack.
	expectedStack := append(initialStack, value)

	// Test the dup operation.
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil)
}

func TestSwapN(t *testing.T) {
	initialStack := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	for i := 1; i <= 16; i++ {
		t.Run(fmt.Sprintf("Swap%d", i), func(t *testing.T) {
			testSwapOperation(t, i, initialStack)
		})
	}
}

func testSwapOperation(t *testing.T, dupSize int, initialStack []uint64) {
	// Generate the SWAP opcode function dynamically.
	op := func(evm IEVM) error {
		methodName := fmt.Sprintf("Swap%d", dupSize)
		method := reflect.ValueOf(evm).MethodByName(methodName)
		if !method.IsValid() {
			return fmt.Errorf("method %s not found", methodName)
		}
		results := method.Call(nil)
		if len(results) > 0 && !results[0].IsNil() {
			return results[0].Interface().(error)
		}
		return nil
	}

	// Swap the elements.
	expectedStack := make([]uint64, len(initialStack))
	copy(expectedStack, initialStack)
	index := len(initialStack) - (dupSize + 1)
	expectedStack[len(initialStack)-1], expectedStack[index] = expectedStack[index], expectedStack[len(initialStack)-1]

	// Test the dup operation.
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, nil, nil)
}
