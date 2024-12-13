package core

import (
	"errors"
	"sync"
)

var ErrKeyNotFound = errors.New("key not found")

// InMemoryStore stores the key-value data in memory (RAM) using a simple hashmap.
// Note: this implementation is thread-safe using a ReadWrite lock.
type InMemoryStore struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		data: make(map[string][]byte),
	}
}

// Set sets a new value (or overrides an existing one) for the given key.
func (s *InMemoryStore) Set(key string, val []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = val
}

// Get retrieves the value associated with the given key or returns an error if it does not exist.
func (s *InMemoryStore) Get(key string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	if !ok {
		return nil, ErrKeyNotFound
	}

	return val, nil
}

// Delete deletes the entry associated with the given key.
func (s *InMemoryStore) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
}
