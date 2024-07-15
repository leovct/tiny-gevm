# 🧱 Tiny gEVM

A tiny [Ethereum Virtual Machine](https://ethereum.github.io/yellowpaper/paper.pdf) (EVM) implementation from scratch, written in Go.

The aim is to keep it simple, quick to implement and interesting to learn more about the EVM.

Inspired by [gevm](https://github.com/Jesserc/gevm) by [Jesserc](https://twitter.com/jesserc_).

## Stack

```go
// IStack defines the methods that a stack implementation should have.
type IStack interface {
  Push([32]byte) error
  Pop() ([32]byte, error)
}

// Stack represents a last-in-first-out (LIFO) stack of 32-byte arrays.
type Stack struct {
  data [][32]byte
}
```

## Memory

```go
// IMemory defines the methods that a memory implementation should have.
type IMemory interface {
	Store(value []byte, offset int)
	Access(offset, size int) []byte
}

// Memory represents a byte-addressable memory structure.
type Memory struct {
	data []byte
}
```
