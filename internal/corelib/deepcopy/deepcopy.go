// Copyright (c) 2025 Focela Technologies.
// This software is provided "as is", without any warranty.
// Licensed under the MIT License – see LICENSE file for details.

// Package deepcopy provides functionality to create deep copies of Go values using reflection.
// It supports primitive types, structs, slices, maps, pointers, and interfaces.
//
// This package is based on: https://github.com/mohae/deepcopy
package deepcopy

import (
	"reflect"
	"time"
)

// Interface allows types to implement their own deep copy logic.
// Types that implement this interface will be copied using their DeepCopy method
// rather than the generic reflection-based approach.
type Interface interface {
	// DeepCopy returns a deep copy of the implementing object.
	DeepCopy() interface{}
}

// Copy creates a deep copy of the provided value and returns it as an interface{}.
// The caller is responsible for type-asserting the result to the correct type.
//
// Supported types:
// - Basic types (numeric, string, bool) are copied by value
// - Types implementing the deepcopy.Interface use their DeepCopy method
// - Pointers, interfaces, structs, slices, and maps are recursively deep-copied
// - time.Time values are copied by value
//
// For nil inputs, nil is returned.
func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	// Copy by type assertion for basic types that don't need deep copying
	switch src := src.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128,
		string, bool:
		return src

	default:
		// Check if the type implements the DeepCopy interface
		if v, ok := src.(Interface); ok {
			return v.DeepCopy()
		}

		// Use reflection for other types
		original := reflect.ValueOf(src)
		dst := reflect.New(original.Type()).Elem()

		// Recursively copy the original value
		copyRecursive(original, dst)

		return dst.Interface()
	}
}

// copyRecursive performs deep copying of values using reflection.
// It handles various types including pointers, interfaces, structs, slices, and maps.
// The function modifies the destination value (cpy) to be a deep copy of the original.
func copyRecursive(original, cpy reflect.Value) {
	// Check if the original value implements deepcopy.Interface
	if original.CanInterface() && original.IsValid() && !original.IsZero() {
		if copier, ok := original.Interface().(Interface); ok {
			cpy.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	// Handle different types based on their Kind
	switch original.Kind() {
	case reflect.Ptr:
		// Handle pointer types
		originalValue := original.Elem()

		// If the pointer is nil or points to an invalid value, return
		if !originalValue.IsValid() {
			return
		}

		// Create a new pointer of the same type and recursively copy its value
		cpy.Set(reflect.New(originalValue.Type()))
		copyRecursive(originalValue, cpy.Elem())

	case reflect.Interface:
		// Handle interface types
		if original.IsNil() {
			return
		}

		// Get the concrete value and recursively copy it
		originalValue := original.Elem()
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue)
		cpy.Set(copyValue)

	case reflect.Struct:
		// Special handling for time.Time
		if t, ok := original.Interface().(time.Time); ok {
			cpy.Set(reflect.ValueOf(t))
			return
		}

		// For other structs, copy each exported field
		for i := 0; i < original.NumField(); i++ {
			// Skip unexported fields (those with a non-empty PkgPath)
			if original.Type().Field(i).PkgPath != "" {
				continue
			}
			copyRecursive(original.Field(i), cpy.Field(i))
		}

	case reflect.Slice:
		if original.IsNil() {
			return
		}

		// Create a new slice with the same length and capacity
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))

		// Copy each element
		for i := 0; i < original.Len(); i++ {
			copyRecursive(original.Index(i), cpy.Index(i))
		}

	case reflect.Map:
		if original.IsNil() {
			return
		}

		// Create a new map of the same type
		cpy.Set(reflect.MakeMap(original.Type()))

		// Copy each key-value pair
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(originalValue, copyValue)

			// Also deep copy the map keys
			copyKey := Copy(key.Interface())
			cpy.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}

	default:
		// For other types, set the value directly
		cpy.Set(original)
	}
}
