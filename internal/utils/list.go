// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"fmt"
)

// ListToMapByKey converts a list of maps into a map based on a specified key.
//
// Parameters:
// - list: A slice of maps with string keys and interface{} values.
// - key: The key used to group map items.
//
// Returns:
// - A map[string]interface{}, where keys are derived from the specified key and values are either single items or slices.
func ListToMapByKey(list []map[string]interface{}, key string) map[string]interface{} {
	result := make(map[string]interface{})
	groupedItems := make(map[string][]interface{})

	// Group items by key
	for _, item := range list {
		if k, ok := item[key]; ok {
			keyStr := fmt.Sprintf(`%v`, k)
			groupedItems[keyStr] = append(groupedItems[keyStr], item)
		}
	}

	// Flatten grouped items into the result map
	for k, v := range groupedItems {
		if len(v) > 1 {
			result[k] = v
		} else {
			result[k] = v[0]
		}
	}

	return result
}
