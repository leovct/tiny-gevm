package evm

import (
	"github.com/holiman/uint256"
)

// IPushOps defines operations on the EVM memory.
type IMemoryOps interface {
	// MLoad loads a word from memory.
	// It pops an item from the stack, this is the offset.
	// Then it reads the word from memory at the given offset.
	// Finally, it pushes the result to the top of the stack.
	// Stack: [offset, ...] -> [word, ...]
	// Memory: [offset:offset+32] = word
	MLoad() error

	// MStore saves a word to memory.
	// It pops two items from the stack, offset and value (32-bytes).
	// Then it writes the word at the given offset in the memory.
	// Stack: [offset, word, ...] -> [...]
	// Memory: [offset:offset+32] = ??? -> [offset:offset+32] = value
	MStore() error

	// MStore8 saves a byte to memory.
	// It pops two items from the stack, offset and value (1-byte).
	// Then it writes the value at the given offset in the memory.
	// Stack: [offset, value, ...] -> [...]
	// Memory: [offset:offset+8] = ??? -> [offset:offset+8] = value
	MStore8() error
}

func (e *EVM) MLoad() error {
	// Load offset from stack.
	offset, err := e.stack.Pop()
	if err != nil {
		return err
	}

	// Load memory from memory at given offset.
	word := e.memory.LoadWord(int(offset.Uint64()))

	// Store word at the top of the stack.
	value := new(uint256.Int).SetBytes(word[:])
	return e.stack.Push(value)
}

func (e *EVM) MStore() error {
	// Load offset from the stack.
	offset, err := e.stack.Pop()
	if err != nil {
		return err
	}

	// Load word from the stack.
	var value *uint256.Int
	value, err = e.stack.Pop()
	if err != nil {
		return err
	}

	// Store word at the given offset in memory.
	word := value.Bytes32()
	e.memory.StoreWord(word, int(offset.Uint64()))
	return nil
}

func (e *EVM) MStore8() error {
	// Load offset from the stack.
	offset, err := e.stack.Pop()
	if err != nil {
		return err
	}

	// Load value from the stack.
	var value *uint256.Int
	value, err = e.stack.Pop()
	if err != nil {
		return err
	}

	// Store byte at the given offset in memory.
	word := value.Bytes()
	// TODO: Should we check the size of word?
	e.memory.StoreByte(word[0], int(offset.Uint64()))
	return nil
}
