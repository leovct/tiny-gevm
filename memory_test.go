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
	// The memory should be equal to [0x01, 0x02, 0x03].
	data := []byte{0x01, 0x02, 0x03}
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
	// The memory should be equal to [0x01, 0x02, 0x03].
	m.Store([]byte{0x01, 0x02, 0x03}, 0)

	// Store another element to the memory which should overwrite partially the previous element.
	// The memory should be equal to [0x01, 0x04, 0x03].
	data := []byte{0x04}
	offset := 1
	initialMemory := data[:]
	m.Store(data, offset)

	// Access the memory.
	memory := m.Access(0, 3)
	expectedMemory := []byte{0x01, 0x04, 0x03}
	if !bytes.Equal(memory, expectedMemory) {
		t.Errorf("Store() stored %v at offset %d in %v and resulted in %v, want %v", data, offset, initialMemory, memory, expectedMemory)
	}
}

func TestStoreAndExpandMemory(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store an element to the memory at offset 3. The memory should be expanded.
	// The memory should be equal to [0x00, 0x00, 0x00, 0x01, 0x02, 0x03].
	data := []byte{0x01, 0x02, 0x03}
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
	// The memory should be equal to [0x00, 0x01, 0x02, 0x03, 0x04, 0x05].
	m.Store([]byte{0x01, 0x02, 0x03, 0x04, 0x05}, 1)

	// Store an element to the memory at offset 3.
	// It should overwrite partially the previous element and the memory should also be expanded.
	// The memory should be equal to [0x00, 0x01, 0x02, 0x0a, 0x0b, 0x0c].
	data := []byte{0x0a, 0x0b, 0x0c}
	offset := 3
	initialMemory := data[:]
	m.Store(data, offset)

	// Access the memory.
	memory := m.Access(0, 6)
	expectedMemory := []byte{0x00, 0x01, 0x02, 0x0a, 0x0b, 0x0c}
	if !bytes.Equal(memory, expectedMemory) {
		t.Errorf("Store() stored %v at offset %d in %v and resulted in %v, want %v", data, offset, initialMemory, memory, expectedMemory)
	}
}

func TestAccessInBounds(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store an element to the memory at index 2.
	// The memory should be equal to [0x00, 0x00, 0x01, 0x02, 0x03].
	m.Store([]byte{0x01, 0x02, 0x03}, 2)

	// Access memory at offset 2 with size 3.
	offset := 2
	size := 3
	value := m.Access(offset, size)

	// Access the memory.
	memory := m.Access(0, 5)
	expectedValue := []byte{0x01, 0x02, 0x03}
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
	expectedValue := []byte{0x00, 0x00, 0x00}
	if !bytes.Equal(value, expectedValue) {
		t.Errorf("Access() accessed element at offset %d with size %d in empty memory %v and resulted in %v, want %v", offset, size, memory, value, expectedValue)
	}
}

func TestAccess0utOfBonds(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store an element to the memory at index 2.
	// The memory should be equal to [0x00, 0x00, 0x01, 0x02, 0x03].
	m.Store([]byte{0x01, 0x02, 0x03}, 2)

	// Access memory at offset 10 with size 3. It is out of bonds.
	offset := 10
	size := 3
	value := m.Access(offset, size)

	// Access the memory.
	memory := m.Access(0, 5)
	expectedValue := []byte{0x00, 0x00, 0x00}
	if !bytes.Equal(value, expectedValue) {
		t.Errorf("Access() accessed out-of-bonds element at offset %d with size %d in %v and resulted in %v, want %v", offset, size, memory, value, expectedValue)
	}
}

func TestAccessPartial0utOfBonds(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store an element to the memory at index 2.
	// The memory should be equal to [0x00, 0x00, 0x01, 0x02, 0x03].
	m.Store([]byte{0x01, 0x02, 0x03}, 2)

	// Access memory at offset 2 with size 5.
	offset := 2
	size := 5
	value := m.Access(offset, size)

	// Access the memory.
	memory := m.Access(0, 5)
	expectedValue := []byte{0x01, 0x02, 0x03, 0x00, 0x00}
	if !bytes.Equal(value, expectedValue) {
		t.Errorf("Access() accessed element at offset %d with size %d in %v and resulted in %v, want %v", offset, size, memory, value, expectedValue)
	}
}
