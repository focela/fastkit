// Copyright (c) 2025 Focela Technologies.
// This software is provided "as is", without any warranty.
// Licensed under the MIT License – see LICENSE file for details.

// Package empty provides functions for checking empty/nil variables.
// It offers optimized checks for various Go types including primitives, collections, and custom interfaces.
package empty

import (
	"reflect"
	"time"

	"github.com/focela/aegis/internal/corelib/reflection"
)

// Interface definitions for type assertions used in emptiness checks.
// These interfaces allow the package to efficiently handle custom types.

// Stringer defines an interface for types that can convert to string.
type Stringer interface {
	String() string
}

// InterfaceProvider defines an interface for types that can provide a slice of interfaces.
type InterfaceProvider interface {
	Interfaces() []interface{}
}

// MapStringConverter defines an interface for types that can convert to map[string]interface{}.
type MapStringConverter interface {
	MapStrAny() map[string]interface{}
}

// TimeInfo defines an interface compatible with time.Time.
type TimeInfo interface {
	Date() (year int, month time.Month, day int)
	IsZero() bool
}

// IsEmpty checks whether the given `value` is empty.
//
// Returns true if `value` is one of:
// - nil
// - zero numeric value (0, 0.0)
// - false boolean
// - empty string
// - empty collection (slice/map/chan with length 0)
// - struct with all empty fields
// - nil interface or pointer
//
// The parameter `traceSource` is used for tracing to the source variable if given `value` is a pointer
// that also points to a pointer. It returns true if the source is empty when `traceSource` is true.
// Note that it might use reflection which affects performance.
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	// First check the variable as common types using type assertion to enhance performance,
	// before falling back to reflection.
	switch result := value.(type) {
	case int:
		return result == 0
	case int8:
		return result == 0
	case int16:
		return result == 0
	case int32:
		return result == 0
	case int64:
		return result == 0
	case uint:
		return result == 0
	case uint8:
		return result == 0
	case uint16:
		return result == 0
	case uint32:
		return result == 0
	case uint64:
		return result == 0
	case float32:
		return result == 0
	case float64:
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

	default:
		// Fall back to reflection for other types
		var rv reflect.Value
		if v, ok := value.(reflect.Value); ok {
			rv = v
		} else {
			rv = reflect.ValueOf(value)
			if IsNil(rv) {
				return true
			}

			// Check common interfaces
			if f, ok := value.(TimeInfo); ok {
				if f == (*time.Time)(nil) {
					return true
				}
				return f.IsZero()
			}
			if f, ok := value.(Stringer); ok {
				if f == nil {
					return true
				}
				return f.String() == ""
			}
			if f, ok := value.(InterfaceProvider); ok {
				if f == nil {
					return true
				}
				return len(f.Interfaces()) == 0
			}
			if f, ok := value.(MapStringConverter); ok {
				if f == nil {
					return true
				}
				return len(f.MapStrAny()) == 0
			}
		}

		// Handle different value kinds using reflection
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

		case reflect.Struct:
			// For structs, check all fields recursively
			var fieldValueInterface interface{}
			for i := 0; i < rv.NumField(); i++ {
				fieldValueInterface, _ = reflection.ValueToInterface(rv.Field(i))
				if !IsEmpty(fieldValueInterface) {
					return false
				}
			}
			return true

		case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
			return rv.Len() == 0

		case reflect.Ptr:
			if len(traceSource) > 0 && traceSource[0] {
				return IsEmpty(rv.Elem())
			}
			return rv.IsNil()

		case reflect.Func, reflect.Interface, reflect.UnsafePointer:
			return rv.IsNil()

		case reflect.Invalid:
			return true

		default:
			return false
		}
	}
}

// IsNil checks whether the given `value` is nil.
// It's particularly useful for interface{} type values which can't be directly compared with nil.
//
// Parameter `traceSource` is used for tracing to the source variable if given `value` is a pointer
// that also points to a pointer. It returns true if the source is nil when `traceSource` is true.
// Note that it uses reflection which affects performance.
func IsNil(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	// Convert to reflect.Value if needed
	var rv reflect.Value
	if v, ok := value.(reflect.Value); ok {
		rv = v
	} else {
		rv = reflect.ValueOf(value)
	}

	// Check based on value kind
	switch rv.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return !rv.IsValid() || rv.IsNil()

	case reflect.Ptr:
		if len(traceSource) > 0 && traceSource[0] {
			// Trace through pointer chain to the source
			for rv.Kind() == reflect.Ptr {
				rv = rv.Elem()
			}
			if !rv.IsValid() {
				return true
			}
			if rv.Kind() == reflect.Ptr {
				return rv.IsNil()
			}
		} else {
			return !rv.IsValid() || rv.IsNil()
		}

	default:
		return false
	}

	return false
}
