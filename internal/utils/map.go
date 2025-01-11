// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

// MapPossibleItemByKey searches for a key-value pair in the map `data` that matches the given `key`.
//
// Parameters:
// - data: The map to search in.
// - key: The key to look for.
//
// Returns:
// - foundKey: The actual key found in the map that matches `key`.
// - foundValue: The value corresponding to the `foundKey`.
//
// Notes:
// - This function performs a case-insensitive and symbol-ignoring comparison.
// - If no match is found, it returns an empty string and nil.
// - May have low performance due to iteration over all map keys.
func MapPossibleItemByKey(data map[string]interface{}, key string) (foundKey string, foundValue interface{}) {
	if len(data) == 0 {
		return "", nil
	}

	// Direct lookup for exact match
	if value, ok := data[key]; ok {
		return key, value
	}

	// Iterative lookup for case-insensitive and symbol-ignored match
	for k, v := range data {
		if EqualFoldWithoutChars(k, key) {
			return k, v
		}
	}

	return "", nil
}

// MapContainsPossibleKey checks if the map `data` contains a key that matches the given `key`.
//
// Parameters:
// - data: The map to search in.
// - key: The key to look for.
//
// Returns:
// - true if the map contains a matching key, false otherwise.
//
// Notes:
// - This function performs a case-insensitive and symbol-ignoring comparison.
// - May have low performance due to reliance on `MapPossibleItemByKey`.
func MapContainsPossibleKey(data map[string]interface{}, key string) bool {
	foundKey, _ := MapPossibleItemByKey(data, key)
	return foundKey != ""
}
