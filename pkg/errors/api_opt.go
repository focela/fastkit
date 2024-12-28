// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package errors provides rich functionalities to manipulate errors.
package errors

import (
	"github.com/focela/loom/pkg/errors/code"
)

// Option defines the configuration for creating custom errors.
type Option struct {
	Error error     // Wrapped error if any.
	Stack bool      // Whether to record stack trace information.
	Text  string    // Custom error text.
	Code  code.Code // Associated error code.
}

// NewWithOption creates and returns a custom error using the provided Option.
// It is primarily used internally within the framework.
func NewWithOption(option Option) error {
	err := &Error{
		error: option.Error,
		text:  option.Text,
		code:  option.Code,
	}
	if option.Stack {
		err.stack = callers()
	}
	return err
}

// NewOption is a deprecated alias for NewWithOption.
// Deprecated: Use NewWithOption instead.
func NewOption(option Option) error {
	return NewWithOption(option)
}
