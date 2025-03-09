// Copyright (c) 2025 Focela Technologies.
// This software is provided "as is", without any warranty.
// Licensed under the MIT License – see LICENSE file for details.

// Package rwmutex provides a read-write mutex implementation with optional concurrent safety.
// It wraps sync.RWMutex with the ability to disable locking for better performance when safety isn't needed.
package rwmutex

import (
	"sync"
)

// RWMutex is a sync.RWMutex with a switch for concurrent safe feature.
// If its underlying mutex is not nil, concurrent safety is enabled.
// By default, the underlying mutex is nil, making this struct much lightweight when safety isn't required.
type RWMutex struct {
	// mutex is the underlying sync.RWMutex for thread-safety.
	// When nil, locking operations become no-ops for better performance.
	mutex *sync.RWMutex
}

// New creates and returns a new *RWMutex.
// The optional parameter `safe` specifies whether to enable concurrent safety.
// By default (without parameters), safety is disabled for better performance.
func New(safe ...bool) *RWMutex {
	mu := Create(safe...)
	return &mu
}

// Create returns a new RWMutex value (not a pointer).
// The optional parameter `safe` specifies whether to enable concurrent safety.
// By default (without parameters), safety is disabled for better performance.
func Create(safe ...bool) RWMutex {
	if len(safe) > 0 && safe[0] {
		return RWMutex{
			mutex: new(sync.RWMutex),
		}
	}
	return RWMutex{}
}

// IsSafe returns whether concurrent safety is enabled for this mutex.
// Returns true if the underlying mutex is initialized, false otherwise.
func (mu *RWMutex) IsSafe() bool {
	return mu.mutex != nil
}

// Lock acquires an exclusive lock for writing.
// If safety is disabled, this operation does nothing.
func (mu *RWMutex) Lock() {
	if mu.mutex != nil {
		mu.mutex.Lock()
	}
}

// Unlock releases an exclusive lock.
// If safety is disabled, this operation does nothing.
func (mu *RWMutex) Unlock() {
	if mu.mutex != nil {
		mu.mutex.Unlock()
	}
}

// RLock acquires a shared lock for reading.
// Multiple goroutines can hold read locks simultaneously.
// If safety is disabled, this operation does nothing.
func (mu *RWMutex) RLock() {
	if mu.mutex != nil {
		mu.mutex.RLock()
	}
}

// RUnlock releases a shared lock.
// If safety is disabled, this operation does nothing.
func (mu *RWMutex) RUnlock() {
	if mu.mutex != nil {
		mu.mutex.RUnlock()
	}
}
