// Copyright (c) 2024 Focela Technologies
// Internal Use Only - Unauthorized use prohibited
// Contact: legal@focela.com

// Package errors provides functionalities to manipulate errors for internal usage purpose.
package errors

import (
	"github.com/focela/altura/internal/framework/errors/code"
)

// Equaler defines an interface for error equality comparison.
type Equaler interface {
	Error() string
	Equal(target error) bool
}

// Coder defines an interface for retrieving error codes.
type Coder interface {
	Error() string
	Code() code.Code
}

// StackTracer defines an interface for retrieving the stack trace of an error.
type StackTracer interface {
	Error() string
	Stack() string
}

// Causer defines an interface for accessing the root cause of an error.
type Causer interface {
	Error() string
	Cause() error
}

// CurrentError defines an interface for retrieving the current error.
type CurrentError interface {
	Error() string
	Current() error
}

// Unwrapper defines an interface for unwrapping nested errors.
type Unwrapper interface {
	Error() string
	Unwrap() error
}

// commaSeparatorSpace is the comma separator with a space, used in error formatting.
const (
	commaSeparatorSpace = ", "
)
