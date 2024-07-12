package main

import "fmt"

func main() {
	evm := NewEVM()
	fmt.Println("EVM:", evm)
}

type EVM struct {
	*Stack
	*Memory
	*Storage
}

func NewEVM() *EVM {
	return &EVM{
		Stack:   NewStack(),
		Memory:  NewMemory(),
		Storage: NewStorage(),
	}
}

type Stack struct {
	data [][32]byte
}

func NewStack() *Stack {
	return &Stack{data: make([][32]byte, 0)}
}

type Memory struct {
	data []byte
}

func NewMemory() *Memory {
	return &Memory{data: make([]byte, 0)}
}

type Storage struct {
	data map[int][32]byte
}

func NewStorage() *Storage {
	return &Storage{data: make(map[int][32]byte)}
}
