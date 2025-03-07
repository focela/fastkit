// Copyright (c) 2025 Focela Technologies.
// This software is provided "as is", without any warranty.
// Licensed under the MIT License – see LICENSE file for details.

// Package reflection provides enhanced reflection utilities for the Aegis framework.
// It simplifies common reflection operations like dereferencing pointers and type conversions
// that are frequently needed in framework internals.
package reflection

import (
	"reflect"
)

// OriginValueAndKindOutput holds the results of value and kind analysis.
// It contains both the input value/kind and the dereferenced original value/kind.
type OriginValueAndKindOutput struct {
	// InputValue is the reflect.Value of the input parameter.
	InputValue reflect.Value

	// InputKind is the reflect.Kind of the input parameter.
	InputKind reflect.Kind

	// OriginValue is the reflect.Value after dereferencing any pointers.
	OriginValue reflect.Value

	// OriginKind is the reflect.Kind after dereferencing any pointers.
	OriginKind reflect.Kind
}

// OriginValueAndKind retrieves and returns the original reflect value and kind.
// It traverses through pointer references to get to the underlying value.
//
// Parameters:
//   - value: can be any value or a reflect.Value
//
// Returns:
//   - OriginValueAndKindOutput containing both input and dereferenced information
func OriginValueAndKind(value interface{}) (out OriginValueAndKindOutput) {
	// Handle the input, which can be either a reflect.Value or any other value
	if v, ok := value.(reflect.Value); ok {
		out.InputValue = v
	} else {
		out.InputValue = reflect.ValueOf(value)
	}
	out.InputKind = out.InputValue.Kind()

	// Start with the input value
	out.OriginValue = out.InputValue
	out.OriginKind = out.InputKind

	// Dereference pointers until we reach a non-pointer
	for out.OriginKind == reflect.Ptr {
		out.OriginValue = out.OriginValue.Elem()
		out.OriginKind = out.OriginValue.Kind()
	}

	return
}

// OriginTypeAndKindOutput holds the results of type and kind analysis.
// It contains both the input type/kind and the dereferenced original type/kind.
type OriginTypeAndKindOutput struct {
	// InputType is the reflect.Type of the input parameter.
	InputType reflect.Type

	// InputKind is the reflect.Kind of the input parameter.
	InputKind reflect.Kind

	// OriginType is the reflect.Type after dereferencing any pointers.
	OriginType reflect.Type

	// OriginKind is the reflect.Kind after dereferencing any pointers.
	OriginKind reflect.Kind
}

// OriginTypeAndKind retrieves and returns the original reflect type and kind.
// It traverses through pointer types to get to the underlying type.
//
// Parameters:
//   - value: can be any value, a reflect.Type, or a reflect.Value
//
// Returns:
//   - OriginTypeAndKindOutput containing both input and dereferenced information
func OriginTypeAndKind(value interface{}) (out OriginTypeAndKindOutput) {
	// Handle nil value
	if value == nil {
		return
	}

	// Handle different possible input types
	if reflectType, ok := value.(reflect.Type); ok {
		// Input is already a reflect.Type
		out.InputType = reflectType
	} else {
		if reflectValue, ok := value.(reflect.Value); ok {
			// Input is a reflect.Value
			out.InputType = reflectValue.Type()
		} else {
			// Input is any other value
			out.InputType = reflect.TypeOf(value)
		}
	}
	out.InputKind = out.InputType.Kind()

	// Start with the input type
	out.OriginType = out.InputType
	out.OriginKind = out.InputKind

	// Dereference pointer types until we reach a non-pointer
	for out.OriginKind == reflect.Ptr {
		out.OriginType = out.OriginType.Elem()
		out.OriginKind = out.OriginType.Kind()
	}

	return
}

// ValueToInterface converts a reflect.Value to its corresponding interface{} value.
// It handles cases where direct Interface() calls might not be possible.
//
// Parameters:
//   - v: the reflect.Value to convert to interface{}
//
// Returns:
//   - value: the interface{} representation of the input value
//   - ok: whether the conversion was successful
func ValueToInterface(v reflect.Value) (value interface{}, ok bool) {
	// Handle the common case where we can directly call Interface()
	if v.IsValid() && v.CanInterface() {
		return v.Interface(), true
	}

	// Handle specific types that we can convert manually
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
		// For pointers and interfaces, dereference and try again
		return ValueToInterface(v.Elem())

	default:
		// For types we can't handle, return failure
		return nil, false
	}
}
