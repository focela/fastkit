// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package debug provides utilities for enabling and managing debug mode in the application.
package debug

import (
	"regexp"
	"runtime"
	"strconv"
)

// gridRegex extracts the goroutine ID from stack trace information.
var gridRegex = regexp.MustCompile(`^\w+\s+(\d+)\s+`)

// GoroutineId retrieves the current goroutine ID from stack information.
//
// Warning: This function uses runtime.Stack, which is not efficient. Avoid using
// it frequently in performance-critical code. It is mainly intended for debugging purposes.
//
// Returns:
// - int: The ID of the current goroutine.
func GoroutineId() int {
	// Allocate a small buffer for the stack trace.
	buf := make([]byte, 64) // Increased buffer size for safety.

	// Capture the stack trace.
	runtime.Stack(buf, false)

	// Extract the goroutine ID using regex.
	match := gridRegex.FindSubmatch(buf)
	if len(match) < 2 {
		// Return -1 if the ID cannot be extracted.
		return -1
	}

	// Convert the extracted ID to an integer.
	id, _ := strconv.Atoi(string(match[1]))
	return id
}
