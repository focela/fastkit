// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"bytes"
	"strings"
)

// DefaultTrimChars are the default characters removed by Trim* functions.
var DefaultTrimChars = string([]byte{
	// Whitespace and control characters
	'\t', '\v', '\n', '\r', '\f', ' ', 0x00, 0x85, 0xA0,
})

// Utility functions for character checks

// IsLetterUpper checks if the given byte `b` is an uppercase letter.
func IsLetterUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

// IsLetterLower checks if the given byte `b` is a lowercase letter.
func IsLetterLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

// IsLetter checks if the given byte `b` is a letter (uppercase or lowercase).
func IsLetter(b byte) bool {
	return IsLetterUpper(b) || IsLetterLower(b)
}

// String manipulation utilities

// IsNumeric checks if the given string `s` is numeric, including floats like "123.456".
func IsNumeric(s string) bool {
	dotCount := 0
	length := len(s)
	if length == 0 {
		return false
	}
	for i := 0; i < length; i++ {
		if (s[i] == '-' || s[i] == '+') && i == 0 {
			if length == 1 {
				return false
			}
			continue
		}
		if s[i] == '.' {
			dotCount++
			if dotCount > 1 || i == 0 || i == length-1 {
				return false
			}
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// UcFirst converts the first letter of the string `s` to uppercase.
func UcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	if IsLetterLower(s[0]) {
		return string(s[0]-32) + s[1:]
	}
	return s
}

// ReplaceByMap replaces substrings in `origin` based on the `replaces` map.
func ReplaceByMap(origin string, replaces map[string]string) string {
	for k, v := range replaces {
		origin = strings.ReplaceAll(origin, k, v)
	}
	return origin
}

// RemoveSymbols removes all symbols from the string `s`, leaving only numbers and letters.
func RemoveSymbols(s string) string {
	result := make([]rune, 0, len(s))
	for _, c := range s {
		if (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || c > 127 {
			result = append(result, c)
		}
	}
	return string(result)
}

// EqualFoldWithoutChars checks if `s1` and `s2` are equal, ignoring case and certain characters.
func EqualFoldWithoutChars(s1, s2 string) bool {
	return strings.EqualFold(RemoveSymbols(s1), RemoveSymbols(s2))
}

// SplitAndTrim splits the string `str` by `delimiter` and trims each element, removing empty elements.
func SplitAndTrim(str, delimiter string, characterMask ...string) []string {
	trimChars := DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	result := make([]string, 0)
	for _, part := range strings.Split(str, delimiter) {
		trimmed := strings.Trim(part, trimChars)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

// Trim removes specified characters (or defaults) from the start and end of the string `str`.
func Trim(str string, characterMask ...string) string {
	trimChars := DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.Trim(str, trimChars)
}

// Format utilities

// FormatCmdKey formats the string `s` as a command key (lowercase with dots).
func FormatCmdKey(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, "_", "."))
}

// FormatEnvKey formats the string `s` as an environment key (uppercase with underscores).
func FormatEnvKey(s string) string {
	return strings.ToUpper(strings.ReplaceAll(s, ".", "_"))
}

// StripSlashes removes escape slashes from the string `str`.
func StripSlashes(str string) string {
	var buf bytes.Buffer
	skipNext := false
	for i, char := range str {
		if skipNext {
			skipNext = false
			continue
		}
		if char == '\\' && i+1 < len(str) && str[i+1] == '\\' {
			skipNext = true
			continue
		}
		buf.WriteRune(char)
	}
	return buf.String()
}
