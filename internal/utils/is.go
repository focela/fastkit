// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"reflect"
)

// IsNil checks whether the given value is nil.
//
// Parameters:
// - value: The input value to check.
//
// Returns:
// - true if the value is nil, false otherwise.
//
// Note:
// - This function is especially useful for interface{} type values, as it uses reflection to determine nil status.
func IsNil(value interface{}) bool {
	return IsItNil(value)
}

// IsEmpty checks whether the given value is considered empty.
//
// Parameters:
// - value: The input value to check.
//
// Returns:
// - true if the value is empty, false otherwise.
//
// Note:
// - A value is considered empty if it is 0, nil, false, "", or has a zero length (e.g., slice, map, or chan).
func IsEmpty(value interface{}) bool {
	return IsItEmpty(value)
}

// IsInt checks whether the given value is of integer type.
//
// Parameters:
// - value: The input value to check.
//
// Returns:
// - true if the value is an integer type, false otherwise.
func IsInt(value interface{}) bool {
	switch value.(type) {
	case int, *int, int8, *int8, int16, *int16, int32, *int32, int64, *int64:
		return true
	}
	return false
}

// IsUint checks whether the given value is of unsigned integer type.
//
// Parameters:
// - value: The input value to check.
//
// Returns:
// - true if the value is an unsigned integer type, false otherwise.
func IsUint(value interface{}) bool {
	switch value.(type) {
	case uint, *uint, uint8, *uint8, uint16, *uint16, uint32, *uint32, uint64, *uint64:
		return true
	}
	return false
}

// IsFloat checks whether the given value is of float type.
//
// Parameters:
// - value: The input value to check.
//
// Returns:
// - true if the value is a float type, false otherwise.
func IsFloat(value interface{}) bool {
	switch value.(type) {
	case float32, *float32, float64, *float64:
		return true
	}
	return false
}

// IsSlice checks whether the given value is a slice or array.
//
// Parameters:
// - value: The input value to check.
//
// Returns:
// - true if the value is a slice or array, false otherwise.
func IsSlice(value interface{}) bool {
	rv := reflect.ValueOf(value)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	return rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array
}

// IsMap checks whether the given value is a map.
//
// Parameters:
// - value: The input value to check.
//
// Returns:
// - true if the value is a map, false otherwise.
func IsMap(value interface{}) bool {
	rv := reflect.ValueOf(value)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	return rv.Kind() == reflect.Map
}

// IsStruct checks whether the given value is a struct.
//
// Parameters:
// - value: The input value to check.
//
// Returns:
// - true if the value is a struct, false otherwise.
func IsStruct(value interface{}) bool {
	rt := reflect.TypeOf(value)
	if rt == nil {
		return false
	}
	for rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	return rt.Kind() == reflect.Struct
}
