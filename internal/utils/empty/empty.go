// Copyright (c) 2025 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package empty provides utilities for determining if values are empty or nil.
// It supports checking various Go types including primitives, collections, interfaces,
// structs and custom types that implement specific interfaces.
package empty

import (
	"reflect"
	"time"

	"github.com/focela/aegis/internal/core/introspection"
)

// Stringer is the interface for types that provide a string representation.
// Used to check emptiness for types implementing the String() method.
type Stringer interface {
	String() string
}

// InterfacesProvider is the interface for types that can provide a slice of interfaces.
// Used to check emptiness for types implementing the Interfaces() method.
type InterfacesProvider interface {
	Interfaces() []interface{}
}

// MapConverter is the interface for types that can be converted to map[string]interface{}.
// Used for checking emptiness of types that provide MapStrAny() method.
type MapConverter interface {
	MapStrAny() map[string]interface{}
}

// TimeInfo is the interface for time-related types with date components and zero checking.
// Used primarily for handling time.Time and similar custom time types.
type TimeInfo interface {
	Date() (year int, month time.Month, day int)
	IsZero() bool
}

// IsEmpty checks whether given `value` is empty.
// It returns true if `value` is in: 0, nil, false, "", len(slice/map/chan) == 0,
// or else it returns false.
//
// For struct types, it checks if all fields are empty. For custom types, it checks
// if they implement certain interfaces like iString, iTime, etc., and uses those
// to determine emptiness.
//
// Parameters:
//   - value: The value to check
//   - traceSource: Optional. When true and value is a pointer, it will follow pointers
//     to check if the ultimately referenced value is empty
//
// Returns:
//   - true if the value is considered empty
//   - false otherwise
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	// First check common types using type assertion for better performance
	isEmpty, handled := checkCommonTypesEmpty(value)
	if handled {
		return isEmpty
	}

	// If not a common type, use reflection for more complex types
	return checkReflectionEmpty(value, traceSource...)
}

// checkCommonTypesEmpty performs emptiness check on common types using type assertions.
// This approach is faster than using reflection for these types.
//
// Returns:
//   - isEmpty: whether the value is empty
//   - handled: whether the type was handled by this function
func checkCommonTypesEmpty(value interface{}) (isEmpty bool, handled bool) {
	switch result := value.(type) {
	// Numeric types
	case int:
		return result == 0, true
	case int8:
		return result == 0, true
	case int16:
		return result == 0, true
	case int32:
		return result == 0, true
	case int64:
		return result == 0, true
	case uint:
		return result == 0, true
	case uint8:
		return result == 0, true
	case uint16:
		return result == 0, true
	case uint32:
		return result == 0, true
	case uint64:
		return result == 0, true
	case float32:
		return result == 0, true
	case float64:
		return result == 0, true

	// Boolean
	case bool:
		return !result, true

	// String
	case string:
		return result == "", true

	// Collections
	case []byte:
		return len(result) == 0, true
	case []rune:
		return len(result) == 0, true
	case []int:
		return len(result) == 0, true
	case []string:
		return len(result) == 0, true
	case []float32:
		return len(result) == 0, true
	case []float64:
		return len(result) == 0, true
	case map[string]interface{}:
		return len(result) == 0, true

	// Not handled by this function
	default:
		return false, false
	}
}

// checkReflectionEmpty checks emptiness using reflection for complex types
// not handled by checkCommonTypesEmpty.
func checkReflectionEmpty(value interface{}, traceSource ...bool) bool {
	// Get reflect.Value from the interface
	var rv reflect.Value
	if v, ok := value.(reflect.Value); ok {
		rv = v
	} else {
		rv = reflect.ValueOf(value)
		if IsNil(rv) {
			return true
		}

		// Check for special interface implementations
		if isEmpty, handled := checkInterfaceImplementationsEmpty(value); handled {
			return isEmpty
		}
	}

	// Handle different kinds of values using reflection
	return checkReflectionKindEmpty(rv, traceSource...)
}

// checkInterfaceImplementationsEmpty checks emptiness for types implementing special interfaces.
func checkInterfaceImplementationsEmpty(value interface{}) (isEmpty bool, handled bool) {
	// Check for time-related interface
	if f, ok := value.(TimeInfo); ok {
		if f == (*time.Time)(nil) {
			return true, true
		}
		return f.IsZero(), true
	}

	// Check for string interface
	if f, ok := value.(Stringer); ok {
		if f == nil {
			return true, true
		}
		return f.String() == "", true
	}

	// Check for interfaces interface
	if f, ok := value.(InterfacesProvider); ok {
		if f == nil {
			return true, true
		}
		return len(f.Interfaces()) == 0, true
	}

	// Check for map interface
	if f, ok := value.(MapConverter); ok {
		if f == nil {
			return true, true
		}
		return len(f.MapStrAny()) == 0, true
	}

	return false, false
}

// checkReflectionKindEmpty handles emptiness checks based on reflect.Kind.
func checkReflectionKindEmpty(rv reflect.Value, traceSource ...bool) bool {
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
		return isStructEmpty(rv)

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

// isStructEmpty checks if all fields in a struct are empty.
func isStructEmpty(rv reflect.Value) bool {
	for i := 0; i < rv.NumField(); i++ {
		fieldValue, _ := introspection.ValueToInterface(rv.Field(i))
		if !IsEmpty(fieldValue) {
			return false
		}
	}
	return true
}

// IsNil checks whether given `value` is nil.
// This function is particularly useful for checking interfaces and other reference types.
//
// Parameters:
//   - value: The value to check
//   - traceSource: Optional. When true and value is a pointer to a pointer,
//     it will recursively follow the pointers to their source
//
// Returns:
//   - true if the value is nil or invalid
//   - false otherwise
func IsNil(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	var rv reflect.Value
	if v, ok := value.(reflect.Value); ok {
		rv = v
	} else {
		rv = reflect.ValueOf(value)
	}

	switch rv.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return !rv.IsValid() || rv.IsNil()

	case reflect.Ptr:
		if len(traceSource) > 0 && traceSource[0] {
			return tracePointerToSource(rv)
		}
		return !rv.IsValid() || rv.IsNil()

	default:
		return false
	}
}

// tracePointerToSource follows pointer chains to their source and checks if the source is nil.
func tracePointerToSource(rv reflect.Value) bool {
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if !rv.IsValid() {
		return true
	}

	if rv.Kind() == reflect.Ptr {
		return rv.IsNil()
	}

	return false
}
