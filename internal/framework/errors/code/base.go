// Copyright (c) 2024 Focela Technologies
// Internal Use Only - Unauthorized use prohibited
// Contact: legal@focela.com

package code

// Code defines the universal error code interface.
type Code interface {
	// Code returns the numeric value of the error code.
	Code() int

	// Message returns the brief message of the error code.
	Message() string

	// Detail returns additional details or metadata of the error code.
	Detail() interface{}
}

// Predefined error codes for common scenarios.
// Codes below 1000 are reserved for internal framework usage.
var (
	CodeNil                       = localCode{-1, "", nil}                             // No specific error code.
	CodeOK                        = localCode{0, "OK", nil}                            // Success.
	CodeInternalError             = localCode{50, "Internal Error", nil}               // Generic internal error.
	CodeValidationFailed          = localCode{51, "Validation Failed", nil}            // Data validation error.
	CodeDbOperationError          = localCode{52, "Database Operation Error", nil}     // Database operation error.
	CodeInvalidParameter          = localCode{53, "Invalid Parameter", nil}            // Invalid parameter for the operation.
	CodeMissingParameter          = localCode{54, "Missing Parameter", nil}            // Missing parameter for the operation.
	CodeInvalidOperation          = localCode{55, "Invalid Operation", nil}            // Operation is not valid in this context.
	CodeInvalidConfiguration      = localCode{56, "Invalid Configuration", nil}        // Configuration is invalid.
	CodeMissingConfiguration      = localCode{57, "Missing Configuration", nil}        // Required configuration is missing.
	CodeNotImplemented            = localCode{58, "Not Implemented", nil}              // Functionality is not yet implemented.
	CodeNotSupported              = localCode{59, "Not Supported", nil}                // Functionality is not supported.
	CodeOperationFailed           = localCode{60, "Operation Failed", nil}             // Operation could not complete successfully.
	CodeNotAuthorized             = localCode{61, "Not Authorized", nil}               // Unauthorized access.
	CodeSecurityReason            = localCode{62, "Security Reason", nil}              // Restricted due to security reasons.
	CodeServerBusy                = localCode{63, "Server Is Busy", nil}               // Server is busy; try later.
	CodeUnknown                   = localCode{64, "Unknown Error", nil}                // Unknown error occurred.
	CodeNotFound                  = localCode{65, "Not Found", nil}                    // Resource not found.
	CodeInvalidRequest            = localCode{66, "Invalid Request", nil}              // Request format is invalid.
	CodeNecessaryPackageNotImport = localCode{67, "Necessary Package Not Import", nil} // Missing required package.
	CodeInternalPanic             = localCode{68, "Internal Panic", nil}               // Panic occurred internally.
	CodeBusinessValidationFailed  = localCode{300, "Business Validation Failed", nil}  // Business validation error.
)

// New creates and returns a new error code.
func New(code int, message string, detail interface{}) Code {
	return localCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode creates and returns a new error code based on an existing Code.
func WithCode(code Code, detail interface{}) Code {
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
