// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package errors provides rich functionalities to manipulate errors.
package errors

import (
	"github.com/focela/loom/pkg/errors/code"
)

// Code retrieves the error code associated with the error instance.
// It returns `CodeNil` if no code is set.
func (err *Error) Code() code.Code {
	if err == nil {
		return code.CodeNil
	}
	if err.code == code.CodeNil {
		return Code(err.Unwrap())
	}
	return err.code
}

// SetCode sets or updates the error code for the current error instance.
func (err *Error) SetCode(c code.Code) {
	if err == nil {
		return
	}
	err.code = c
}
