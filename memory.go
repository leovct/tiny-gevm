package main

// IMemory defines the methods that a memory implementation should have.
type IMemory interface {
	// Store writes a byte slice to memory at the specified offset.
	// If the offset plus the length of the value exceeds the current memory size,
	// the memory is automatically expanded to accommodate the new data.
	Store(value []byte, offset int)

	// Load retrieves a slice of memory starting at the given offset with the specified size.
	// It handles cases where the requested region may extend beyond the current memory size.
	// Returns a byte slice of length 'size', zero-padded if necessary.
	Load(offset, size int) []byte

	// Load a byte from memory at the given offset.
	LoadByte(offset int) byte

	// Load a word (32 bytes) from memory at the given offset.
	LoadWord(offset int) [32]byte

	// Store a byte to memory at the given offset.
	StoreByte(value byte, offset int)

	// Store a word (32 bytes) to memory at the given offset.
	StoreWord(word [32]byte, offset int)
}

// Memory represents a byte-addressable memory structure.
type Memory struct {
	data []byte
}

// NewMemory creates and returns a new, empty Memory instance.
func NewMemory() IMemory {
	return &Memory{data: make([]byte, 0)}
}

func (m *Memory) Store(value []byte, offset int) {
	// Expand the memory if needed.
	requiredSize := offset + len(value)
	if len(m.data) < requiredSize {
		m.data = append(m.data, make([]byte, requiredSize-len(m.data))...)
	}

	// Copy the value into memory at the specified offset.
	copy(m.data[offset:], value)
}

func (m *Memory) Load(offset, size int) []byte {
	// Return a zero-filled slice if memory is empty or offset is out of bounds.
	if len(m.data) == 0 || offset >= len(m.data) {
		return make([]byte, size)
	}

	// Handle partial out-of-bounds access.
	end := offset + size
	if end > len(m.data) {
		result := make([]byte, size)
		copy(result, m.data[offset:])
		return result
	}

	// Handle partial out-of-bounds access.
	return m.data[offset : offset+size]
}

func (m *Memory) LoadByte(offset int) byte {
	array := m.Load(offset, 1)
	return array[0]
}

func (m *Memory) LoadWord(offset int) [32]byte {
	array := m.Load(offset, 32)
	var word [32]byte
	copy(word[:], array)
	return word
}

func (m *Memory) StoreByte(value byte, offset int) {
	m.Store([]byte{value}, offset)
}

func (m *Memory) StoreWord(word [32]byte, offset int) {
	m.Store(word[:], offset)
}
