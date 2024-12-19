// Copyright (c) 2024 Focela Technologies
// Internal Use Only - Unauthorized use prohibited
// Contact: legal@focela.com

// Package code provides universal error code definition and common error code implementations.
package code

// Code defines a universal error code interface.
type Code interface {
	// Code returns the integer value of the error code.
	Code() int

	// Message returns the brief message for the error code.
	Message() string

	// Detail returns additional details or context for the error code.
	Detail() interface{}
}

// Common error code definitions.
var (
	CodeNil                       = localCode{-1, "", nil}                             // No error code specified.
	CodeOK                        = localCode{0, "OK", nil}                            // It is OK.
	CodeInternalError             = localCode{50, "Internal Error", nil}               // An error occurred internally.
	CodeValidationFailed          = localCode{51, "Validation Failed", nil}            // Data validation failed.
	CodeDbOperationError          = localCode{52, "Database Operation Error", nil}     // Database operation error.
	CodeInvalidParameter          = localCode{53, "Invalid Parameter", nil}            // Invalid parameter for the operation.
	CodeMissingParameter          = localCode{54, "Missing Parameter", nil}            // Missing parameter for the operation.
	CodeInvalidOperation          = localCode{55, "Invalid Operation", nil}            // Operation not valid.
	CodeInvalidConfiguration      = localCode{56, "Invalid Configuration", nil}        // Invalid configuration.
	CodeMissingConfiguration      = localCode{57, "Missing Configuration", nil}        // Missing configuration.
	CodeNotImplemented            = localCode{58, "Not Implemented", nil}              // Operation not implemented yet.
	CodeNotSupported              = localCode{59, "Not Supported", nil}                // Operation not supported.
	CodeOperationFailed           = localCode{60, "Operation Failed", nil}             // Operation failed.
	CodeNotAuthorized             = localCode{61, "Not Authorized", nil}               // Not authorized.
	CodeSecurityReason            = localCode{62, "Security Reason", nil}              // Security-related error.
	CodeServerBusy                = localCode{63, "Server Is Busy", nil}               // Server is busy.
	CodeUnknown                   = localCode{64, "Unknown Error", nil}                // Unknown error.
	CodeNotFound                  = localCode{65, "Not Found", nil}                    // Resource not found.
	CodeInvalidRequest            = localCode{66, "Invalid Request", nil}              // Invalid request.
	CodeNecessaryPackageNotImport = localCode{67, "Necessary Package Not Import", nil} // Required package not imported.
	CodeInternalPanic             = localCode{68, "Internal Panic", nil}               // Internal panic occurred.
	CodeBusinessValidationFailed  = localCode{300, "Business Validation Failed", nil}  // Business validation failed.
)

// New creates and returns a new error code instance.
func New(code int, message string, detail interface{}) Code {
	return localCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode creates a new error code by extending an existing one.
// It uses the code and message from the given `code`, but replaces the detail with the provided `detail`.
func WithCode(code Code, detail interface{}) Code {
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
