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

// Package reflection provides utility functions for working with reflection in Go.
package reflection

import (
	"reflect"
)

// OriginValueAndKindOutput represents the result of retrieving the origin value and kind of an input.
type OriginValueAndKindOutput struct {
	InputValue  reflect.Value
	InputKind   reflect.Kind
	OriginValue reflect.Value
	OriginKind  reflect.Kind
}

// OriginTypeAndKindOutput represents the result of retrieving the origin type and kind of an input.
type OriginTypeAndKindOutput struct {
	InputType  reflect.Type
	InputKind  reflect.Kind
	OriginType reflect.Type
	OriginKind reflect.Kind
}

// OriginValueAndKind retrieves the original value and kind of the given input using reflection.
// It handles pointers and resolves them to their base value.
func OriginValueAndKind(value interface{}) (out OriginValueAndKindOutput) {
	if v, ok := value.(reflect.Value); ok {
		out.InputValue = v
	} else {
		out.InputValue = reflect.ValueOf(value)
	}
	out.InputKind = out.InputValue.Kind()
	out.OriginValue = out.InputValue
	out.OriginKind = out.InputKind

	// Resolve pointer chain to base value
	for out.OriginKind == reflect.Ptr && out.OriginValue.IsValid() && !out.OriginValue.IsZero() {
		out.OriginValue = out.OriginValue.Elem()
		out.OriginKind = out.OriginValue.Kind()
	}
	return
}

// OriginTypeAndKind retrieves the original type and kind of the given input using reflection.
// It handles pointers and resolves them to their base type.
func OriginTypeAndKind(value interface{}) (out OriginTypeAndKindOutput) {
	if value == nil {
		return
	}

	// Determine input type
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

	// Resolve pointer chain to base type
	for out.OriginKind == reflect.Ptr && out.OriginType != nil {
		out.OriginType = out.OriginType.Elem()
		out.OriginKind = out.OriginType.Kind()
	}
	return
}

// It supports primitive types, pointers, interfaces, and recursively resolves their values.
func ValueToInterface(v reflect.Value) (interface{}, bool) {
	if !v.IsValid() {
		return nil, false
	}
	if v.CanInterface() {
		return v.Interface(), true
	}

	// Handle specific kinds explicitly
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
