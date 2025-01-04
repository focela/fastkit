// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package binary provides APIs for handling binary/bytes data.
// It uses LittleEndian encoding by default.
package binary

// Encode serializes multiple values into binary format.
func Encode(values ...interface{}) []byte {
	return LeEncode(values...)
}

// EncodeByLength serializes values into binary with a specified length.
func EncodeByLength(length int, values ...interface{}) []byte {
	return LeEncodeByLength(length, values...)
}

// EncodeString encodes a string into binary format.
func EncodeString(s string) []byte {
	return LeEncodeString(s)
}

// EncodeBool encodes a boolean value into binary format.
func EncodeBool(b bool) []byte {
	return LeEncodeBool(b)
}

// EncodeInt encodes an integer into binary format.
func EncodeInt(i int) []byte {
	return LeEncodeInt(i)
}

// EncodeUint encodes an unsigned integer into binary format.
func EncodeUint(i uint) []byte {
	return LeEncodeUint(i)
}

// EncodeInt8 encodes an 8-bit integer into binary format.
func EncodeInt8(i int8) []byte {
	return LeEncodeInt8(i)
}

// EncodeUint8 encodes an 8-bit unsigned integer into binary format.
func EncodeUint8(i uint8) []byte {
	return LeEncodeUint8(i)
}

// EncodeInt16 encodes a 16-bit integer into binary format.
func EncodeInt16(i int16) []byte {
	return LeEncodeInt16(i)
}

// EncodeUint16 encodes a 16-bit unsigned integer into binary format.
func EncodeUint16(i uint16) []byte {
	return LeEncodeUint16(i)
}

// EncodeInt32 encodes a 32-bit integer into binary format.
func EncodeInt32(i int32) []byte {
	return LeEncodeInt32(i)
}

// EncodeUint32 encodes a 32-bit unsigned integer into binary format.
func EncodeUint32(i uint32) []byte {
	return LeEncodeUint32(i)
}

// EncodeInt64 encodes a 64-bit integer into binary format.
func EncodeInt64(i int64) []byte {
	return LeEncodeInt64(i)
}

// EncodeUint64 encodes a 64-bit unsigned integer into binary format.
func EncodeUint64(i uint64) []byte {
	return LeEncodeUint64(i)
}

// EncodeFloat32 encodes a 32-bit float into binary format.
func EncodeFloat32(f float32) []byte {
	return LeEncodeFloat32(f)
}

// EncodeFloat64 encodes a 64-bit float into binary format.
func EncodeFloat64(f float64) []byte {
	return LeEncodeFloat64(f)
}

// Decode deserializes binary data into provided values.
func Decode(b []byte, values ...interface{}) error {
	return LeDecode(b, values...)
}

// DecodeToString decodes binary data into a string.
func DecodeToString(b []byte) string {
	return LeDecodeToString(b)
}

// DecodeToBool decodes binary data into a boolean.
func DecodeToBool(b []byte) bool {
	return LeDecodeToBool(b)
}

// DecodeToInt decodes binary data into an integer.
func DecodeToInt(b []byte) int {
	return LeDecodeToInt(b)
}

// DecodeToUint decodes binary data into an unsigned integer.
func DecodeToUint(b []byte) uint {
	return LeDecodeToUint(b)
}

// DecodeToInt8 decodes binary data into an 8-bit integer.
func DecodeToInt8(b []byte) int8 {
	return LeDecodeToInt8(b)
}

// DecodeToUint8 decodes binary data into an 8-bit unsigned integer.
func DecodeToUint8(b []byte) uint8 {
	return LeDecodeToUint8(b)
}

// DecodeToInt16 decodes binary data into a 16-bit integer.
func DecodeToInt16(b []byte) int16 {
	return LeDecodeToInt16(b)
}

// DecodeToUint16 decodes binary data into a 16-bit unsigned integer.
func DecodeToUint16(b []byte) uint16 {
	return LeDecodeToUint16(b)
}

// DecodeToInt32 decodes binary data into a 32-bit integer.
func DecodeToInt32(b []byte) int32 {
	return LeDecodeToInt32(b)
}

// DecodeToUint32 decodes binary data into a 32-bit unsigned integer.
func DecodeToUint32(b []byte) uint32 {
	return LeDecodeToUint32(b)
}

// DecodeToInt64 decodes binary data into a 64-bit integer.
func DecodeToInt64(b []byte) int64 {
	return LeDecodeToInt64(b)
}

// DecodeToUint64 decodes binary data into a 64-bit unsigned integer.
func DecodeToUint64(b []byte) uint64 {
	return LeDecodeToUint64(b)
}

// DecodeToFloat32 decodes binary data into a 32-bit float.
func DecodeToFloat32(b []byte) float32 {
	return LeDecodeToFloat32(b)
}

// DecodeToFloat64 decodes binary data into a 64-bit float.
func DecodeToFloat64(b []byte) float64 {
	return LeDecodeToFloat64(b)
}
