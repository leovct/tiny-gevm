package main

import (
	"testing"

	"github.com/holiman/uint256"
)

func TestMLoad(t *testing.T) {
	op := func(evm IEVM) error { return evm.MLoad() }

	// Stack
	offset := 32 // read the second word
	initialStack := []uint64{3, 2, uint64(offset)}

	// Memory
	word1 := uint256.NewInt(333).Bytes32()
	word2 := uint256.NewInt(222).Bytes32()
	word3 := uint256.NewInt(111).Bytes32()
	var memory []byte
	memory = append(append(append(memory, word1[:]...), word2[:]...), word3[:]...)

	expectedStack := []uint64{3, 2, 222}
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, memory, nil, nil)
}

func TestMLoadOnEmptyStack(t *testing.T) {
	op := func(evm IEVM) error { return evm.MLoad() }
	testStackOperationWithNewEVM(t, op, ErrStackUnderflow, nil, nil, nil, nil, nil)
}

func TestMStore(t *testing.T) {
	op := func(evm IEVM) error { return evm.MStore() }

	// Initial Stack
	// [32, 444, 2, 3]
	offset := 32
	initialStack := []uint64{3, 2, 444, uint64(offset)}

	// Initial Memory
	// Each word is represented by 32 bytes.
	// [111, 222, 333]
	word1 := uint256.NewInt(111).Bytes32()
	word2 := uint256.NewInt(222).Bytes32()
	word3 := uint256.NewInt(333).Bytes32()
	var initialMemory []byte
	initialMemory = append(append(append(initialMemory, word1[:]...), word2[:]...), word3[:]...)

	// Expected Stack
	// [2, 3]
	expectedStack := []uint64{3, 2}

	// Expected Memory
	// Each word is represented by 32 bytes.
	// [111, 444, 333]
	word4 := uint256.NewInt(444).Bytes32()
	var expectedMemory []byte
	expectedMemory = append(append(append(expectedMemory, word1[:]...), word4[:]...), word3[:]...)

	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, initialMemory, expectedMemory, nil)
}

func TestMStoreOnEmptyStack(t *testing.T) {
	op := func(evm IEVM) error { return evm.MStore() }
	testStackOperationWithNewEVM(t, op, ErrStackUnderflow, nil, nil, nil, nil, nil)
}

func TestMStoreOnOneElementStack(t *testing.T) {
	op := func(evm IEVM) error { return evm.MStore() }
	initialStack := []uint64{1}
	testStackOperationWithNewEVM(t, op, ErrStackUnderflow, initialStack, nil, nil, nil, nil)
}

func TestMStore8(t *testing.T) {
	op := func(evm IEVM) error { return evm.MStore8() }

	// Initial Stack
	// [1, 444, 2, 3]
	offset := 1
	initialStack := []uint64{3, 2, 0x20, uint64(offset)}

	// Initial Memory
	// Each word is represented by 1 byte.
	// [0x01, 0x02, 0x03, 0x04]
	initialMemory := []byte{0x01, 0x02, 0x03, 0x04}

	// Expected Stack
	// [2, 3]
	expectedStack := []uint64{3, 2}

	// Expected Memory
	// Each word is represented by 1 byte.
	// [0x01, 0x20, 0x03, 0x04]
	expectedMemory := []byte{0x01, 0x20, 0x03, 0x04}

	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, initialMemory, expectedMemory, nil)
}

func TestMStore8OnEmptyStack(t *testing.T) {
	op := func(evm IEVM) error { return evm.MStore8() }
	testStackOperationWithNewEVM(t, op, ErrStackUnderflow, nil, nil, nil, nil, nil)
}

func TestMStore8OnOneElementStack(t *testing.T) {
	op := func(evm IEVM) error { return evm.MStore8() }
	initialStack := []uint64{1}
	testStackOperationWithNewEVM(t, op, ErrStackUnderflow, initialStack, nil, nil, nil, nil)
}
