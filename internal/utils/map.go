// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

// MapPossibleItemByKey searches for a key-value pair in a map, ignoring case and symbols.
//
// It performs a case-insensitive and symbol-agnostic comparison to find the closest matching key.
// Note: This function might have low performance with large datasets.
//
// Parameters:
// - data: A map with string keys and interface{} values.
// - key: The key to search for.
//
// Returns:
// - foundKey: The matched key if found.
// - foundValue: The associated value of the matched key, or nil if not found.
func MapPossibleItemByKey(data map[string]interface{}, key string) (foundKey string, foundValue interface{}) {
	if len(data) == 0 {
		return "", nil
	}

	// Direct match
	if v, ok := data[key]; ok {
		return key, v
	}

	// Case and symbol insensitive search
	for k, v := range data {
		if EqualFoldWithoutChars(k, key) {
			return k, v
		}
	}
	return "", nil
}

// MapContainsPossibleKey checks whether a key exists in a map, ignoring case and symbols.
//
// It uses MapPossibleItemByKey to perform the key lookup.
// Note: This function might have low performance with large datasets.
//
// Parameters:
// - data: A map with string keys and interface{} values.
// - key: The key to search for.
//
// Returns:
// - true if the key exists, otherwise false.
func MapContainsPossibleKey(data map[string]interface{}, key string) bool {
	foundKey, _ := MapPossibleItemByKey(data, key)
	return foundKey != ""
}
