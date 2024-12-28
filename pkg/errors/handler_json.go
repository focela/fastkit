// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package errors provides rich functionalities to manipulate errors.
package errors

import (
	"encoding/json"
)

// MarshalJSON serializes the error into a JSON string representation.
// It implements the `json.Marshaler` interface to ensure compatibility with `json.Marshal`.
//
// The error message is wrapped in double quotes as a standard JSON string.
func (err Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(err.Error())
}
