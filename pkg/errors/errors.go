// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package errors provides rich functionalities to manipulate errors.
//
// For maintainers, please note that this package is a foundational package.
// It SHOULD NOT import extra packages except standard packages and internal packages
// to avoid cyclic imports.
package errors

import (
	"github.com/focela/loom/pkg/errors/code"
)

// Interface definitions for error handling features.

// Equaler defines an interface for comparing errors.
type Equaler interface {
	Error() string
	Equal(target error) bool
}

// Coder defines an interface for retrieving error codes.
type Coder interface {
	Error() string
	Code() code.Code
}

// Stacker defines an interface for retrieving stack trace.
type Stacker interface {
	Error() string
	Stack() string
}

// Causer defines an interface for retrieving root cause of an error.
type Causer interface {
	Error() string
	Cause() error
}

// Currenter defines an interface for retrieving the current error.
type Currenter interface {
	Error() string
	Current() error
}

// Unwrapper defines an interface for unwrapping errors.
type Unwrapper interface {
	Error() string
	Unwrap() error
}

// Common constants used in error formatting.
const (
	// commaSeparatorSpace is the comma separator with space.
	commaSeparatorSpace = ", "
)
