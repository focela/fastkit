// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package errors provides rich functionalities to manipulate errors.
package errors

import (
	"fmt"
	"strings"

	"github.com/focela/loom/pkg/errors/code"
)

// NewCode creates and returns an error with an error code and optional text.
func NewCode(code code.Code, text ...string) error {
	return &Error{
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// NewCodef creates and returns an error with an error code and formatted text.
func NewCodef(code code.Code, format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// NewCodeSkip creates and returns an error with an error code and optional text.
// The `skip` parameter specifies the number of stack frames to skip.
func NewCodeSkip(code code.Code, skip int, text ...string) error {
	return &Error{
		stack: callers(skip),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// NewCodeSkipf creates and returns an error with an error code and formatted text.
// The `skip` parameter specifies the number of stack frames to skip.
func NewCodeSkipf(code code.Code, skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCode wraps an existing error with a code and optional text.
// Returns nil if the provided error is nil.
func WrapCode(code code.Code, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// WrapCodef wraps an existing error with a code and formatted text.
// Returns nil if the provided error is nil.
func WrapCodef(code code.Code, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCodeSkip wraps an existing error with a code and optional text.
// The `skip` parameter specifies the number of stack frames to skip.
func WrapCodeSkip(code code.Code, skip int, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// WrapCodeSkipf wraps an existing error with a code and formatted text.
// The `skip` parameter specifies the number of stack frames to skip.
func WrapCodeSkipf(code code.Code, skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// Code retrieves the error code from an error.
// Returns `CodeNil` if the error does not have an associated code.
func Code(err error) code.Code {
	if err == nil {
		return code.CodeNil
	}
	if e, ok := err.(Coder); ok {
		return e.Code()
	}
	if e, ok := err.(Unwrapper); ok {
		return Code(e.Unwrap())
	}
	return code.CodeNil
}

// HasCode checks if the error or any error in its chain has the specified error code.
func HasCode(err error, code code.Code) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(Coder); ok && code == e.Code() {
		return true
	}
	if e, ok := err.(Unwrapper); ok {
		return HasCode(e.Unwrap(), code)
	}
	return false
}
