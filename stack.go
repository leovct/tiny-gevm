package main

type Stack struct {
	data [][32]byte
}

func NewStack() *Stack {
	return &Stack{data: make([][32]byte, 0)}
}
