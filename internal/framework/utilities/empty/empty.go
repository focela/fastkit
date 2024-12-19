// Copyright (c) 2024 Focela Technologies
// Internal Use Only - Unauthorized use prohibited
// Contact: legal@focela.com

// Package empty provides functions for checking empty or nil variables.
package empty

import (
	"reflect"
	"time"

	"github.com/focela/altura/internal/framework/reflection"
)

// Stringer defines an interface with a String() method.
type Stringer interface {
	String() string
}

// InterfacesProvider defines an interface with a method returning a slice of interfaces.
type InterfacesProvider interface {
	Interfaces() []interface{}
}

// MapStringAnyProvider defines an interface for converting struct parameters to a map.
type MapStringAnyProvider interface {
	MapStrAny() map[string]interface{}
}

// Timer defines a subset of time-related methods.
type Timer interface {
	Date() (year int, month time.Month, day int)
	IsZero() bool
}

// IsEmpty checks if a given value is empty.
// Returns true if value is nil, 0, false, "", or a collection with length 0.
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	// Optimize for common types using type assertions.
	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		return v == 0
	case uint, uint8, uint16, uint32, uint64:
		return v == 0
	case float32, float64:
		return v == 0
	case bool:
		return !v
	case string:
		return v == ""
	case []byte:
		return len(v) == 0
	case []rune:
		return len(v) == 0
	case map[string]interface{}:
		return len(v) == 0
	}

	// Fallback to reflection for other types.
	rv := reflect.ValueOf(value)
	if IsNil(rv) {
		return true
	}

	// Check common interfaces.
	if t, ok := value.(Timer); ok && (t == nil || t.IsZero()) {
		return true
	}
	if s, ok := value.(Stringer); ok && (s == nil || s.String() == "") {
		return true
	}
	if i, ok := value.(InterfacesProvider); ok && (i == nil || len(i.Interfaces()) == 0) {
		return true
	}
	if m, ok := value.(MapStringAnyProvider); ok && (m == nil || len(m.MapStrAny()) == 0) {
		return true
	}

	// Use reflection for advanced checks.
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
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return rv.Len() == 0
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			fieldValue, _ := reflection.ValueToInterface(rv.Field(i))
			if !IsEmpty(fieldValue) {
				return false
			}
		}
		return true
	case reflect.Ptr:
		if len(traceSource) > 0 && traceSource[0] && rv.IsValid() {
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

// IsNil checks if a given value is nil.
// Supports complex cases such as pointers to pointers.
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
			return !rv.IsValid() || rv.IsNil()
		}
		return rv.IsNil()
	default:
		return false
	}
}
