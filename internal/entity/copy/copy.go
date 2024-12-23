/*
 * FOCELA TECHNOLOGIES INTERNAL USE ONLY LICENSE AGREEMENT
 *
 * Copyright (c) 2024 Focela Technologies. All rights reserved.
 *
 * Permission is hereby granted to employees or authorized personnel of Focela
 * Technologies (the "Company") to use this software solely for internal business
 * purposes within the Company.
 *
 * For inquiries or permissions, please contact: legal@focela.com
 */

// Package deepcopy provides functionalities for deep copying objects using reflection.
// This package is maintained from: https://github.com/mohae/deepcopy
package deepcopy

import (
	"reflect"
	"time"
)

// Interface defines a custom interface for types that support their own DeepCopy implementation.
type Interface interface {
	DeepCopy() interface{}
}

// Copy creates a deep copy of the given object using reflection.
// The returned value will need to be asserted to the correct type.
func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	// Directly return simple types
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
		// Custom DeepCopy implementation
		if v, ok := src.(Interface); ok {
			return v.DeepCopy()
		}

		// Use reflection for deep copy
		original := reflect.ValueOf(src)
		dst := reflect.New(original.Type()).Elem()
		copyRecursive(original, dst)
		return dst.Interface()
	}
}

// copyRecursive performs the actual recursive copying using reflection.
// It supports types like Ptr, Interface, Struct, Slice, and Map.
func copyRecursive(original, cpy reflect.Value) {
	if !original.IsValid() || original.IsZero() {
		return
	}

	// Handle DeepCopy interface implementation
	if original.CanInterface() {
		if copier, ok := original.Interface().(Interface); ok {
			cpy.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	// Handle based on kind
	switch original.Kind() {
	case reflect.Ptr:
		if original.IsNil() {
			return
		}
		cpy.Set(reflect.New(original.Type().Elem()))
		copyRecursive(original.Elem(), cpy.Elem())

	case reflect.Interface:
		if original.IsNil() {
			return
		}
		originalValue := original.Elem()
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue)
		cpy.Set(copyValue)

	case reflect.Struct:
		if t, ok := original.Interface().(time.Time); ok {
			cpy.Set(reflect.ValueOf(t))
			return
		}
		for i := 0; i < original.NumField(); i++ {
			field := original.Type().Field(i)
			if !field.IsExported() {
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
