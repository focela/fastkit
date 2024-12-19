// Copyright (c) 2024 Focela Technologies
// Internal Use Only - Unauthorized use prohibited
// Contact: legal@focela.com

package code

import (
	"fmt"
)

// localCode represents an internal implementation of the error code interface.
type localCode struct {
	code    int         // Numeric error code.
	message string      // Short description of the error.
	detail  interface{} // Additional details or context about the error.
}

// Code returns the integer value of the error code.
func (c localCode) Code() int {
	return c.code
}

// Message returns the short description of the error code.
func (c localCode) Message() string {
	return c.message
}

// Detail returns the additional details or context of the error code.
func (c localCode) Detail() interface{} {
	return c.detail
}

// String converts the error code to a formatted string representation.
// Includes the code, message, and detail (if present).
func (c localCode) String() string {
	if c.detail != nil {
		return fmt.Sprintf(`%d:%s %v`, c.code, c.message, c.detail)
	}
	if c.message != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}
