package main

import "fmt"

func main() {
	evm := NewEVM()
	fmt.Println("EVM:", evm)
}

type EVM struct {
	Stack IStack
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
