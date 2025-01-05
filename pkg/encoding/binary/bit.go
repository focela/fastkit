// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package binary provides APIs for handling binary/bytes data with BigEndian encoding.
package binary

// Bit represents a binary bit (0 | 1).
type Bit int8

// EncodeBits encodes an integer `i` into bits with length `l` and appends it to the given `bits` slice.
func EncodeBits(bits []Bit, i int, l int) []Bit {
	return EncodeBitsWithUint(bits, uint(i), l)
}

// EncodeBitsWithUint encodes an unsigned integer `ui` into bits with length `l` and appends it to the given `bits` slice.
func EncodeBitsWithUint(bits []Bit, ui uint, l int) []Bit {
	encodedBits := make([]Bit, l)
	for i := l - 1; i >= 0; i-- {
		encodedBits[i] = Bit(ui & 1)
		ui >>= 1
	}
	if bits != nil {
		return append(bits, encodedBits...)
	}
	return encodedBits
}

// EncodeBitsToBytes converts a slice of bits into a slice of bytes.
// If the number of bits is not a multiple of 8, it pads with zeros.
func EncodeBitsToBytes(bits []Bit) []byte {
	// Padding with zeros to make bits length a multiple of 8.
	padding := (8 - len(bits)%8) % 8
	for i := 0; i < padding; i++ {
		bits = append(bits, 0)
	}

	// Encode bits to bytes.
	bytes := make([]byte, len(bits)/8)
	for i := 0; i < len(bits); i += 8 {
		bytes[i/8] = byte(DecodeBitsToUint(bits[i : i+8]))
	}
	return bytes
}

// DecodeBits decodes a slice of bits into an integer.
func DecodeBits(bits []Bit) int {
	result := 0
	for _, bit := range bits {
		result = (result << 1) | int(bit)
	}
	return result
}

// DecodeBitsToUint decodes a slice of bits into an unsigned integer.
func DecodeBitsToUint(bits []Bit) uint {
	result := uint(0)
	for _, bit := range bits {
		result = (result << 1) | uint(bit)
	}
	return result
}

// DecodeBytesToBits decodes a slice of bytes into a slice of bits.
func DecodeBytesToBits(bytes []byte) []Bit {
	var bits []Bit
	for _, b := range bytes {
		bits = EncodeBitsWithUint(bits, uint(b), 8)
	}
	return bits
}
