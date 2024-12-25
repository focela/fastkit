// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package code provides structures and utilities for managing error codes.
package code

import (
	"fmt"
)

// localCode represents an error code with associated message and detail.
// It is designed for internal usage only.
type localCode struct {
	code    int         // Error code, usually represented as an integer.
	message string      // Brief message describing the error code.
	detail  interface{} // Additional details or context for the error code.
}

// Code returns the integer representation of the error code.
func (c localCode) Code() int {
	return c.code
}

// Message returns the brief message associated with the error code.
func (c localCode) Message() string {
	return c.message
}

// Detail returns additional details associated with the error code.
// This field is mainly used for extending error context.
func (c localCode) Detail() interface{} {
	return c.detail
}

// String formats and returns the error code, message, and detail as a string.
func (c localCode) String() string {
	if c.detail != nil {
		return fmt.Sprintf(`%d: %s %v`, c.code, c.message, c.detail)
	}
	if c.message != "" {
		return fmt.Sprintf(`%d: %s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}
