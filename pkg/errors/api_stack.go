// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package errors provides rich functionalities to manipulate errors.
package errors

import (
	"errors"
	"runtime"
)

// stack represents a stack of program counters.
type stack []uintptr

const (
	// maxStackDepth marks the maximum stack depth for error back traces.
	maxStackDepth = 64
)

// Cause retrieves the root cause of the given error.
// If the error implements the Causer interface, it will return the root cause.
func Cause(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(Causer); ok {
		return e.Cause()
	}
	if e, ok := err.(Unwrapper); ok {
		return Cause(e.Unwrap())
	}
	return err
}

// Stack retrieves the stack trace of the given error.
// If the error implements the Stacker interface, it will return the stack trace.
func Stack(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(Stacker); ok {
		return e.Stack()
	}
	return err.Error()
}

// Current retrieves the current level error.
// If the error implements the Currenter interface, it will return the current error instance.
func Current(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(Currenter); ok {
		return e.Current()
	}
	return err
}

// Unwrap retrieves the next level error.
// If the error implements the Unwrapper interface, it will return the next error in the chain.
func Unwrap(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(Unwrapper); ok {
		return e.Unwrap()
	}
	return nil
}

// HasStack checks whether the given error implements the Stacker interface.
func HasStack(err error) bool {
	_, ok := err.(Stacker)
	return ok
}

// Equal compares two errors for equality.
// If either error implements the Equaler interface, it will use the Equal method for comparison.
func Equal(err, target error) bool {
	if err == target {
		return true
	}
	if e, ok := err.(Equaler); ok {
		return e.Equal(target)
	}
	if e, ok := target.(Equaler); ok {
		return e.Equal(err)
	}
	return false
}

// Is checks whether the given error matches the target error.
// It uses the standard library's errors.Is for comparison.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As finds the first error in the error chain that matches the target.
// It uses the standard library's errors.As for matching.
func As(err error, target any) bool {
	return errors.As(err, target)
}

// HasError is an alias for Is.
// Deprecated: Use Is instead.
func HasError(err, target error) bool {
	return errors.Is(err, target)
}

// callers retrieves the stack callers.
// It collects the caller's memory addresses but does not include caller details.
func callers(skip ...int) stack {
	var (
		pcs [maxStackDepth]uintptr
		n   = 3
	)
	if len(skip) > 0 {
		n += skip[0]
	}
	return pcs[:runtime.Callers(n, pcs[:])]
}
