package main

// IMemory defines the methods that a memory implementation should have.
type IMemory interface {
	// Store writes a byte slice to memory at the specified offset.
	// If the offset plus the length of the value exceeds the current memory size,
	// the memory is automatically expanded to accommodate the new data.
	Store(value []byte, offset int)

	// Access retrieves a slice of memory starting at the given offset with the specified size.
	// It handles cases where the requested region may extend beyond the current memory size.
	// Returns a byte slice of length 'size', zero-padded if necessary.
	Access(offset, size int) []byte

	// Load a word (32 bytes) from memory at the given offset.
	Load(offset int) []byte
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

func (m *Memory) Access(offset, size int) []byte {
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

func (m *Memory) Load(offset int) []byte {
	return m.Access(offset, 32)
}
