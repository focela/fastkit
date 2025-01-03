// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package debug provides utilities for enabling and managing debug mode in the application.
package debug

import (
	"bytes"
	"fmt"
	"runtime"
)

// PrintStack prints the current goroutine stack trace to standard error.
// Optional `skip` parameter allows skipping specific stack frames.
func PrintStack(skip ...int) {
	fmt.Print(Stack(skip...))
}

// Stack returns a formatted stack trace of the current goroutine.
// Optional `skip` parameter allows skipping specific stack frames.
func Stack(skip ...int) string {
	return StackWithFilter(nil, skip...)
}

// StackWithFilter returns a filtered stack trace of the current goroutine.
// The `filters` parameter allows filtering stack frames based on path substrings.
// Optional `skip` parameter allows skipping specific stack frames.
func StackWithFilter(filters []string, skip ...int) string {
	return StackWithFilters(filters, skip...)
}

// StackWithFilters returns a filtered stack trace of the current goroutine.
// Filters are applied to remove unwanted stack frames.
// Optional `skip` parameter allows skipping specific stack frames.
func StackWithFilters(filters []string, skip ...int) string {
	skipFrames := 0
	if len(skip) > 0 {
		skipFrames = skip[0]
	}

	var (
		buffer = bytes.NewBuffer(nil)
		index  = 1
		space  = "  "
		ok     = true
		pc     uintptr
		file   string
		line   int
	)

	// Bắt đầu từ caller index sau khi áp dụng bộ lọc
	_, _, _, start := callerFromIndex(filters)

	for i := start + skipFrames; i < maxCallerDepth; i++ {
		pc, file, line, ok = runtime.Caller(i)
		if !ok {
			break
		}

		if filterFileByFilters(file, filters) {
			continue
		}

		funcName := "unknown"
		if fn := runtime.FuncForPC(pc); fn != nil {
			funcName = fn.Name()
		}

		if index > 9 {
			space = " "
		}

		buffer.WriteString(fmt.Sprintf("%d.%s%s\n    %s:%d\n", index, space, funcName, file, line))
		index++
	}

	return buffer.String()
}
