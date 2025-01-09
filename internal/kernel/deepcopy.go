// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package kernel provides core utilities and foundational components for the Aegis framework.
package kernel

import (
	"reflect"
	"time"
)

// Interface defines a custom deep copy behavior.
type Interface interface {
	// DeepCopy creates a deep copy of the object.
	DeepCopy() interface{}
}

// Copy creates a deep copy of the given object and returns the copy.
// If the object implements the Interface, its DeepCopy method will be used.
func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	// Handle primitive types and basic type assertion.
	switch r := src.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128,
		string, bool:
		return r
	default:
		if v, ok := src.(Interface); ok {
			return v.DeepCopy()
		}

		original := reflect.ValueOf(src)
		dst := reflect.New(original.Type()).Elem()

		// Perform recursive deep copy.
		copyRecursive(original, dst)
		return dst.Interface()
	}
}

// copyRecursive performs the recursive deep copy operation.
func copyRecursive(original, cpy reflect.Value) {
	// Use custom DeepCopy if the type implements Interface.
	if original.IsValid() && original.CanInterface() && !original.IsZero() {
		if copier, ok := original.Interface().(Interface); ok {
			cpy.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	// Handle specific types based on kind.
	switch original.Kind() {
	case reflect.Ptr:
		originalValue := original.Elem()
		if !originalValue.IsValid() {
			return
		}
		cpy.Set(reflect.New(originalValue.Type()))
		copyRecursive(originalValue, cpy.Elem())
	case reflect.Interface:
		if original.IsNil() {
			return
		}
		originalValue := original.Elem()
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue)
		cpy.Set(copyValue)
	case reflect.Struct:
		// Special case for time.Time
		if t, ok := original.Interface().(time.Time); ok {
			cpy.Set(reflect.ValueOf(t))
			return
		}
		for i := 0; i < original.NumField(); i++ {
			if original.Type().Field(i).PkgPath != "" {
				continue
			}
			copyRecursive(original.Field(i), cpy.Field(i))
		}
	case reflect.Slice:
		if original.IsNil() {
			return
		}
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			copyRecursive(original.Index(i), cpy.Index(i))
		}
	case reflect.Map:
		if original.IsNil() {
			return
		}
		cpy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(originalValue, copyValue)
			copyKey := Copy(key.Interface())
			cpy.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}
	default:
		cpy.Set(original)
	}
}
