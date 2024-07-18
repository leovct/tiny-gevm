# 🧱 Tiny gEVM

A tiny [Ethereum Virtual Machine](https://ethereum.github.io/yellowpaper/paper.pdf) (EVM) implementation from scratch, written in Go.

The aim is to keep it simple, quick to implement and interesting to learn more about the EVM.

Inspired by [gevm](https://github.com/Jesserc/gevm) by [Jesserc](https://twitter.com/jesserc_).

## EVM

```go
type IEVM interface {
  //// Stack operations
  // Push an item to the stack.
  Push(*uint256.Int) error

  // Pop an item from the stack.
  Pop() (*uint256.Int, error)

  //// Math operations
  // Add the top two elements of the stack and push the result, x + y, back to the stack.
  Add() error

  // Multiply the top two elements of the stack and push the result, x * y, back to the stack.
  Mul() error

  // Subtract the top two elements of the stack and push the result, x - y, back to the stack.
  Sub() error

  // Perform the integer divison operation on the top two elements of the stack and push the result, x // y, back to the stack.
  Div() error

  // Perform the signed integer division operation (trunced) on the top two elements of the stack and push the result, x // y, back to the stack.
  SDiv() error

  // Perform the modulo remained operation on the top two elements of the stack and push the result, x % y, back to the stack.
  Mod() error

  // Perform the signed modulo remained operation on the top two elements of the stack and push the result, x % y, back to the stack.
  SMod() error

  // Perform the modulo addition operation on the top two elements of the stack and push the result, (x + y) % m, back to the stack.
  // The third top element of the stack is the integer denominator m.
  AddMod() error

  // Perform the modulo multiplication operation on the top two elements of the stack and push the result, (x * y) % m, back to the stack.
  // The third top element of the stack is the integer denominator N.
  MulMod() error

  // Perform the exponential operation on the top two elements of the stack and push the result, x ** y, back to the stack.
  Exp() error

  // Extend the length of two’s complement signed integer.
  // The first top element of the stack, b, represents the size in byte - 1 of the integer to sign extend.
  // The second top element of the stack, x, represents the integer value to sign extend.
  SignExtend() error
}

// EVM represents an Ethereum Virtual Machine.
type EVM struct {
	stack   IStack
	memory  IMemory
	storage IStorage
}
```

## Stack

```go
type IStack interface {
  // Push adds a new element to the top of the stack.
  // It returns an error if the stack is full.
  Push(*uint256.Int) error

  // Pop removes and returns the top element from the stack.
  // If the stack is empty, it returns a zero-value 32-byte array and an error.
  Pop() (*uint256.Int, error)
}

// Stack represents a last-in-first-out (LIFO) stack of 32-byte arrays.
type Stack struct {
  data []uint256.Int
}
```

## Memory

```go
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
