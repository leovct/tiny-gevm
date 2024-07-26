package evm

// IStorage defines the methods that a storage implementation should have.
type IStorage interface {
	// Store writes a 32-byte word to storage at the specified key.
	// If the key already exists, its value will be overwritten.
	Store(key int, value [32]byte)

	// Load retrieves a 32-byte word from storage using the specified key.
	// If the key does not exist in the storage, it returns an empty 32-byte word.
	Load(key int) [32]byte
}

// Storage represents a word-addressable storage structure.
type Storage struct {
	data map[int][32]byte
}

// NewStorage creates and returns a new, empty Storage instance.
func NewStorage() IStorage {
	return &Storage{data: make(map[int][32]byte)}
}

func (s *Storage) Store(key int, value [32]byte) {
	s.data[key] = value
}

func (s *Storage) Load(key int) [32]byte {
	return s.data[key]
}
