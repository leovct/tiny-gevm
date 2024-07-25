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
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, memory, nil)
}

func TestMStore(t *testing.T) {
	op := func(evm IEVM) error { return evm.MStore() }

	// Stack
	// [32, 444, 2, 3]
	offset := 32
	initialStack := []uint64{3, 2, 444, uint64(offset)}

	// Memory
	// Each word is represented by 32 bytes.
	// [111, 222, 333]
	word1 := uint256.NewInt(111).Bytes32()
	word2 := uint256.NewInt(222).Bytes32()
	word3 := uint256.NewInt(333).Bytes32()
	var memory []byte
	memory = append(append(append(memory, word1[:]...), word2[:]...), word3[:]...)

	// Expected
	// - Stack: [2, 3]
	// - Memory: [111, 444, 333]
	expectedStack := []uint64{3, 2}
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, memory, nil)
}
