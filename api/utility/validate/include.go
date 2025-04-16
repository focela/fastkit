// Copyright 2025 Focela – MIT Licensed. See LICENSE.
// Part of the Focela Open Initiative.
// Built to empower developers and partners.

// Package validate provides utility functions for validating various data structures.
package validate

// InSlice checks if a given element exists in a slice.
// It supports any comparable type through Go generics.
//
// Parameters:
//   - slice: The slice to search in
//   - element: The element to search for
//
// Returns:
//   - bool: true if the element is found, false otherwise
func InSlice[T comparable](slice []T, element T) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}
