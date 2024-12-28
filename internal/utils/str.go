// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"bytes"
	"strings"
)

// DefaultTrimChars defines default characters for trimming operations.
var DefaultTrimChars = string([]byte{
	'\t', '\v', '\n', '\r', '\f', ' ', 0x00, 0x85, 0xA0,
})

// IsLetterUpper checks if a byte is an uppercase letter.
func IsLetterUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

// IsLetterLower checks if a byte is a lowercase letter.
func IsLetterLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

// IsLetter checks if a byte is a letter.
func IsLetter(b byte) bool {
	return IsLetterUpper(b) || IsLetterLower(b)
}

// IsNumeric checks if a string is numeric, including floats.
func IsNumeric(s string) bool {
	var dotCount int
	if len(s) == 0 {
		return false
	}
	for i, ch := range s {
		switch {
		case (ch == '-' || ch == '+') && i == 0:
			if len(s) == 1 {
				return false
			}
		case ch == '.':
			dotCount++
			if i == 0 || i == len(s)-1 {
				return false
			}
		case ch < '0' || ch > '9':
			return false
		}
	}
	return dotCount <= 1
}

// UcFirst capitalizes the first letter of a string.
func UcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	if IsLetterLower(s[0]) {
		return string(s[0]-32) + s[1:]
	}
	return s
}

// ReplaceByMap replaces substrings in a string based on a map.
func ReplaceByMap(origin string, replaces map[string]string) string {
	for k, v := range replaces {
		origin = strings.ReplaceAll(origin, k, v)
	}
	return origin
}

// RemoveSymbols removes all non-alphanumeric symbols from a string.
func RemoveSymbols(s string) string {
	var b []rune
	for _, c := range s {
		if c > 127 || ('0' <= c && c <= '9') || ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') {
			b = append(b, c)
		}
	}
	return string(b)
}

// EqualFoldWithoutChars compares two strings case-insensitively, ignoring symbols.
func EqualFoldWithoutChars(s1, s2 string) bool {
	return strings.EqualFold(RemoveSymbols(s1), RemoveSymbols(s2))
}

// SplitAndTrim splits and trims a string by a delimiter.
func SplitAndTrim(str, delimiter string, characterMask ...string) []string {
	var array []string
	for _, v := range strings.Split(str, delimiter) {
		v = Trim(v, characterMask...)
		if v != "" {
			array = append(array, v)
		}
	}
	return array
}

// Trim removes specified characters from both ends of a string.
func Trim(str string, characterMask ...string) string {
	trimChars := DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.Trim(str, trimChars)
}

// FormatCmdKey converts a string to lowercase with '.' as a separator.
func FormatCmdKey(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, "_", "."))
}

// FormatEnvKey converts a string to uppercase with '_' as a separator.
func FormatEnvKey(s string) string {
	return strings.ToUpper(strings.ReplaceAll(s, ".", "_"))
}

// StripSlashes removes escape slashes from a string.
func StripSlashes(str string) string {
	var buf bytes.Buffer
	l, skip := len(str), false
	for i, char := range str {
		if skip {
			skip = false
		} else if char == '\\' && i+1 < l && str[i+1] == '\\' {
			skip = true
			continue
		}
		buf.WriteRune(char)
	}
	return buf.String()
}
