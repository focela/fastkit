// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package empty provides functions for checking empty/nil variables.
package empty

import (
	"reflect"
	"time"

	"github.com/focela/aid/internal/reflection"
)

// StringProvider is used for type assertion for String().
type StringProvider interface {
	String() string
}

// InterfaceProvider is used for type assertion for Interfaces().
type InterfaceProvider interface {
	Interfaces() []interface{}
}

// MapProvider is the interface for converting struct parameter to map.
type MapProvider interface {
	MapStrAny() map[string]interface{}
}

// Timer is used for type assertion for time-related functions.
type Timer interface {
	Date() (year int, month time.Month, day int)
	IsZero() bool
}

// IsEmpty checks whether the given `value` is empty.
// It returns true if `value` is in: 0, nil, false, "", len(slice/map/chan) == 0,
// or else it returns false.
//
// The parameter `traceSource` is used for tracing to the source variable if the given `value` is a pointer
// that also points to a pointer. It returns true if the source is empty when `traceSource` is true.
// Note that it might use reflection which affects performance a little.
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}
	// It firstly checks the variable as common types using assertion to enhance the performance,
	// and then using reflection.
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
		// Finally, using reflect.
		var rv reflect.Value
		if v, ok := value.(reflect.Value); ok {
			rv = v
		} else {
			// =========================
			// Common interfaces checks.
			// =========================
			if f, ok := value.(Timer); ok {
				if f == (*time.Time)(nil) {
					return true
				}
				return f.IsZero()
			}
			if f, ok := value.(StringProvider); ok {
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
			if f, ok := value.(MapProvider); ok {
				if f == nil {
					return true
				}
				return len(f.MapStrAny()) == 0
			}

			rv = reflect.ValueOf(value)
		}

		switch rv.Kind() {
		case reflect.Bool:
			return !rv.Bool()

		case
			reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64:
			return rv.Int() == 0

		case
			reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64,
			reflect.Uintptr:
			return rv.Uint() == 0

		case
			reflect.Float32,
			reflect.Float64:
			return rv.Float() == 0

		case reflect.String:
			return rv.Len() == 0

		case reflect.Struct:
			var fieldValueInterface interface{}
			for i := 0; i < rv.NumField(); i++ {
				fieldValueInterface, _ = reflection.ValueToInterface(rv.Field(i))
				if !IsEmpty(fieldValueInterface) {
					return false
				}
			}
			return true

		case
			reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Array:
			return rv.Len() == 0

		case reflect.Ptr:
			if len(traceSource) > 0 && traceSource[0] {
				return IsEmpty(rv.Elem())
			}
			return rv.IsNil()

		case
			reflect.Func,
			reflect.Interface,
			reflect.UnsafePointer:
			return rv.IsNil()

		case reflect.Invalid:
			return true
		}
	}
	return false
}

// IsNil checks whether given `value` is nil, especially for interface{} type value.
// Parameter `traceSource` is used for tracing to the source variable if given `value` is type of pointer
// that also points to a pointer. It returns nil if the source is nil when `traceSource` is true.
// Note that it might use reflection which affects performance a little.
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
	}
	return false
}
