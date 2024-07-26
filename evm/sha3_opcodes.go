package evm

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

// ISHA3Ops defines SHA3 operations for the EVM.
// All operations pop their operands from the stack and push the result back.
// All methods return an error if there are not enough elements on the stack.
type ISHA3Ops interface {
	// Compute Keccak-256 hash of the given data in memory.
	// Stack: [offset, size, ...] -> [hash, ...]
	Keccak256() error
}

func (e *EVM) Keccak256() error {
	op := func(operands ...*uint256.Int) *uint256.Int {
		offset, size := operands[0], operands[1]
		data := e.memory.Load(int(offset.Uint64()), int(size.Uint64()))
		hash := crypto.Keccak256(data)
		return new(uint256.Int).SetBytes(hash)
	}
	return e.performBinaryStackOperation(2, op)
}
