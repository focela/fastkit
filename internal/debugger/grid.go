// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package debugger provides utilities for debugging, including logging and tracking application state.
package debugger

import (
	"regexp"
	"runtime"
	"strconv"
)

// gridRegex is the regular expression for extracting the goroutine ID
// from the runtime stack trace. It matches the format of goroutine headers.
var gridRegex = regexp.MustCompile(`^\w+\s+(\d+)\s+`)

// GoroutineId retrieves and returns the current goroutine ID from stack information.
//
// Note:
// - This function uses runtime.Stack, which is not performant and should be used for debugging purposes only.
// - If the goroutine ID cannot be retrieved, it returns -1.
func GoroutineId() int {
	// Buffer size of 26 is sufficient for parsing goroutine headers in the stack trace.
	buf := make([]byte, 26)
	n := runtime.Stack(buf, false)

	// Match the stack trace against the regular expression.
	match := gridRegex.FindSubmatch(buf[:n])
	if len(match) < 2 {
		// Return -1 if the ID could not be extracted.
		return -1
	}

	// Convert the matched ID from string to integer.
	id, err := strconv.Atoi(string(match[1]))
	if err != nil {
		// Return -1 if conversion fails.
		return -1
	}

	return id
}
