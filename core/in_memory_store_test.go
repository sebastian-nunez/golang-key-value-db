package core

import (
	"testing"
)

func TestInMemoryStore(t *testing.T) {
	t.Parallel()

	t.Run("retrieves the value of a key within the store", func(t *testing.T) {
		store := NewInMemoryStore()
		key := "some-key"
		val := "some-value"
		store.Set(key, []byte(val))

		value, err := store.Get(key)

		if err != nil {
			t.Errorf("Expected err to be nil, got %v", err)
		}
		if string(value) != val {
			t.Errorf("Expected value to be %s, got %v", val, value)
		}
	})

	t.Run("fails to retrieve a key that is not present in the store", func(t *testing.T) {
		store := NewInMemoryStore()
		key := "does not exist"

		_, err := store.Get(key)

		if err != ErrKeyNotFound {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("deletes a valid key from the store", func(t *testing.T) {
		store := NewInMemoryStore()
		key := "does not exist"
		val := "some value to be deleted"

		store.Set(key, []byte(val))
		store.Delete(key)
		_, err := store.Get(key)

		if err != ErrKeyNotFound {
			t.Errorf("expected error but got nil")
		}
	})
}
