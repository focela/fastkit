// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package random provides high-performance random bytes/number/string generation functionality.
package random

import (
	"encoding/binary"
	"time"
)

const (
	letters    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52 letters
	symbols    = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"                   // 32 symbols
	digits     = "0123456789"                                           // 10 digits
	characters = letters + digits + symbols                             // 94 characters
)

// Intn returns a random integer in the range [0, max).
func Intn(max int) int {
	if max <= 0 {
		return max
	}
	n := int(binary.LittleEndian.Uint32(<-bufferChan)) % max
	if n < 0 {
		return -n
	}
	return n
}

// N returns a random integer between min and max, inclusive.
func N(min, max int) int {
	if min >= max {
		return min
	}
	if min >= 0 {
		return Intn(max-min+1) + min
	}
	return Intn(max+(0-min)+1) - (0 - min)
}

// D returns a random time.Duration between min and max.
func D(min, max time.Duration) time.Duration {
	multiple := int64(1)
	if min != 0 {
		for min%10 == 0 {
			multiple *= 10
			min /= 10
			max /= 10
		}
	}
	n := int64(N(int(min), int(max)))
	return time.Duration(n * multiple)
}

// Perm returns a random permutation of the integers [0, n).
func Perm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		j := Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}

// Meet returns true if a random value satisfies the probability num/total.
func Meet(num, total int) bool {
	return Intn(total) < num
}

// MeetProb returns true if a random value satisfies the given probability.
func MeetProb(prob float32) bool {
	return Intn(1e7) < int(prob*1e7)
}

// --------------------
// Byte Generators
// --------------------

// B generates and returns random bytes of length n.
func B(n int) []byte {
	if n <= 0 {
		return nil
	}
	b := make([]byte, n)
	for i := 0; i < n; i += 4 {
		copy(b[i:], <-bufferChan)
	}
	return b
}

// --------------------
// String Generators
// --------------------

// S generates a random string of length n, optionally including symbols.
func S(n int, useSymbols ...bool) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	numberBytes := B(n)
	for i := range b {
		if len(useSymbols) > 0 && useSymbols[0] {
			b[i] = characters[numberBytes[i]%94]
		} else {
			b[i] = characters[numberBytes[i]%62]
		}
	}
	return string(b)
}

// Str generates a random string of length n from a given string s.
func Str(s string, n int) string {
	if n <= 0 {
		return ""
	}
	b := make([]rune, n)
	runes := []rune(s)
	if len(runes) <= 255 {
		numberBytes := B(n)
		for i := range b {
			b[i] = runes[int(numberBytes[i])%len(runes)]
		}
	} else {
		for i := range b {
			b[i] = runes[Intn(len(runes))]
		}
	}
	return string(b)
}

// Digits generates a random string of digits of length n.
func Digits(n int) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	numberBytes := B(n)
	for i := range b {
		b[i] = digits[numberBytes[i]%10]
	}
	return string(b)
}

// Letters generates a random string of letters of length n.
func Letters(n int) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	numberBytes := B(n)
	for i := range b {
		b[i] = letters[numberBytes[i]%52]
	}
	return string(b)
}

// Symbols generates a random string of symbols of length n.
func Symbols(n int) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	numberBytes := B(n)
	for i := range b {
		b[i] = symbols[numberBytes[i]%32]
	}
	return string(b)
}
