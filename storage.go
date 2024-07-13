package main

type Storage struct {
	data map[int][32]byte
}

func NewStorage() *Storage {
	return &Storage{data: make(map[int][32]byte)}
}
