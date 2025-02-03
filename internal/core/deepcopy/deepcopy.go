// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package deepcopy provides utilities for deep copying data structures using reflection.
//
// This package is maintained from: https://github.com/mohae/deepcopy
package deepcopy

import (
	"reflect"
	"time"
)

// Interface defines a type that can implement its own DeepCopy method.
type Interface interface {
	DeepCopy() interface{}
}

// Copy creates a deep copy of the given value and returns it as an interface{}.
// The caller must assert the returned value to the correct type.
func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	// Handle primitive types directly.
	switch r := src.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128,
		string, bool:
		return r
	}

	// If the type implements DeepCopy, use it.
	if v, ok := src.(Interface); ok {
		return v.DeepCopy()
	}

	// Handle complex types using reflection.
	original := reflect.ValueOf(src)
	copyValue := reflect.New(original.Type()).Elem()
	copyRecursive(original, copyValue)
	return copyValue.Interface()
}

// copyRecursive performs deep copying of the given value recursively.
func copyRecursive(original, copyValue reflect.Value) {
	if !original.IsValid() || original.IsZero() {
		return
	}

	// If the type implements DeepCopy, use it.
	if original.CanInterface() {
		if copier, ok := original.Interface().(Interface); ok {
			copyValue.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	switch original.Kind() {
	case reflect.Ptr:
		if original.IsNil() {
			return
		}
		copyValue.Set(reflect.New(original.Elem().Type()))
		copyRecursive(original.Elem(), copyValue.Elem())

	case reflect.Interface:
		if original.IsNil() {
			return
		}
		copyElem := reflect.New(original.Elem().Type()).Elem()
		copyRecursive(original.Elem(), copyElem)
		copyValue.Set(copyElem)

	case reflect.Struct:
		if t, ok := original.Interface().(time.Time); ok {
			copyValue.Set(reflect.ValueOf(t))
			return
		}
		for i := 0; i < original.NumField(); i++ {
			if original.Type().Field(i).PkgPath == "" {
				copyRecursive(original.Field(i), copyValue.Field(i))
			}
		}

	case reflect.Slice:
		if original.IsNil() {
			return
		}
		copyValue.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			copyRecursive(original.Index(i), copyValue.Index(i))
		}

	case reflect.Map:
		if original.IsNil() {
			return
		}
		copyValue.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			copyKey := Copy(key.Interface())
			copyElem := reflect.New(original.MapIndex(key).Type()).Elem()
			copyRecursive(original.MapIndex(key), copyElem)
			copyValue.SetMapIndex(reflect.ValueOf(copyKey), copyElem)
		}

	default:
		copyValue.Set(original)
	}
}
