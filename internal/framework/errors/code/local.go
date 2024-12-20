// Copyright (c) 2024 Focela Technologies
// Internal Use Only - Unauthorized use prohibited
// Contact: legal@focela.com

package code

import (
	"fmt"
)

// localCode is an internal implementation of the Code interface for managing error codes.
type localCode struct {
	code    int         // Numeric error code, typically unique.
	message string      // Short description for the error code.
	detail  interface{} // Extended information or metadata about the error code.
}

// Code retrieves the numeric value of the error code.
func (c localCode) Code() int {
	return c.code
}

// Message retrieves the short description of the error code.
func (c localCode) Message() string {
	return c.message
}

// Detail retrieves additional information or metadata associated with the error code.
func (c localCode) Detail() interface{} {
	return c.detail
}

// String formats the error code, message, and detail (if available) into a readable string.
func (c localCode) String() string {
	switch {
	case c.detail != nil:
		return fmt.Sprintf("%d: %s %v", c.code, c.message, c.detail)
	case c.message != "":
		return fmt.Sprintf("%d: %s", c.code, c.message)
	default:
		return fmt.Sprintf("%d", c.code)
	}
}
