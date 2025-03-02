// Copyright (c) 2025 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package deepcopy provides functionality for creating deep copies of Go data structures
// using reflection. It can handle complex nested structures like maps, slices, and structs,
// as well as basic types.
//
// This package is adapted from: https://github.com/mohae/deepcopy
package deepcopy

import (
	"reflect"
	"time"
)

// Interface defines a method for custom deep copy implementation.
// Types that implement this interface will have control over their own deep copy process.
type Interface interface {
	// DeepCopy returns a deep copy of the implementing object.
	DeepCopy() interface{}
}

// Copy creates a deep copy of the provided source value and returns the copy
// as an interface{}.
//
// The returned value will need to be type-asserted to the correct type by the caller.
// For nil inputs, nil is returned.
//
// Primitive types (numbers, strings, booleans) are returned as-is since they are already
// passed by value in Go.
//
// For complex types (structs, maps, slices, etc.), a recursive deep copy is performed.
func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	// Handle primitive types and types that implement Interface directly
	switch src := src.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128,
		string, bool:
		// Primitive types are returned as-is (they're already passed by value)
		return src

	case Interface:
		// Use the type's custom deep copy implementation
		return src.DeepCopy()

	default:
		// For other types, use reflection to create a deep copy
		original := reflect.ValueOf(src)

		// Create a new value of the same type as the original
		dst := reflect.New(original.Type()).Elem()

		// Recursively copy the original value to the new value
		copyRecursive(original, dst)

		// Return the new value as an interface{}
		return dst.Interface()
	}
}

// copyRecursive handles the recursive copying of reflected values.
// It supports various kinds of types including pointers, interfaces, structs,
// slices, and maps.
//
// Parameters:
//   - original: reflect.Value of the source object
//   - cpy: reflect.Value of the destination where the copy will be stored
func copyRecursive(original, cpy reflect.Value) {
	// Check if the original implements Interface for custom copying
	if original.CanInterface() && original.IsValid() && !original.IsZero() {
		if copier, ok := original.Interface().(Interface); ok {
			cpy.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	// Handle based on the kind of value being copied
	switch original.Kind() {
	case reflect.Ptr:
		copyPointer(original, cpy)

	case reflect.Interface:
		copyInterface(original, cpy)

	case reflect.Struct:
		copyStruct(original, cpy)

	case reflect.Slice:
		copySlice(original, cpy)

	case reflect.Map:
		copyMap(original, cpy)

	default:
		// For all other types (including primitives when they reach here),
		// just set the original value directly
		cpy.Set(original)
	}
}

// copyPointer handles copying of pointer values.
func copyPointer(original, cpy reflect.Value) {
	// Get the value that the pointer points to
	originalValue := original.Elem()

	// If it isn't valid (nil pointer), return
	if !originalValue.IsValid() {
		return
	}

	// Create a new pointer of the same type
	cpy.Set(reflect.New(originalValue.Type()))

	// Recursively copy the pointed-to value
	copyRecursive(originalValue, cpy.Elem())
}

// copyInterface handles copying of interface values.
func copyInterface(original, cpy reflect.Value) {
	// If this is a nil interface, don't do anything
	if original.IsNil() {
		return
	}

	// Get the concrete value in the interface
	originalValue := original.Elem()

	// Create a new value of the same type
	copyValue := reflect.New(originalValue.Type()).Elem()

	// Recursively copy the concrete value
	copyRecursive(originalValue, copyValue)

	// Set the copied value to the destination interface
	cpy.Set(copyValue)
}

// copyStruct handles copying of struct values with special treatment for time.Time.
func copyStruct(original, cpy reflect.Value) {
	// Special case for time.Time as it contains unexported fields
	if t, ok := original.Interface().(time.Time); ok {
		cpy.Set(reflect.ValueOf(t))
		return
	}

	// Copy each exported field of the struct
	for i := 0; i < original.NumField(); i++ {
		// Skip unexported fields
		if original.Type().Field(i).PkgPath != "" {
			continue
		}

		// Recursively copy the field value
		copyRecursive(original.Field(i), cpy.Field(i))
	}
}

// copySlice handles copying of slice values.
func copySlice(original, cpy reflect.Value) {
	if original.IsNil() {
		return
	}

	// Create a new slice with the same length and capacity
	cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))

	// Copy each element of the slice
	for i := 0; i < original.Len(); i++ {
		copyRecursive(original.Index(i), cpy.Index(i))
	}
}

// copyMap handles copying of map values.
func copyMap(original, cpy reflect.Value) {
	if original.IsNil() {
		return
	}

	// Create a new map of the same type
	cpy.Set(reflect.MakeMap(original.Type()))

	// Copy each key-value pair in the map
	for _, key := range original.MapKeys() {
		// Get the original value for this key
		originalValue := original.MapIndex(key)

		// Create a new value of the same type
		copyValue := reflect.New(originalValue.Type()).Elem()

		// Recursively copy the value
		copyRecursive(originalValue, copyValue)

		// Deep copy the key as well (keys could be complex types)
		copyKey := Copy(key.Interface())

		// Set the copied key-value pair in the new map
		cpy.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
	}
}
