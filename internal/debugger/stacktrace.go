// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package debugger provides utilities for debugging, including logging and tracking application state.
package debugger

import (
	"bytes"
	"fmt"
	"runtime"
)

// Stack returns a formatted stack trace of the goroutine that calls it.
//
// Parameters:
// - skip: Number of stack frames to skip.
//
// This function internally calls StackWithFilter with a nil filter.
func Stack(skip ...int) string {
	return StackWithFilter(nil, skip...)
}

// StackWithFilter returns a formatted stack trace of the goroutine that calls it.
//
// Parameters:
// - filters: A slice of strings to filter specific paths in the stack trace.
// - skip: Number of stack frames to skip.
//
// This function internally calls StackWithFilters.
func StackWithFilter(filters []string, skip ...int) string {
	return StackWithFilters(filters, skip...)
}

// StackWithFilters returns a formatted stack trace of the goroutine that calls it.
//
// Parameters:
// - filters: A slice of strings used to filter the caller paths.
// - skip: Number of stack frames to skip.
//
// This function iterates through the runtime stack and applies the provided filters to format the stack trace.
//
// Note:
// - Performance improvements can be made by switching to debug.Stack or caching results.
func StackWithFilters(filters []string, skip ...int) string {
	number := 0
	if len(skip) > 0 {
		number = skip[0]
	}

	var (
		name                  string
		space                 = "  "
		index                 = 1
		buffer                = bytes.NewBuffer(nil)
		ok                    = true
		pc, file, line, start = callerFromIndex(filters)
	)

	for i := start + number; i < maxCallerDepth; i++ {
		if i != start {
			pc, file, line, ok = runtime.Caller(i)
		}
		if ok {
			if filterFileByFilters(file, filters) {
				continue
			}
			if fn := runtime.FuncForPC(pc); fn == nil {
				name = "unknown"
			} else {
				name = fn.Name()
			}
			if index > 9 {
				space = " "
			}
			buffer.WriteString(fmt.Sprintf("%d.%s%s\n    %s:%d\n", index, space, name, file, line))
			index++
		} else {
			break
		}
	}
	return buffer.String()
}

// PrintStack prints the stack trace of the current goroutine to the standard output.
//
// Parameters:
// - skip: Number of stack frames to skip.
//
// This function calls Stack to retrieve the stack trace.
func PrintStack(skip ...int) {
	fmt.Print(Stack(skip...))
}
