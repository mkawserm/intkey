package store

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type MemStore struct {
	mu    sync.Mutex
	store map[string]uint64
}

func (m *MemStore) Get(_ context.Context, key string) uint64 {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.store[key]
}

func (m *MemStore) Insert(_ context.Context, key string, value uint64) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, found := m.store[key]; !found {
		m.store[key] = value
		return true, nil
	}

	return false, errors.New("key already exists")
}

func (m *MemStore) Delete(_ context.Context, key string) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, found := m.store[key]; !found {
		return false, errors.New("key does not exists")
	}

	delete(m.store, key)

	return true, nil
}

func (m *MemStore) Increment(_ context.Context, key string, value uint64) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if old, found := m.store[key]; found {
		m.store[key] = old + value
		return true, nil
	}

	return false, errors.New("key does not exists")
}

func (m *MemStore) SafeIncrement(ctx context.Context, key string, value uint64) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
		fmt.Println("inside default")
		time.Sleep(time.Second * 40)
		fmt.Println("After 40 seconds")
		return true, nil
	}
	//ctx.Done()
	//m.mu.Lock()
	//defer m.mu.Unlock()
	//
	//if old, found := m.store[key]; found {
	//	m.store[key] = old + value
	//} else {
	//	m.store[key] = value
	//}
	//return true, nil
}

func (m *MemStore) Decrement(_ context.Context, key string, value uint64) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if old, found := m.store[key]; found {
		m.store[key] = old - value
		return true, nil
	}

	return false, errors.New("key does not exists")
}

var instantiated *MemStore
var once sync.Once

func MemStoreIns() *MemStore {
	once.Do(func() {
		instantiated = &MemStore{store: make(map[string]uint64)}
	})
	return instantiated
}
