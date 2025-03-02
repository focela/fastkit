// Copyright (c) 2025 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package lock provides a mutex implementation with the ability to enable or disable
// concurrent safety as needed. This allows for creating mutex objects that can be
// selectively used for synchronization based on application requirements.
package lock

import (
	"sync"
)

// Mutex is a wrapper around sync.Mutex with a switch for concurrent safe feature.
// When the safe mode is disabled, all Lock/Unlock operations become no-ops,
// allowing the same code to run efficiently in both single-threaded and multi-threaded
// contexts without changes.
type Mutex struct {
	// mutex is the underlying sync.Mutex used when safe mode is enabled.
	// When nil, the mutex is in unsafe mode (no locking occurs).
	mutex *sync.Mutex
}

// New creates and returns a new *Mutex.
//
// Parameters:
//   - safe: Optional boolean indicating whether the mutex should operate in thread-safe mode.
//     If not provided or false, the mutex will not perform actual locking (unsafe mode).
//     If true, the mutex will use an underlying sync.Mutex for thread safety.
//
// Returns:
//   - A pointer to a newly created Mutex object
func New(safe ...bool) *Mutex {
	mu := Create(safe...)
	return &mu
}

// Create creates and returns a new Mutex object (not a pointer).
//
// Parameters:
//   - safe: Optional boolean indicating whether the mutex should operate in thread-safe mode.
//     If not provided or false, the mutex will not perform actual locking (unsafe mode).
//     If true, the mutex will use an underlying sync.Mutex for thread safety.
//
// Returns:
//   - A newly created Mutex object
func Create(safe ...bool) Mutex {
	if len(safe) > 0 && safe[0] {
		return Mutex{
			mutex: new(sync.Mutex),
		}
	}
	return Mutex{}
}

// IsSafe checks and returns whether current mutex is in concurrent-safe usage.
//
// Returns:
//   - true if the mutex is operating in thread-safe mode
//   - false if the mutex is operating in unsafe mode (no locking)
func (mu *Mutex) IsSafe() bool {
	return mu.mutex != nil
}

// Lock acquires an exclusive lock on the mutex.
// If the mutex is in unsafe mode (not concurrent-safe), this operation does nothing.
func (mu *Mutex) Lock() {
	if mu.mutex != nil {
		mu.mutex.Lock()
	}
}

// Unlock releases an exclusive lock on the mutex.
// If the mutex is in unsafe mode (not concurrent-safe), this operation does nothing.
func (mu *Mutex) Unlock() {
	if mu.mutex != nil {
		mu.mutex.Unlock()
	}
}
