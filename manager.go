package redfox

import (
	"errors"
	"sync"

	"github.com/evilmagics/go-redfox/lib"
)

// Manager is a generic thread-safe cache for storing and retrieving exceptions
// using a comparable key type. It provides synchronized access to a map of exceptions.
type Manager[T comparable] struct {
	mu    sync.RWMutex
	cache map[T]Exception[T]
}

// Set replaces the entire cache with the provided map of exceptions.
// It acquires a write lock to ensure thread-safe modification of the internal cache.
// Use Set() carefully, as it will overwrite the entire cache.
func (m *Manager[T]) Set(exceptions map[T]Exception[T]) *Manager[T] {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cache = exceptions
	return m
}

// Add adds an exception to the manager's cache, using the exception's error code as the key.
// It acquires a write lock to ensure thread-safe modification of the internal cache.
func (m *Manager[T]) Add(err Exception[T]) *Manager[T] {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cache[err.ErrCode()] = err
	return m
}

// SafeAdd insert an exception to the manager's cache, using the exception's error code as the key.
// It acquires a write lock to ensure thread-safe modification of the internal cache.
// Add to registered key will return an error if the key already exists
func (m *Manager[T]) SafeAdd(err Exception[T]) error {
	if _, ok := m.cache[err.ErrCode()]; ok {
		return errors.New("exception (" + lib.Stringify(err.ErrCode()) + " ) already exists")
	}

	m.Add(err)
	return nil
}

// AddAll adds multiple exceptions to the manager's cache, using the exception's error code as the key.
// It acquires a write lock to ensure thread-safe modification of the internal cache.
func (m *Manager[T]) AddAll(errs ...Exception[T]) *Manager[T] {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, err := range errs {
		m.cache[err.ErrCode()] = err
	}
	return m
}

// SafeAddAll adds multiple exceptions to the manager's cache, using the exception's error code as the key.
// It acquires a write lock to ensure thread-safe modification of the internal cache.
// Add to registered key will return an error if the key already exists
func (m *Manager[T]) SafeAddAll(errs ...Exception[T]) error {
	for _, err := range errs {
		if _, ok := m.cache[err.ErrCode()]; ok {
			return errors.New("exception (" + lib.Stringify(err.ErrCode()) + " ) already exists")
		}
	}

	m.AddAll(errs...)
	return nil
}

// Get retrieves an exception from the manager's cache by its error code.
// It acquires a read lock to ensure thread-safe access to the internal cache
// and returns a clone of the exception to prevent direct modification.
func (m *Manager[T]) Get(code T) Exception[T] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.cache[code].Clone()
}

// GetAll retrieves all exceptions from the manager's cache.
// It acquires a read lock to ensure thread-safe access and returns a slice of cloned exceptions
// to prevent direct modification of the original cache entries.
// The returned slice contains copies of all exceptions currently stored in the cache.
func (m *Manager[T]) GetAll() []Exception[T] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	exceptions := make([]Exception[T], 0, len(m.cache))
	for _, exception := range m.cache {
		exceptions = append(exceptions, exception.Clone())
	}
	return exceptions
}

// Size returns the number of exceptions currently stored in the manager's cache.
// It acquires a read lock to ensure thread-safe access and returns the current size of the cache.
func (m *Manager[T]) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.cache)
}

// IsEmpty checks if the manager's cache is empty.
// It acquires a read lock to ensure thread-safe access and returns true if
func (m *Manager[T]) IsEmpty() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.cache) == 0
}

// IsNotEmpty checks if the manager's cache contains any entries.
// It acquires a read lock to ensure thread-safe access and returns true if
// the cache has at least one entry, otherwise returns false.
func (m *Manager[T]) IsNotEmpty() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.cache) > 0
}

// Contains checks if the specified code exists in the manager's cache.
// It acquires a read lock to ensure thread-safe access and returns true if
// the given code is found in the cache, otherwise returns false.
func (m *Manager[T]) Contains(code T) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, exists := m.cache[code]
	return exists
}

// ContainsAll checks if all of the provided codes exist in the manager's cache.
// It acquires a read lock to ensure thread-safe access and returns true if all
// of the given codes are found in the cache, otherwise returns false.
func (m *Manager[T]) ContainsAll(codes ...T) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, code := range codes {
		if _, exists := m.cache[code]; !exists {
			return false
		}
	}
	return true
}

// ContainsAny checks if any of the provided codes exist in the manager's cache.
// It acquires a read lock to ensure thread-safe access and returns true if at least
// one of the given codes is found in the cache, otherwise returns false.
func (m *Manager[T]) ContainsAny(codes ...T) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, code := range codes {
		if _, exists := m.cache[code]; exists {
			return true
		}
	}
	return false
}

// Remove deletes the specified code from the cache, ensuring thread-safe access
// by acquiring a write lock before removing the entry from the underlying map.
func (m *Manager[T]) Remove(code T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.cache, code)
}

// Clear removes all entries from the cache, ensuring thread-safe access
// by acquiring a write lock before reinitializing the underlying map.
func (m *Manager[T]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cache = make(map[T]Exception[T])

}
