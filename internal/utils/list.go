// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"fmt"
)

// ListToMapByKey converts a slice of maps into a map with keys derived from a specified key field.
//
// Parameters:
// - list: A slice of maps where each map represents an item.
// - key: The field in each map to use as the key in the resulting map.
//
// Returns:
// - A map[string]interface{} where the keys are derived from the specified field in the input maps.
// - If multiple items share the same key, their values are grouped into a slice.
//
// Note:
// - If there are duplicate keys, the resulting map stores slices of items.
// - If all keys are unique, the resulting map stores individual items.
func ListToMapByKey(list []map[string]interface{}, key string) map[string]interface{} {
	// Temporary map to group items by key.
	tempMap := make(map[string][]interface{})

	// Group items by their key field.
	for _, item := range list {
		// Check if the key exists in the current map item.
		if keyValue, ok := item[key]; ok {
			// Convert key value to string for map indexing.
			keyStr := fmt.Sprintf("%v", keyValue)
			// Append the item to the corresponding key group.
			tempMap[keyStr] = append(tempMap[keyStr], item)
		}
	}

	// Final map to store the result.
	resultMap := make(map[string]interface{})

	// Transform tempMap into the final resultMap.
	for k, v := range tempMap {
		if len(v) > 1 {
			// Store as a slice if there are multiple items for the same key.
			resultMap[k] = v
		} else {
			// Otherwise, store as a single item.
			resultMap[k] = v[0]
		}
	}

	return resultMap
}
