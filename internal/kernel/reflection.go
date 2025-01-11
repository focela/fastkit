// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package kernel provides core utilities and foundational components for the Aegis framework.
package kernel

import (
	"reflect"
)

// OriginValueAndKindOutput represents the original value and kind of a reflect.Value.
type OriginValueAndKindOutput struct {
	InputValue  reflect.Value // Initial reflect.Value.
	InputKind   reflect.Kind  // Kind of the initial value.
	OriginValue reflect.Value // Resolved original reflect.Value.
	OriginKind  reflect.Kind  // Resolved original kind.
}

// OriginTypeAndKindOutput represents the original type and kind of a reflect.Type.
type OriginTypeAndKindOutput struct {
	InputType  reflect.Type // Initial reflect.Type.
	InputKind  reflect.Kind // Kind of the initial type.
	OriginType reflect.Type // Resolved original reflect.Type.
	OriginKind reflect.Kind // Resolved original kind.
}

// OriginValueAndKind retrieves the original reflect.Value and reflect.Kind of the input.
//
// Parameters:
// - value: The input value which can be of any type.
//
// Returns:
// - A struct containing the input and resolved original value and kind.
func OriginValueAndKind(value interface{}) (out OriginValueAndKindOutput) {
	if v, ok := value.(reflect.Value); ok {
		out.InputValue = v
	} else {
		out.InputValue = reflect.ValueOf(value)
	}

	out.InputKind = out.InputValue.Kind()
	out.OriginValue = out.InputValue
	out.OriginKind = out.InputKind

	for out.OriginKind == reflect.Ptr {
		if out.OriginValue.IsNil() {
			break
		}
		out.OriginValue = out.OriginValue.Elem()
		out.OriginKind = out.OriginValue.Kind()
	}
	return
}

// OriginTypeAndKind retrieves the original reflect.Type and reflect.Kind of the input.
//
// Parameters:
// - value: The input value which can be of any type.
//
// Returns:
// - A struct containing the input and resolved original type and kind.
func OriginTypeAndKind(value interface{}) (out OriginTypeAndKindOutput) {
	if value == nil {
		return
	}

	if inputType, ok := value.(reflect.Type); ok {
		out.InputType = inputType
	} else if inputValue, ok := value.(reflect.Value); ok {
		out.InputType = inputValue.Type()
	} else {
		out.InputType = reflect.TypeOf(value)
	}

	out.InputKind = out.InputType.Kind()
	out.OriginType = out.InputType
	out.OriginKind = out.InputKind

	for out.OriginKind == reflect.Ptr {
		if out.OriginType == nil {
			break
		}
		out.OriginType = out.OriginType.Elem()
		out.OriginKind = out.OriginType.Kind()
	}
	return
}

// ValueToInterface converts a reflect.Value into its corresponding interface{} type.
//
// Parameters:
// - v: The reflect.Value to be converted.
//
// Returns:
// - value: The corresponding interface{} type.
// - ok: Boolean indicating whether the conversion was successful.
func ValueToInterface(v reflect.Value) (value interface{}, ok bool) {
	if v.IsValid() && v.CanInterface() {
		return v.Interface(), true
	}

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
		return ValueToInterface(v.Elem())
	default:
		return nil, false
	}
}
