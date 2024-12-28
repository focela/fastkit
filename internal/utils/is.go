// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"reflect"

	"github.com/focela/loom/internal/empty"
)

// IsNil checks if a value is nil, including interface{}.
func IsNil(value interface{}) bool {
	return empty.IsNil(value)
}

// IsEmpty checks if a value is empty.
func IsEmpty(value interface{}) bool {
	return empty.IsEmpty(value)
}

// IsInt checks if a value is of integer type.
func IsInt(value interface{}) bool {
	switch value.(type) {
	case int, *int, int8, *int8, int16, *int16, int32, *int32, int64, *int64:
		return true
	}
	return false
}

// IsUint checks if a value is of unsigned integer type.
func IsUint(value interface{}) bool {
	switch value.(type) {
	case uint, *uint, uint8, *uint8, uint16, *uint16, uint32, *uint32, uint64, *uint64:
		return true
	}
	return false
}

// IsFloat checks if a value is of float type.
func IsFloat(value interface{}) bool {
	switch value.(type) {
	case float32, *float32, float64, *float64:
		return true
	}
	return false
}

// getKind dereferences a pointer to get the base kind of a value.
func getKind(value interface{}) reflect.Kind {
	reflectValue := reflect.ValueOf(value)
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue.Kind()
}

// IsSlice checks if a value is of slice or array type.
func IsSlice(value interface{}) bool {
	switch getKind(value) {
	case reflect.Slice, reflect.Array:
		return true
	}
	return false
}

// IsMap checks if a value is of map type.
func IsMap(value interface{}) bool {
	return getKind(value) == reflect.Map
}

// IsStruct checks if a value is of struct type.
func IsStruct(value interface{}) bool {
	reflectType := reflect.TypeOf(value)
	if reflectType == nil {
		return false
	}
	for reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	return reflectType.Kind() == reflect.Struct
}
