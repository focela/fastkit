// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package reflection provides reflection utilities for internal usage.
package reflection

import (
	"reflect"
)

// OriginValueAndKindOutput holds the input and resolved original value and kind.
type OriginValueAndKindOutput struct {
	InputValue  reflect.Value
	InputKind   reflect.Kind
	OriginValue reflect.Value
	OriginKind  reflect.Kind
}

// OriginValueAndKind retrieves the original reflect.Value and reflect.Kind from the given value.
func OriginValueAndKind(value interface{}) OriginValueAndKindOutput {
	var out OriginValueAndKindOutput

	if v, ok := value.(reflect.Value); ok {
		out.InputValue = v
	} else {
		out.InputValue = reflect.ValueOf(value)
	}
	out.InputKind = out.InputValue.Kind()
	out.OriginValue = out.InputValue
	out.OriginKind = out.InputKind

	for out.OriginKind == reflect.Ptr && out.OriginValue.IsValid() {
		out.OriginValue = out.OriginValue.Elem()
		out.OriginKind = out.OriginValue.Kind()
	}
	return out
}

// OriginTypeAndKindOutput holds the input and resolved original type and kind.
type OriginTypeAndKindOutput struct {
	InputType  reflect.Type
	InputKind  reflect.Kind
	OriginType reflect.Type
	OriginKind reflect.Kind
}

// OriginTypeAndKind retrieves the original reflect.Type and reflect.Kind from the given value.
func OriginTypeAndKind(value interface{}) OriginTypeAndKindOutput {
	var out OriginTypeAndKindOutput
	if value == nil {
		return out
	}

	switch v := value.(type) {
	case reflect.Type:
		out.InputType = v
	case reflect.Value:
		out.InputType = v.Type()
	default:
		out.InputType = reflect.TypeOf(value)
	}

	out.InputKind = out.InputType.Kind()
	out.OriginType = out.InputType
	out.OriginKind = out.InputKind

	for out.OriginKind == reflect.Ptr && out.OriginType != nil {
		out.OriginType = out.OriginType.Elem()
		out.OriginKind = out.OriginType.Kind()
	}
	return out
}

// ValueToInterface converts a reflect.Value to its interface type.
func ValueToInterface(v reflect.Value) (interface{}, bool) {
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
