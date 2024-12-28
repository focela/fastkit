// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"reflect"
)

// CanCallIsNil checks if a reflect.Value can safely call IsNil without causing a panic.
// This function ensures compatibility with types that support the IsNil method.
//
// Supported kinds: Interface, Chan, Func, Map, Ptr, Slice, UnsafePointer.
//
// Parameters:
// - v: The value to be checked.
//
// Returns:
// - true if IsNil can be safely called on the given reflect.Value.
// - false otherwise.
func CanCallIsNil(v interface{}) bool {
	// Ensure the input is a reflect.Value
	rv, ok := v.(reflect.Value)
	if !ok {
		return false
	}

	// Check if the kind supports IsNil
	switch rv.Kind() {
	case reflect.Interface,
		reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice,
		reflect.UnsafePointer:
		return true
	default:
		return false
	}
}
