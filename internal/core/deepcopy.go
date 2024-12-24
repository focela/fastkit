// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package core provides essential utilities and foundational tools for the application.
package core

import (
	"reflect"
	"time"
)

// Interface defines a contract for types that support deep copying.
type Interface interface {
	DeepCopy() interface{}
}

// Copy creates a deep copy of the provided source and returns it as an interface{}.
// The returned value will need to be type-asserted to the appropriate type.
func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	// Directly return basic types without further processing.
	switch r := src.(type) {
	case
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128,
		string,
		bool:
		return r

	default:
		if v, ok := src.(Interface); ok {
			return v.DeepCopy()
		}
		original := reflect.ValueOf(src)
		dst := reflect.New(original.Type()).Elem()
		copyRecursive(original, dst)
		return dst.Interface()
	}
}

// copyRecursive performs recursive copying of complex data structures.
// Supports pointers, interfaces, structs, slices, and maps.
func copyRecursive(original, cpy reflect.Value) {
	// Check for custom DeepCopy implementation.
	if original.CanInterface() && original.IsValid() && !original.IsZero() {
		if copier, ok := original.Interface().(Interface); ok {
			cpy.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

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
		cpy.Set(original)
	}
}

// copyPointer handles deep copying of pointer types.
func copyPointer(original, cpy reflect.Value) {
	if !original.Elem().IsValid() {
		return
	}
	cpy.Set(reflect.New(original.Elem().Type()))
	copyRecursive(original.Elem(), cpy.Elem())
}

// copyInterface handles deep copying of interface types.
func copyInterface(original, cpy reflect.Value) {
	if original.IsNil() {
		return
	}
	originalValue := original.Elem()
	copyValue := reflect.New(originalValue.Type()).Elem()
	copyRecursive(originalValue, copyValue)
	cpy.Set(copyValue)
}

// copyStruct handles deep copying of struct types.
func copyStruct(original, cpy reflect.Value) {
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
}

// copySlice handles deep copying of slice types.
func copySlice(original, cpy reflect.Value) {
	if original.IsNil() {
		return
	}
	cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
	for i := 0; i < original.Len(); i++ {
		copyRecursive(original.Index(i), cpy.Index(i))
	}
}

// copyMap handles deep copying of map types.
func copyMap(original, cpy reflect.Value) {
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
}
