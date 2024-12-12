package core

import (
	"errors"
	"sync"
)

var ErrKeyNotFound = errors.New("key not found")

type InMemoryStore struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		data: make(map[string][]byte),
	}
}

func (s *InMemoryStore) Set(key string, val []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = val
}

func (s *InMemoryStore) Get(key string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	if !ok {
		return nil, ErrKeyNotFound
	}

	return val, nil
}

func (s *InMemoryStore) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
}
