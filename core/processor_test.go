package core

import (
	"context"
	"testing"
	"time"
)

func TestProcessor(t *testing.T) {
	ctx := context.Background()
	store := &MockDataStore{}
	cp := NewCommandProcessor(store)

	t.Run("should call GET in the datastore given a valid request", func(t *testing.T) {
		req := Request{Command: CmdGet, Params: []string{"key1"}}

		cp.Process(ctx, req)

		if !store.GetCalled {
			t.Errorf("Get not called")
		}
	})

	t.Run("should call SET in the datastore given a valid request", func(t *testing.T) {
		req := Request{Command: CmdSet, Params: []string{"key1", "value1"}}

		cp.Process(ctx, req)

		if !store.SetCalled {
			t.Errorf("Set not called")
		}
	})

	t.Run("should call DELETE in the datastore given a valid request", func(t *testing.T) {
		req := Request{Command: CmdDelete, Params: []string{"key1"}}

		cp.Process(ctx, req)

		if !store.DeleteCalled {
			t.Errorf("Delete not called")
		}
	})

	t.Run("should call DELETE in the datastore for a key after the TTL has expired", func(t *testing.T) {
		req := Request{Command: CmdSet, Params: []string{"key1", "value1", "2"}}

		cp.Process(ctx, req)

		if !store.SetCalled {
			t.Errorf("Set not called")
		}
		time.Sleep(2 * time.Second)
		if !store.DeleteCalled {
			t.Errorf("Delete not called for TTL")
		}
	})
}

type MockDataStore struct {
	GetCalled    bool
	SetCalled    bool
	DeleteCalled bool
}

func (mds *MockDataStore) Get(key string) ([]byte, error) {
	mds.GetCalled = true
	return []byte("value1"), nil
}
func (mds *MockDataStore) Set(key string, value []byte) {
	mds.SetCalled = true
}
func (mds *MockDataStore) Delete(key string) {
	mds.DeleteCalled = true
}
