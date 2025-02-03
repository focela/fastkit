// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package mutex provides a wrapper for sync.Mutex with an optional concurrency-safe feature.
package mutex

import (
	"sync"
)

// Mutex is a wrapper around sync.Mutex that allows optional concurrency safety.
type Mutex struct {
	mutex *sync.Mutex // Underlying mutex for concurrency control
}

// New creates and returns a pointer to a new Mutex instance.
// The `safe` parameter determines whether the mutex should be concurrency-safe.
// If `safe` is true, the mutex is initialized with a sync.Mutex.
func New(safe ...bool) *Mutex {
	mu := Create(safe...)
	return &mu
}

// Create initializes and returns a new Mutex instance.
// The `safe` parameter determines whether the mutex should be concurrency-safe.
func Create(safe ...bool) Mutex {
	if len(safe) > 0 && safe[0] {
		return Mutex{mutex: new(sync.Mutex)}
	}
	return Mutex{}
}

// IsSafe returns true if the mutex is in concurrency-safe mode.
func (mu *Mutex) IsSafe() bool {
	return mu.mutex != nil
}

// Lock acquires the mutex lock if concurrency-safe mode is enabled.
func (mu *Mutex) Lock() {
	if mu.IsSafe() {
		mu.mutex.Lock()
	}
}

// Unlock releases the mutex lock if concurrency-safe mode is enabled.
func (mu *Mutex) Unlock() {
	if mu.IsSafe() {
		mu.mutex.Unlock()
	}
}
