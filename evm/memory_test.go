package evm

import (
	"bytes"
	"testing"

	"github.com/holiman/uint256"
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
	memory := m.Load(0, 3)
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
	memory := m.Load(0, 3)
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
	memory := m.Load(0, 6)
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
	memory := m.Load(0, 6)
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
	value := m.Load(offset, size)

	// Access the memory.
	memory := m.Load(0, 5)
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
	value := m.Load(offset, size)

	// Access the memory.
	memory := m.Load(0, 5)
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
	value := m.Load(offset, size)

	// Access the memory.
	memory := m.Load(0, 5)
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
	value := m.Load(offset, size)

	// Access the memory.
	memory := m.Load(0, 5)
	expectedValue := []byte{0x1, 0x2, 0x3, 0x0, 0x0}
	if !bytes.Equal(value, expectedValue) {
		t.Errorf("Access() accessed element at offset %d with size %d in %v and resulted in %v, want %v", offset, size, memory, value, expectedValue)
	}
}

func TestLoadByte(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store elements to the memory.
	// The memory should contain 34 bytes.
	data := []byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, // 10 elements
		0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x11, 0x12, 0x13, 0x14,
		0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E,
		0x1F, 0x20, 0x21, 0x22,
	}
	m.Store(data, 0)

	// Load the first byte from memory.
	byte1 := m.LoadByte(0)
	expectedByte1 := byte(0x01)
	if byte1 != expectedByte1 {
		t.Errorf("LoadByte() at offset 0 returned byte %v, wanted %v", byte1, expectedByte1)
	}

	// Load the second byte from memory.
	byte2 := m.LoadByte(1)
	expectedByte2 := byte(0x02)
	if byte2 != expectedByte2 {
		t.Errorf("LoadByte() at offset 1 returned byte %v, wanted %v", byte2, expectedByte2)
	}

	// Load the 33th byte from memory.
	byte33 := m.LoadByte(33)
	expectedByte33 := byte(0x22)
	if byte33 != expectedByte33 {
		t.Errorf("LoadByte() at offset 33 returned byte %v, wanted %v", byte33, expectedByte33)
	}

	// Load the 34th byte from memory.
	byte34 := m.LoadByte(34)
	expectedByte34 := byte(0x00)
	if byte34 != expectedByte34 {
		t.Errorf("LoadByte() at offset 34 returned byte %v, wanted %v", byte34, expectedByte34)
	}
}

func TestLoadWord(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store elements to the memory.
	// The memory should contain 34 bytes.
	data := []byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, // 10 elements
		0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x11, 0x12, 0x13, 0x14,
		0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E,
		0x1F, 0x20, 0x21, 0x22,
	}
	m.Store(data, 0)

	// Load the first word (32 bytes) from memory.
	word1 := m.LoadWord(0)
	expectedWord1 := []byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, // 10 elements
		0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x11, 0x12, 0x13, 0x14,
		0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E,
		0x1F, 0x20,
	}
	if !bytes.Equal(word1[:], expectedWord1) {
		t.Errorf("LoadWord() at offset 0 returned word %v, wanted %v", word1, expectedWord1)
	}

	// Load the second word (32 bytes) from memory.
	word2 := m.LoadWord(32)
	expectedWord2 := []byte{
		0x21, 0x22, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
	}
	if !bytes.Equal(word2[:], expectedWord2) {
		t.Errorf("LoadWord() at offset 32 returned word %v, wanted %v", word2, expectedWord2)
	}
}

func TestStoreByte(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store elements to the memory.
	// The memory should contain 3 bytes.
	m.StoreByte(0xF1, 0)
	m.StoreByte(0xF2, 10)
	m.StoreByte(0xF3, 34)

	// Load the first byte from memory.
	byte1 := m.LoadByte(0)
	expectedByte1 := byte(0xF1)
	if byte1 != expectedByte1 {
		t.Errorf("StoreByte() at offset 0 returned byte %v, wanted %v", byte1, expectedByte1)
	}

	// Load the second byte from memory.
	byte2 := m.LoadByte(1)
	expectedByte2 := byte(0x00)
	if byte2 != expectedByte2 {
		t.Errorf("LoadByte() at offset 1 returned byte %v, wanted %v", byte2, expectedByte2)
	}

	// Load the 10th byte from memory.
	byte10 := m.LoadByte(10)
	expectedByte10 := byte(0xF2)
	if byte10 != expectedByte10 {
		t.Errorf("StoreByte() at offset 10 returned byte %v, wanted %v", byte10, expectedByte10)
	}

	// Load the 34th byte from memory.
	byte34 := m.LoadByte(34)
	expectedByte34 := byte(0xF3)
	if byte34 != expectedByte34 {
		t.Errorf("StoreByte() at offset 34 returned byte %v, wanted %v", byte34, expectedByte34)
	}
}

func TestStoreWord(t *testing.T) {
	// Create an empty memory.
	m := NewMemory()

	// Store elements to the memory.
	// The memory should contain 3 words and an empty word of 32 bytes.
	word1 := uint256.NewInt(333).Bytes32()
	word2 := uint256.NewInt(222).Bytes32()
	word3 := uint256.NewInt(111).Bytes32()
	m.StoreWord(word1, 0)
	m.StoreWord(word2, 32*2)
	m.StoreWord(word3, 32*3)

	// Load the first word from memory.
	expectedWord1 := m.LoadWord(0)
	if !bytes.Equal(word1[:], expectedWord1[:]) {
		t.Errorf("StoreWord() at offset 32 returned word %v, wanted %v", word1, expectedWord1)
	}

	// Load empty word from memory.
	expectedEmptyWord := m.LoadWord(32)
	var emptyWord [32]byte
	if !bytes.Equal(emptyWord[:], expectedEmptyWord[:]) {
		t.Errorf("StoreWord() at offset 32 returned word %v, wanted empty word %v", emptyWord, expectedEmptyWord)
	}

	// Load the second word from memory.
	expectedWord2 := m.LoadWord(32 * 2)
	if !bytes.Equal(word2[:], expectedWord2[:]) {
		t.Errorf("StoreWord() at offset 32*2 returned word %v, wanted %v", word2, expectedWord2)
	}

	// Load the third word from memory.
	expectedWord3 := m.LoadWord(32 * 3)
	if !bytes.Equal(word3[:], expectedWord3[:]) {
		t.Errorf("StoreWord() at offset 32*3 returned word %v, wanted %v", word3, expectedWord3)
	}
}
