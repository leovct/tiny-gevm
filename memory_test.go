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
	} else {
		if len(m.data) != 0 {
			t.Errorf("NewMemory() created a memory with %d elements, want 0", len(m.data))
		}
	}
}

func TestStore(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store an element to the memory at offset 0.
	// The memory should be equal to [0x01, 0x02, 0x03].
	data := []byte{0x01, 0x02, 0x03}
	offset := 0
	initialMemory := data[:]
	m.Store(data, offset)

	if !bytes.Equal(m.data, data) {
		t.Errorf("Store() stored %v at offset %d in %v and resulted in %v, want %v", data, offset, initialMemory, m.data, data)
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

	expectedMemory := []byte{0x01, 0x04, 0x03}
	if !bytes.Equal(m.data, expectedMemory) {
		t.Errorf("Store() stored %v at offset %d in %v and resulted in %v, want %v", data, offset, initialMemory, m.data, expectedMemory)
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

	expectedMemory := append([]byte{0x0, 0x0, 0x0}, data...)
	if !bytes.Equal(m.data, expectedMemory) {
		t.Errorf("Store() stored %v at offset %d in %v and resulted in %v, want %v", data, offset, initialMemory, m.data, expectedMemory)
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

	expectedMemory := []byte{0x00, 0x01, 0x02, 0x0a, 0x0b, 0x0c}
	if !bytes.Equal(m.data, expectedMemory) {
		t.Errorf("Store() stored %v at offset %d in %v and resulted in %v, want %v", data, offset, initialMemory, m.data, expectedMemory)
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

	expectedValue := []byte{0x01, 0x02, 0x03}
	if !bytes.Equal(value, expectedValue) {
		t.Errorf("Access() accessed element at offset %d with size %d in %v and resulted in %v, want %v", offset, size, m.data, value, expectedValue)
	}
}

func TestAccessEmptyMemory(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Access memory at offset 2 with size 3 in empty memory.
	offset := 2
	size := 3
	value := m.Access(offset, size)

	expectedValue := []byte{0x00, 0x00, 0x00}
	if !bytes.Equal(value, expectedValue) {
		t.Errorf("Access() accessed element at offset %d with size %d in empty memory %v and resulted in %v, want %v", offset, size, m.data, value, expectedValue)
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

	expectedValue := []byte{0x00, 0x00, 0x00}
	if !bytes.Equal(value, expectedValue) {
		t.Errorf("Access() accessed out-of-bonds element at offset %d with size %d in %v and resulted in %v, want %v", offset, size, m.data, value, expectedValue)
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

	expectedValue := []byte{0x01, 0x02, 0x03, 0x00, 0x00}
	if !bytes.Equal(value, expectedValue) {
		t.Errorf("Access() accessed element at offset %d with size %d in %v and resulted in %v, want %v", offset, size, m.data, value, expectedValue)
	}
}
