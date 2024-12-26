// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package code provides structures and utilities for managing error codes.
package code

// Code is a universal error code interface definition.
type Code interface {
	// Code returns the integer number of the current error code.
	Code() int
	// Message returns a brief message for the current error code.
	Message() string
	// Detail returns detailed information for the current error code.
	Detail() interface{}
}

// Common error code definitions.
// Reserved internal error codes by the framework: code < 1000.
var (
	// General Codes
	CodeNil            = localCode{-1, "", nil}                // No error code specified.
	CodeOK             = localCode{0, "OK", nil}               // Everything is fine.
	CodeUnknown        = localCode{64, "Unknown Error", nil}   // Unknown error.
	CodeNotFound       = localCode{65, "Not Found", nil}       // Resource does not exist.
	CodeInvalidRequest = localCode{66, "Invalid Request", nil} // Invalid request.

	// Internal Codes
	CodeInternalError = localCode{50, "Internal Error", nil} // An error occurred internally.
	CodeInternalPanic = localCode{68, "Internal Panic", nil} // A panic occurred internally.
	CodeServerBusy    = localCode{63, "Server Is Busy", nil} // Server is busy, please try again later.

	// Validation & Configuration Codes
	CodeValidationFailed     = localCode{51, "Validation Failed", nil}     // Data validation failed.
	CodeInvalidParameter     = localCode{53, "Invalid Parameter", nil}     // Invalid parameter.
	CodeMissingParameter     = localCode{54, "Missing Parameter", nil}     // Missing parameter.
	CodeInvalidConfiguration = localCode{56, "Invalid Configuration", nil} // Invalid configuration.
	CodeMissingConfiguration = localCode{57, "Missing Configuration", nil} // Missing configuration.

	// Authorization & Security Codes
	CodeNotAuthorized  = localCode{61, "Not Authorized", nil}  // Not authorized.
	CodeSecurityReason = localCode{62, "Security Reason", nil} // Security-related issue.

	// Operation & Support Codes
	CodeInvalidOperation = localCode{55, "Invalid Operation", nil} // Invalid operation.
	CodeOperationFailed  = localCode{60, "Operation Failed", nil}  // Operation failed.
	CodeNotImplemented   = localCode{58, "Not Implemented", nil}   // Not implemented yet.
	CodeNotSupported     = localCode{59, "Not Supported", nil}     // Operation not supported.

	// Business Logic Codes
	CodeBusinessValidationFailed = localCode{300, "Business Validation Failed", nil} // Business validation failed.
)

// New creates and returns a new error code.
// It generates a localCode instance with the specified code, message, and detail.
func New(code int, message string, detail interface{}) Code {
	return localCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode creates and returns a new error code based on an existing Code instance.
// The code and message are copied from the given `code`, while the detail comes from the provided `detail`.
func WithCode(code Code, detail interface{}) Code {
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
