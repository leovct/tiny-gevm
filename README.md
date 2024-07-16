# 🧱 Tiny gEVM

A tiny [Ethereum Virtual Machine](https://ethereum.github.io/yellowpaper/paper.pdf) (EVM) implementation from scratch, written in Go.

The aim is to keep it simple, quick to implement and interesting to learn more about the EVM.

Inspired by [gevm](https://github.com/Jesserc/gevm) by [Jesserc](https://twitter.com/jesserc_).

## Stack

```go
type IStack interface {
  Push(*uint256.Int) error
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
  Store(value []byte, offset int)
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
  Store(key int, value [32]byte)
  Load(key int) [32]byte
}

// Storage represents a word-addressable storage structure.
type Storage struct {
  data map[int][32]byte
}
```
