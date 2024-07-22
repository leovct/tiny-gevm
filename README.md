# 🧱 Tiny gEVM

A tiny [Ethereum Virtual Machine](https://ethereum.github.io/yellowpaper/paper.pdf) (EVM) implementation from scratch, written in Go.

The aim is to keep it simple, quick to implement and interesting to learn more about the EVM.

Inspired by [gevm](https://github.com/Jesserc/gevm) by [Jesserc](https://twitter.com/jesserc_).

## EVM

```go
// IEVM defines the methods that an Ethereum Virtual Machine implementation should have.
type IEVM interface {
	// Stack operations.
	// Push an item to the stack.
	Push(*uint256.Int) error
	// Pop an item from the stack.
	Pop() (*uint256.Int, error)

	// Memory operations.
	// Write byte slice to memory at the specified offset.
	Store(value []byte, offset int)

	// Arithmetic operations.
	IArithmeticOps

	// Comparison and bitwise logic operations.
	IComparisonAndBitwiseOps

	// SHAA3 operations.
	ISHA3Ops

	// Stack operations.
	IStackOps
}

// EVM represents an Ethereum Virtual Machine.
type EVM struct {
	stack	IStack
	memory	IMemory
	storage	IStorage
	env	ExecutionEnvironment
	state	MachineState
}

// ExecutionEnvironment represents the EVM execution environment.
type ExecutionEnvironment struct {
	// Machine code to be executed by the EVM.
	code []byte
}

// MachineState represents the EVM state.
type MachineState struct {
	// Program counter.
	pc int
}
```

## Stack

```go
// IStack defines the methods that a stack implementation should have.
type IStack interface {
	// Push adds a new element to the top of the stack.
	// It returns an error if the stack is full.
	Push(*uint256.Int) error

	// Pop removes and returns the top element from the stack.
	// If the stack is empty, it returns a zero-value 32-byte array and an error.
	Pop() (*uint256.Int, error)

	// Size returns the number of elements currently on the stack.
	Size() int
}

// Stack represents a last-in-first-out (LIFO) stack of 32-byte arrays.
type Stack struct {
	data []uint256.Int
}
```

## Memory

```go
// IMemory defines the methods that a memory implementation should have.
type IMemory interface {
	// Store writes a byte slice to memory at the specified offset.
	// If the offset plus the length of the value exceeds the current memory size,
	// the memory is automatically expanded to accommodate the new data.
	Store(value []byte, offset int)

	// Access retrieves a slice of memory starting at the given offset with the specified size.
	// It handles cases where the requested region may extend beyond the current memory size.
	// Returns a byte slice of length 'size', zero-padded if necessary.
	Access(offset, size int) []byte
}

// Memory represents a byte-addressable memory structure.
type Memory struct {
	data []byte
}
```

## Storage

```go
// IStorage defines the methods that a storage implementation should have.
type IStorage interface {
	// Store writes a 32-byte word to storage at the specified key.
	// If the key already exists, its value will be overwritten.
	Store(key int, value [32]byte)

	// Load retrieves a 32-byte word from storage using the specified key.
	// If the key does not exist in the storage, it returns an empty 32-byte word.
	Load(key int) [32]byte
}

// Storage represents a word-addressable storage structure.
type Storage struct {
	data map[int][32]byte
}
```

## References

- [evm.codes](https://www.evm.codes/)
