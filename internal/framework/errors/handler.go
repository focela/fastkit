// Copyright (c) 2024 Focela Technologies
// Internal Use Only - Unauthorized use prohibited
// Contact: legal@focela.com

package errors

import (
	"github.com/focela/altura/internal/framework/errors/code"
)

// Equaler defines functionality for comparing equality between errors.
type Equaler interface {
	// Error returns the error message.
	Error() string

	// Equal compares the current error with a target error.
	Equal(target error) bool
}

// Coder defines functionality for retrieving error codes.
type Coder interface {
	// Error returns the error message.
	Error() string

	// Code retrieves the error code associated with the error.
	Code() code.Code
}

// Stacker defines functionality for retrieving error stack traces.
type Stacker interface {
	// Error returns the error message.
	Error() string

	// Stack retrieves the stack trace of the error.
	Stack() string
}

// Causer defines functionality for retrieving the root cause of an error.
type Causer interface {
	// Error returns the error message.
	Error() string

	// Cause retrieves the root cause of the error.
	Cause() error
}

// Currenter defines functionality for retrieving the current error.
type Currenter interface {
	// Error returns the error message.
	Error() string

	// Current retrieves the current error instance.
	Current() error
}

// Unwrapper defines functionality for unwrapping nested errors.
type Unwrapper interface {
	// Error returns the error message.
	Error() string

	// Unwrap retrieves the inner error wrapped within the current error.
	Unwrap() error
}

// CommaSeparator defines a comma separator with a space.
const CommaSeparator = ", "
