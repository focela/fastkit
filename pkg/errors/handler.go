// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package errors provides rich functionalities to manipulate errors.
package errors

import (
	"runtime"
	"strings"

	"github.com/focela/loom/pkg/errors/code"
)

// Error represents a custom error with additional features.
type Error struct {
	error error     // Wrapped error.
	stack stack     // Stack array, records stack trace information when error is created.
	text  string    // Custom error message, may be empty if a code is provided.
	code  code.Code // Associated error code.
}

const (
	// stackFilterKeyLocal filters paths for the current error module.
	stackFilterKeyLocal = "/pkg/errors/errors"
)

// goRootForFilter is used for stack filtering, primarily in development environments.
var goRootForFilter = runtime.GOROOT()

// init initializes the stack filtering path.
// It normalizes the GOROOT path for stack trace filtering.
func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
	}
}
