// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"reflect"
	"time"

	"github.com/focela/aegis/internal/kernel"
)

// Stringer defines an interface with a String() method.
type Stringer interface {
	String() string
}

// InterfaceProvider defines an interface for providing multiple interface values.
type InterfaceProvider interface {
	Interfaces() []interface{}
}

// Mapper defines an interface for converting a struct to a map.
type Mapper interface {
	MapStrAny() map[string]interface{}
}

// TimeHandler defines an interface for handling time-related operations.
type TimeHandler interface {
	Date() (year int, month time.Month, day int)
	IsZero() bool
}

// IsEmpty checks whether the given value is considered empty.
//
// Parameters:
// - value: The input value to check.
// - traceSource: Optional boolean indicating whether to trace pointers to the source value.
//
// Returns:
// - true if the value is empty, false otherwise.
//
// Note:
// - A value is considered empty if it is 0, nil, false, "", or has a zero length (e.g., slice, map, or chan).
// - This function may use reflection, which can impact performance.
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	// Handle common types for performance.
	switch result := value.(type) {
	case int, int8, int16, int32, int64:
		return result == 0
	case uint, uint8, uint16, uint32, uint64:
		return result == 0
	case float32, float64:
		return result == 0
	case bool:
		return !result
	case string:
		return result == ""
	case []byte:
		return len(result) == 0
	case []rune:
		return len(result) == 0
	case []int:
		return len(result) == 0
	case []string:
		return len(result) == 0
	case []float32:
		return len(result) == 0
	case []float64:
		return len(result) == 0
	case map[string]interface{}:
		return len(result) == 0
	}

	// Handle reflection for more complex types.
	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Ptr && rv.IsNil() {
		return true
	}

	switch rv.Kind() {
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.String:
		return rv.Len() == 0
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return rv.Len() == 0
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			fieldValue, _ := kernel.ValueToInterface(rv.Field(i))
			if !IsEmpty(fieldValue) {
				return false
			}
		}
		return true
	case reflect.Ptr:
		if len(traceSource) > 0 && traceSource[0] {
			return IsEmpty(rv.Elem())
		}
		return rv.IsNil()
	default:
		return false
	}
}

// IsNil checks whether the given value is nil.
//
// Parameters:
// - value: The input value to check.
// - traceSource: Optional boolean indicating whether to trace pointers to the source value.
//
// Returns:
// - true if the value is nil, false otherwise.
//
// Note:
// - This function works for interface{} types and uses reflection to determine if the value is nil.
func IsNil(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return !rv.IsValid() || rv.IsNil()
	case reflect.Ptr:
		if len(traceSource) > 0 && traceSource[0] {
			for rv.Kind() == reflect.Ptr {
				rv = rv.Elem()
			}
			return !rv.IsValid() || rv.Kind() == reflect.Ptr && rv.IsNil()
		}
		return !rv.IsValid() || rv.IsNil()
	default:
		return false
	}
}
