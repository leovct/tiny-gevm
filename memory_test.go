package main

import (
	"bytes"
	"testing"
)

func TestNewMemory(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()
	if m == nil {
		t.Error("NewMemory() returned nil")
	}
}

func TestStoreAndAccess(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store an element to the memory at offset 0.
	// The memory should be equal to [0x1, 0x2, 0x3].
	data := []byte{0x1, 0x2, 0x3}
	offset := 0
	initialMemory := data[:]
	m.Store(data, offset)

	// Access the memory.
	memory := m.Access(0, 3)
	if !bytes.Equal(memory, data) {
		t.Errorf("Store() stored %v at offset %d in %v and resulted in %v, want %v", data, offset, initialMemory, memory, data)
	}
}

func TestStoreOverWrite(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store an element to the memory at offset 0.
	// The memory should be equal to [0x1, 0x2, 0x3].
	m.Store([]byte{0x1, 0x2, 0x3}, 0)

	// Store another element to the memory which should overwrite partially the previous element.
	// The memory should be equal to [0x1, 0x4, 0x3].
	data := []byte{0x4}
	offset := 1
	initialMemory := data[:]
	m.Store(data, offset)

	// Access the memory.
	memory := m.Access(0, 3)
	expectedMemory := []byte{0x1, 0x4, 0x3}
	if !bytes.Equal(memory, expectedMemory) {
		t.Errorf("Store() stored %v at offset %d in %v and resulted in %v, want %v", data, offset, initialMemory, memory, expectedMemory)
	}
}

func TestStoreAndExpandMemory(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store an element to the memory at offset 3. The memory should be expanded.
	// The memory should be equal to [0x0, 0x0, 0x0, 0x1, 0x2, 0x3].
	data := []byte{0x1, 0x2, 0x3}
	offset := 3
	initialMemory := data[:]
	m.Store(data, offset)

	// Access the memory.
	memory := m.Access(0, 6)
	expectedMemory := append([]byte{0x0, 0x0, 0x0}, data...)
	if !bytes.Equal(memory, expectedMemory) {
		t.Errorf("Store() stored %v at offset %d in %v and resulted in %v, want %v", data, offset, initialMemory, memory, expectedMemory)
	}
}

func TestStoreOverWriteAndExpandMemory(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store an element to the memory at offset 1.
	// The memory should be equal to [0x0, 0x1, 0x2, 0x3, 0x4, 0x5].
	m.Store([]byte{0x1, 0x2, 0x3, 0x4, 0x5}, 1)

	// Store an element to the memory at offset 3.
	// It should overwrite partially the previous element and the memory should also be expanded.
	// The memory should be equal to [0x0, 0x1, 0x2, 0xA, 0xB, 0xC].
	data := []byte{0xA, 0xB, 0xC}
	offset := 3
	initialMemory := data[:]
	m.Store(data, offset)

	// Access the memory.
	memory := m.Access(0, 6)
	expectedMemory := []byte{0x0, 0x1, 0x2, 0xA, 0xB, 0xC}
	if !bytes.Equal(memory, expectedMemory) {
		t.Errorf("Store() stored %v at offset %d in %v and resulted in %v, want %v", data, offset, initialMemory, memory, expectedMemory)
	}
}

func TestAccessInBounds(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store an element to the memory at index 2.
	// The memory should be equal to [0x0, 0x0, 0x1, 0x2, 0x3].
	m.Store([]byte{0x1, 0x2, 0x3}, 2)

	// Access memory at offset 2 with size 3.
	offset := 2
	size := 3
	value := m.Access(offset, size)

	// Access the memory.
	memory := m.Access(0, 5)
	expectedValue := []byte{0x1, 0x2, 0x3}
	if !bytes.Equal(value, expectedValue) {
		t.Errorf("Access() accessed element at offset %d with size %d in %v and resulted in %v, want %v", offset, size, memory, value, expectedValue)
	}
}

func TestAccessEmptyMemory(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Access memory at offset 2 with size 3 in empty memory.
	offset := 2
	size := 3
	value := m.Access(offset, size)

	// Access the memory.
	memory := m.Access(0, 5)
	expectedValue := []byte{0x0, 0x0, 0x0}
	if !bytes.Equal(value, expectedValue) {
		t.Errorf("Access() accessed element at offset %d with size %d in empty memory %v and resulted in %v, want %v", offset, size, memory, value, expectedValue)
	}
}

func TestAccess0utOfBonds(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store an element to the memory at index 2.
	// The memory should be equal to [0x0, 0x0, 0x1, 0x2, 0x3].
	m.Store([]byte{0x1, 0x2, 0x3}, 2)

	// Access memory at offset 10 with size 3. It is out of bonds.
	offset := 10
	size := 3
	value := m.Access(offset, size)

	// Access the memory.
	memory := m.Access(0, 5)
	expectedValue := []byte{0x0, 0x0, 0x0}
	if !bytes.Equal(value, expectedValue) {
		t.Errorf("Access() accessed out-of-bonds element at offset %d with size %d in %v and resulted in %v, want %v", offset, size, memory, value, expectedValue)
	}
}

func TestAccessPartial0utOfBonds(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store an element to the memory at index 2.
	// The memory should be equal to [0x0, 0x0, 0x1, 0x2, 0x3].
	m.Store([]byte{0x1, 0x2, 0x3}, 2)

	// Access memory at offset 2 with size 5.
	offset := 2
	size := 5
	value := m.Access(offset, size)

	// Access the memory.
	memory := m.Access(0, 5)
	expectedValue := []byte{0x1, 0x2, 0x3, 0x0, 0x0}
	if !bytes.Equal(value, expectedValue) {
		t.Errorf("Access() accessed element at offset %d with size %d in %v and resulted in %v, want %v", offset, size, memory, value, expectedValue)
	}
}
