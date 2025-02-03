// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package empty provides utilities for checking if a value is empty or nil.
package empty

import (
	"reflect"
	"time"

	"github.com/focela/aegis/internal/core/reflection"
)

// Interfaces for type assertions.
type (
	Stringer interface {
		String() string
	}

	InterfacesProvider interface {
		Interfaces() []interface{}
	}

	MapConverter interface {
		MapStrAny() map[string]interface{}
	}

	TimeProvider interface {
		Date() (year int, month time.Month, day int)
		IsZero() bool
	}
)

// IsEmpty checks if the given `value` is empty.
// Empty values include: 0, nil, false, "", empty slices, maps, and channels.
// If `traceSource` is true, it traces through pointer references.
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	// First, check common primitive types to enhance performance before using reflection.
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
		// Use reflection for complex types.
		var rv reflect.Value
		if v, ok := value.(reflect.Value); ok {
			rv = v
		} else {
			rv = reflect.ValueOf(value)
			if IsNil(rv) {
				return true
			}

			// Check for common interfaces.
			if f, ok := value.(TimeProvider); ok {
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
			if f, ok := value.(InterfacesProvider); ok {
				if f == nil {
					return true
				}
				return len(f.Interfaces()) == 0
			}
			if f, ok := value.(MapConverter); ok {
				if f == nil {
					return true
				}
				return len(f.MapStrAny()) == 0
			}
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

		default:
			return false
		}
	}
}

// IsNil checks whether a given `value` is nil, especially for interface{} types.
// If `traceSource` is true, it traces through pointer references.
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
