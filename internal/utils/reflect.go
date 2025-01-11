// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"reflect"
)

// CanCallIsNil determines if the reflect.Value `v` can safely call IsNil without causing a panic.
//
// Parameters:
// - v: An interface that can be a reflect.Value or any other value.
//
// Returns:
// - true if IsNil can be safely called on `v`, false otherwise.
//
// Notes:
// - This function ensures safe usage of reflect.Value.IsNil by checking if the kind of `v` is valid for IsNil.
// - Applicable kinds are: Interface, Chan, Func, Map, Ptr, Slice, and UnsafePointer.
func CanCallIsNil(v interface{}) bool {
	rv, ok := v.(reflect.Value)
	if !ok {
		return false
	}

	// Check if the kind is valid for calling IsNil.
	switch rv.Kind() {
	case reflect.Interface, reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return true
	default:
		return false
	}
}
