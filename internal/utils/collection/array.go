// Copyright (c) 2025 Focela Technologies.
// This software is provided "as is", without any warranty.
// Licensed under the MIT License – see LICENSE file for details.

// Package collection provides utility functions for working with collections in Go.
// It includes functions for inspecting and manipulating arrays, slices, maps and other collection types.
package collection

import (
	"reflect"
)

// IsArray checks whether the given value is an array or slice.
// It handles both direct values and pointers to arrays/slices.
//
// Parameters:
//   - value: The value to check, can be of any type.
//
// Returns:
//   - true if the value is an array or slice (or a pointer to one).
//   - false otherwise.
//
// Note that this function uses reflection, which may affect performance in critical paths.
func IsArray(value interface{}) bool {
	// Get the reflection value
	rv := reflect.ValueOf(value)
	kind := rv.Kind()

	// Dereference pointer if needed
	if kind == reflect.Ptr {
		rv = rv.Elem()
		kind = rv.Kind()
	}

	// Check if the kind is array or slice
	return kind == reflect.Array || kind == reflect.Slice
}
