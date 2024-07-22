package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

func TestKeccak256(t *testing.T) {
	op := func(evm IEVM) error { return evm.Keccak256() }

	offset := uint64(1)
	size := uint64(3)
	initialStack := []uint64{1, size, offset}
	memory := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	data := memory[offset : offset+size]

	hash := crypto.Keccak256(data)
	value := new(uint256.Int).SetBytes([]byte(hash)).Uint64()
	expectedStack := []uint64{1, value}
	testStackOperationWithNewEVM(t, op, nil, initialStack, expectedStack, memory, nil)
}
