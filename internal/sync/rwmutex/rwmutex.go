// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package rwmutex provides a wrapper around sync.RWMutex with an optional concurrency-safe feature.
package rwmutex

import (
	"sync"
)

// RWMutex is a wrapper around sync.RWMutex that allows optional concurrency safety.
type RWMutex struct {
	mutex *sync.RWMutex
}

// New creates and returns a pointer to a new RWMutex instance.
// The `safe` parameter determines whether the mutex should be concurrency-safe.
// If `safe` is true, the mutex is initialized with a sync.RWMutex.
func New(safe ...bool) *RWMutex {
	mu := Create(safe...)
	return &mu
}

// Create initializes and returns a new RWMutex instance.
// The `safe` parameter determines whether the mutex should be concurrency-safe.
func Create(safe ...bool) RWMutex {
	if len(safe) > 0 && safe[0] {
		return RWMutex{mutex: new(sync.RWMutex)}
	}
	return RWMutex{}
}

// IsSafe returns true if the mutex is in concurrency-safe mode.
func (mu *RWMutex) IsSafe() bool {
	return mu.mutex != nil
}

// Lock acquires the write lock if concurrency-safe mode is enabled.
func (mu *RWMutex) Lock() {
	if mu.IsSafe() {
		mu.mutex.Lock()
	}
}

// Unlock releases the write lock if concurrency-safe mode is enabled.
func (mu *RWMutex) Unlock() {
	if mu.IsSafe() {
		mu.mutex.Unlock()
	}
}

// RLock acquires the read lock if concurrency-safe mode is enabled.
func (mu *RWMutex) RLock() {
	if mu.IsSafe() {
		mu.mutex.RLock()
	}
}

// RUnlock releases the read lock if concurrency-safe mode is enabled.
func (mu *RWMutex) RUnlock() {
	if mu.IsSafe() {
		mu.mutex.RUnlock()
	}
}
