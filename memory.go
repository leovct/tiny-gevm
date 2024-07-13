package main

type Memory struct {
	data []byte
}

func NewMemory() *Memory {
	return &Memory{data: make([]byte, 0)}
}
