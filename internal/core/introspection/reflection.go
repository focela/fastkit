// Copyright (c) 2025 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package introspection provides utilities for working with Go's reflection system.
// It offers helpers to safely extract and manipulate type information and values
// at runtime with less boilerplate code.
package introspection

import (
	"reflect"
)

// OriginValueAndKindOutput contains both the input and dereferenced (origin) value information.
// It helps track both the original input and the final value after following pointer references.
type OriginValueAndKindOutput struct {
	// InputValue is the reflect.Value of the original input.
	InputValue reflect.Value

	// InputKind is the reflect.Kind of the original input.
	InputKind reflect.Kind

	// OriginValue is the reflect.Value after following all pointers (if any).
	OriginValue reflect.Value

	// OriginKind is the reflect.Kind after following all pointers (if any).
	OriginKind reflect.Kind
}

// OriginValueAndKind retrieves and returns the original reflect value and kind.
// It dereferences pointers until a non-pointer value is found, tracking both
// the input and the dereferenced "origin" value.
//
// The function accepts either a reflect.Value or any other value, which will
// be converted to a reflect.Value internally.
func OriginValueAndKind(value interface{}) (out OriginValueAndKindOutput) {
	// Handle the input, which might be a reflect.Value or any other type
	if v, ok := value.(reflect.Value); ok {
		out.InputValue = v
	} else {
		out.InputValue = reflect.ValueOf(value)
	}
	out.InputKind = out.InputValue.Kind()

	// Start with the input value
	out.OriginValue = out.InputValue
	out.OriginKind = out.InputKind

	// Follow pointers to their target values
	for out.OriginKind == reflect.Ptr {
		out.OriginValue = out.OriginValue.Elem()
		out.OriginKind = out.OriginValue.Kind()
	}

	return
}

// OriginTypeAndKindOutput contains both the input and dereferenced (origin) type information.
// It helps track both the original input type and the final type after following pointer references.
type OriginTypeAndKindOutput struct {
	// InputType is the reflect.Type of the original input.
	InputType reflect.Type

	// InputKind is the reflect.Kind of the original input.
	InputKind reflect.Kind

	// OriginType is the reflect.Type after following all pointers (if any).
	OriginType reflect.Type

	// OriginKind is the reflect.Kind after following all pointers (if any).
	OriginKind reflect.Kind
}

// OriginTypeAndKind retrieves and returns the original reflect type and kind.
// It dereferences pointer types until a non-pointer type is found, tracking both
// the input and the dereferenced "origin" type.
//
// The function accepts:
// - reflect.Type: used directly
// - reflect.Value: the Type() is extracted
// - any other value: reflect.TypeOf() is used
// - nil: returns an empty output structure
func OriginTypeAndKind(value interface{}) (out OriginTypeAndKindOutput) {
	// Handle nil input
	if value == nil {
		return
	}

	// Handle the input, which could be one of several types
	if reflectType, ok := value.(reflect.Type); ok {
		// Input is already a reflect.Type
		out.InputType = reflectType
	} else if reflectValue, ok := value.(reflect.Value); ok {
		// Input is a reflect.Value
		out.InputType = reflectValue.Type()
	} else {
		// Input is any other value
		out.InputType = reflect.TypeOf(value)
	}

	out.InputKind = out.InputType.Kind()

	// Start with the input type
	out.OriginType = out.InputType
	out.OriginKind = out.InputKind

	// Follow pointer types to their target types
	for out.OriginKind == reflect.Ptr {
		out.OriginType = out.OriginType.Elem()
		out.OriginKind = out.OriginType.Kind()
	}

	return
}

// ValueToInterface converts a reflect.Value to its native Go interface{} representation.
// This is useful when you need to extract the actual value from a reflect.Value object.
//
// For values that cannot be directly converted with Interface(), such as unexported fields,
// the function attempts to extract the underlying primitive value based on its Kind.
//
// Returns:
//   - value: The extracted interface{} value
//   - ok: Whether the extraction was successful
func ValueToInterface(v reflect.Value) (value interface{}, ok bool) {
	// First try the direct approach if the value is accessible
	if v.IsValid() && v.CanInterface() {
		return v.Interface(), true
	}

	// Handle different kinds of values that might not be directly accessible
	switch v.Kind() {
	case reflect.Bool:
		return v.Bool(), true

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), true

	case reflect.Float32, reflect.Float64:
		return v.Float(), true

	case reflect.Complex64, reflect.Complex128:
		return v.Complex(), true

	case reflect.String:
		return v.String(), true

	case reflect.Ptr, reflect.Interface:
		// Recursively handle pointers and interfaces
		return ValueToInterface(v.Elem())

	default:
		// Cannot convert other kinds (map, slice, struct, etc.) if not accessible
		return nil, false
	}
}
