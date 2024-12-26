// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package errors provides rich functionalities to manipulate errors.
package errors

import (
	"errors"
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

// Error returns the error message as a string.
// If both `text` and `code` are empty, it returns the wrapped error's message.
func (err *Error) Error() string {
	if err == nil {
		return ""
	}
	errStr := err.text
	if errStr == "" && err.code != nil {
		errStr = err.code.Message()
	}
	if err.error != nil {
		if errStr != "" {
			errStr += ": "
		}
		errStr += err.error.Error()
	}
	return errStr
}

// Cause retrieves the root cause of the error.
// It unwraps the error recursively until the base cause is found.
func (err *Error) Cause() error {
	if err == nil {
		return nil
	}
	loop := err
	for loop != nil {
		if loop.error != nil {
			if e, ok := loop.error.(*Error); ok {
				// Internal Error struct.
				loop = e
			} else if e, ok := loop.error.(Causer); ok {
				// External error implementing the Causer interface.
				return e.Cause()
			} else {
				return loop.error
			}
		} else {
			return errors.New(loop.text)
		}
	}
	return nil
}

// Current returns a copy of the current error instance.
// If the error is nil, it returns nil.
func (err *Error) Current() error {
	if err == nil {
		return nil
	}
	return &Error{
		error: nil,
		stack: err.stack,
		text:  err.text,
		code:  err.code,
	}
}

// Unwrap returns the next error in the error chain.
// It is compatible with Go's standard library errors.Unwrap.
func (err *Error) Unwrap() error {
	if err == nil {
		return nil
	}
	return err.error
}
