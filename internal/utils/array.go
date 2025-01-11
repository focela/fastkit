// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"reflect"
)

// IsArray checks whether the given value is an array or slice.
//
// Parameters:
// - value: The value to be checked.
//
// Returns:
// - true if the value is an array or slice, false otherwise.
//
// Note:
//   - This function handles both direct values and pointers. If a pointer is provided,
//     it dereferences the pointer to determine the underlying type.
func IsArray(value interface{}) bool {
	rv := reflect.ValueOf(value)
	valueKind := rv.Kind()

	// Dereference pointer to get the underlying type.
	if valueKind == reflect.Ptr {
		rv = rv.Elem()
		valueKind = rv.Kind()
	}

	// Check if the value is an array or slice.
	switch valueKind {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}
