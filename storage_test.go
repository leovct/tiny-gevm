package main

import (
	"bytes"
	"testing"
)

func TestNewStorage(t *testing.T) {
	// Create an empty storage.
	s := NewStorage()
	if s == nil {
		t.Error("NewStorage() returned nil")
	}
}

func TestStoreAndLoad(t *testing.T) {
	// Create an empty storage.
	s := NewStorage()

	// Store a few values in the storage.
	value1 := [32]byte{0x01, 0x02, 0x03}
	s.Store(0, value1)

	value2 := [32]byte{0x04, 0x05, 0x06}
	s.Store(10, value2)

	// Load the values stored in storage.
	loaded1 := s.Load(0)
	if !bytes.Equal(loaded1[:], value1[:]) {
		t.Errorf("Expected %v, got %v", value1, loaded1)
	}

	loaded2 := s.Load(10)
	if !bytes.Equal(loaded2[:], value2[:]) {
		t.Errorf("Expected %v, got %v", value2, loaded2)
	}

	// Load an empty value from the storage.
	loaded3 := s.Load(20)
	emptyValue := [32]byte{}
	if !bytes.Equal(loaded3[:], emptyValue[:]) {
		t.Errorf("Expected %v, got %v", emptyValue, loaded3)
	}
}
