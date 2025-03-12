// Copyright (c) 2025 Focela Technologies.
// This software is provided "as is", without any warranty.
// Licensed under the MIT License – see LICENSE file for details.

// Package collection provides utility functions for working with collections in Go.
// It includes functions for inspecting and manipulating arrays, slices, maps and other collection types.
package collection

import (
	"fmt"
)

// ListToMapByKey converts a list of maps to a single map, indexed by a specified key's value.
//
// Parameters:
//   - list: A slice of map[string]interface{} containing the data to be converted
//   - key: The key to extract from each map to use as the index in the resulting map
//
// Returns:
//   - A map[string]interface{} where:
//   - Keys are string representations of the values found at `key` in the original maps
//   - Values are either:
//   - The original map if only one item has that key
//   - A slice of maps if multiple items share the same key
//
// Example:
//
//	Input: list=[{"id":"a","val":1}, {"id":"b","val":2}, {"id":"a","val":3}], key="id"
//	Output: {"a":[{"id":"a","val":1}, {"id":"a","val":3}], "b":{"id":"b","val":2}}
func ListToMapByKey(list []map[string]interface{}, key string) map[string]interface{} {
	// Initialize result maps
	resultMap := make(map[string]interface{})

	// This temporary map stores lists of items grouped by key
	groupedItems := make(map[string][]interface{})

	// Track if any key appears multiple times
	hasMultipleValues := false

	// Group items by key
	for _, item := range list {
		if keyValue, exists := item[key]; exists {
			// Convert key value to string for map indexing
			keyString := fmt.Sprintf("%v", keyValue)

			// Add this item to the group for this key
			groupedItems[keyString] = append(groupedItems[keyString], item)

			// Check if this key now has multiple values
			if len(groupedItems[keyString]) > 1 {
				hasMultipleValues = true
			}
		}
	}

	// Build the final result map
	for keyString, items := range groupedItems {
		if hasMultipleValues {
			// If any key has multiple values, preserve the list structure for all keys
			resultMap[keyString] = items
		} else {
			// Otherwise, simplify by storing just the single item
			resultMap[keyString] = items[0]
		}
	}

	return resultMap
}
