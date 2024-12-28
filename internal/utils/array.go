// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"reflect"
)

// IsArray checks if the given value is an array or slice.
func IsArray(value interface{}) bool {
	// Get the reflection value of the input
	rv := reflect.ValueOf(value)

	// Handle pointer type by dereferencing
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}

	// Check if the value is an array or slice
	return rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice
}
