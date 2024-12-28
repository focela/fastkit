// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package errors provides rich functionalities to manipulate errors.
package errors

import (
	"fmt"

	"github.com/focela/loom/pkg/errors/code"
)

// New creates and returns an error with the given text.
func New(text string) error {
	return &Error{
		stack: callers(),
		text:  text,
		code:  code.CodeNil,
	}
}

// Newf creates and returns an error with formatted text.
func Newf(format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code.CodeNil,
	}
}

// NewSkip creates and returns an error with the given text.
// The `skip` parameter specifies the number of stack frames to skip.
func NewSkip(skip int, text string) error {
	return &Error{
		stack: callers(skip),
		text:  text,
		code:  code.CodeNil,
	}
}

// NewSkipf creates and returns an error with formatted text.
// The `skip` parameter specifies the number of stack frames to skip.
func NewSkipf(skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code.CodeNil,
	}
}

// Wrap wraps an existing error with additional text.
// Returns nil if the given error is nil.
func Wrap(err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  text,
		code:  Code(err),
	}
}

// Wrapf wraps an existing error with formatted text.
// Returns nil if the given error is nil.
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  Code(err),
	}
}

// WrapSkip wraps an existing error with additional text.
// The `skip` parameter specifies the number of stack frames to skip.
// Returns nil if the given error is nil.
func WrapSkip(skip int, err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  text,
		code:  Code(err),
	}
}

// WrapSkipf wraps an existing error with formatted text.
// The `skip` parameter specifies the number of stack frames to skip.
// Returns nil if the given error is nil.
func WrapSkipf(skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  Code(err),
	}
}
