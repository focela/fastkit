// Copyright (c) 2025 Focela Technologies.
// This software is provided "as is", without any warranty.
// Licensed under the MIT License – see LICENSE file for details.

// Package mutex provides a mutex implementation with optional concurrent safety.
// It wraps sync.Mutex with the ability to disable locking for better performance when safety isn't needed.
package mutex

import (
	"sync"
)

// Mutex is a sync.Mutex with a switch for concurrent safe feature.
// If its underlying mutex is not nil, concurrent safety is enabled.
// By default, the underlying mutex is nil, making this struct lightweight when safety isn't required.
type Mutex struct {
	// mutex is the underlying sync.Mutex for thread-safety.
	// When nil, locking operations become no-ops for better performance.
	mutex *sync.Mutex
}

// New creates and returns a new *Mutex.
// The optional parameter `safe` specifies whether to enable concurrent safety.
// By default (without parameters), safety is disabled for better performance.
func New(safe ...bool) *Mutex {
	mu := Create(safe...)
	return &mu
}

// Create returns a new Mutex value (not a pointer).
// The optional parameter `safe` specifies whether to enable concurrent safety.
// By default (without parameters), safety is disabled for better performance.
func Create(safe ...bool) Mutex {
	if len(safe) > 0 && safe[0] {
		return Mutex{
			mutex: new(sync.Mutex),
		}
	}
	return Mutex{}
}

// IsSafe returns whether concurrent safety is enabled for this mutex.
// Returns true if the underlying mutex is initialized, false otherwise.
func (mu *Mutex) IsSafe() bool {
	return mu.mutex != nil
}

// Lock acquires an exclusive lock.
// If safety is disabled, this operation does nothing.
func (mu *Mutex) Lock() {
	if mu.mutex != nil {
		mu.mutex.Lock()
	}
}

// Unlock releases an exclusive lock.
// If safety is disabled, this operation does nothing.
func (mu *Mutex) Unlock() {
	if mu.mutex != nil {
		mu.mutex.Unlock()
	}
}
