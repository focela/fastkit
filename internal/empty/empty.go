// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package empty provides utility functions to check emptiness and nil values for various types.
package empty

import (
	"reflect"
	"time"

	"github.com/focela/loom/internal/reflection"
)

// Interface definitions for type assertions.
type (
	Stringer interface {
		String() string
	}

	InterfaceProvider interface {
		Interfaces() []interface{}
	}

	MapStringAnyProvider interface {
		MapStrAny() map[string]interface{}
	}

	TimeProvider interface {
		Date() (year int, month time.Month, day int)
		IsZero() bool
	}
)

// IsEmpty checks whether the given `value` is empty.
// It evaluates emptiness for basic types, slices, maps, structs, and pointers.
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	// Check common types directly.
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

	// Use reflection for complex types.
	rv := reflect.ValueOf(value)
	if IsNil(rv) {
		return true
	}

	// Check custom interfaces.
	if f, ok := value.(TimeProvider); ok {
		return f.IsZero()
	}
	if f, ok := value.(Stringer); ok {
		return f.String() == ""
	}
	if f, ok := value.(InterfaceProvider); ok {
		return len(f.Interfaces()) == 0
	}
	if f, ok := value.(MapStringAnyProvider); ok {
		return len(f.MapStrAny()) == 0
	}

	// Check reflection kinds.
	switch rv.Kind() {
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.String:
		return rv.Len() == 0
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			fieldValue, _ := reflection.ValueToInterface(rv.Field(i))
			if !IsEmpty(fieldValue) {
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

// IsNil checks whether the given `value` is nil.
// It evaluates nil values for various types, including pointers, slices, maps, and interfaces.
func IsNil(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Ptr && len(traceSource) > 0 && traceSource[0] {
		for rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
		if !rv.IsValid() {
			return true
		}
		if rv.Kind() == reflect.Ptr {
			return rv.IsNil()
		}
	}

	switch rv.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return !rv.IsValid() || rv.IsNil()
	case reflect.Ptr:
		return !rv.IsValid() || rv.IsNil()
	default:
		return false
	}
}
