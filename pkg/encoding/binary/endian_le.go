// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package binary provides APIs for handling binary/bytes data.
package binary

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/focela/loom/internal/core"
	"github.com/focela/loom/pkg/errors"
)

// LeEncode encodes multiple values into little-endian binary format.
func LeEncode(values ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for _, value := range values {
		if value == nil {
			return buf.Bytes()
		}
		switch v := value.(type) {
		case int:
			buf.Write(LeEncodeInt(v))
		case int8:
			buf.Write(LeEncodeInt8(v))
		case int16:
			buf.Write(LeEncodeInt16(v))
		case int32:
			buf.Write(LeEncodeInt32(v))
		case int64:
			buf.Write(LeEncodeInt64(v))
		case uint:
			buf.Write(LeEncodeUint(v))
		case uint8:
			buf.Write(LeEncodeUint8(v))
		case uint16:
			buf.Write(LeEncodeUint16(v))
		case uint32:
			buf.Write(LeEncodeUint32(v))
		case uint64:
			buf.Write(LeEncodeUint64(v))
		case bool:
			buf.Write(LeEncodeBool(v))
		case string:
			buf.Write(LeEncodeString(v))
		case []byte:
			buf.Write(v)
		case float32:
			buf.Write(LeEncodeFloat32(v))
		case float64:
			buf.Write(LeEncodeFloat64(v))
		default:
			if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
				core.Errorf(context.TODO(), `%+v`, err)
				buf.Write(LeEncodeString(fmt.Sprintf("%v", v)))
			}
		}
	}
	return buf.Bytes()
}

// LeEncodeByLength encodes values into a fixed-length byte slice.
func LeEncodeByLength(length int, values ...interface{}) []byte {
	b := LeEncode(values...)
	if len(b) < length {
		b = append(b, make([]byte, length-len(b))...)
	} else if len(b) > length {
		b = b[:length]
	}
	return b
}

// LeEncodeString converts a string into a byte slice.
func LeEncodeString(s string) []byte { return []byte(s) }

// LeEncodeBool encodes a boolean into one byte.
func LeEncodeBool(b bool) []byte {
	if b {
		return []byte{1}
	}
	return []byte{0}
}

// LeEncodeInt encodes an integer into bytes based on size.
func LeEncodeInt(i int) []byte {
	switch {
	case i <= math.MaxInt8:
		return EncodeInt8(int8(i))
	case i <= math.MaxInt16:
		return EncodeInt16(int16(i))
	case i <= math.MaxInt32:
		return EncodeInt32(int32(i))
	default:
		return EncodeInt64(int64(i))
	}
}

// LeEncodeUint encodes an unsigned integer into bytes based on size.
func LeEncodeUint(i uint) []byte {
	switch {
	case i <= math.MaxUint8:
		return EncodeUint8(uint8(i))
	case i <= math.MaxUint16:
		return EncodeUint16(uint16(i))
	case i <= math.MaxUint32:
		return EncodeUint32(uint32(i))
	default:
		return EncodeUint64(uint64(i))
	}
}

// LeEncodeInt8 encodes an int8 into one byte.
func LeEncodeInt8(i int8) []byte { return []byte{byte(i)} }

// LeEncodeUint8 encodes a uint8 into one byte.
func LeEncodeUint8(i uint8) []byte { return []byte{i} }

// LeEncodeInt16 encodes an int16 into two bytes.
func LeEncodeInt16(i int16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(i))
	return b
}

// LeEncodeUint16 encodes a uint16 into two bytes.
func LeEncodeUint16(i uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, i)
	return b
}

// LeEncodeInt32 encodes an int32 into four bytes.
func LeEncodeInt32(i int32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(i))
	return b
}

// LeEncodeUint32 encodes a uint32 into four bytes.
func LeEncodeUint32(i uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	return b
}

// LeEncodeInt64 encodes an int64 into eight bytes.
func LeEncodeInt64(i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

// LeEncodeUint64 encodes a uint64 into eight bytes.
func LeEncodeUint64(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}

// LeEncodeFloat32 encodes a float32 into four bytes.
func LeEncodeFloat32(f float32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, math.Float32bits(f))
	return b
}

// LeEncodeFloat64 encodes a float64 into eight bytes.
func LeEncodeFloat64(f float64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, math.Float64bits(f))
	return b
}

// LeDecode decodes binary data into provided variables.
func LeDecode(b []byte, values ...interface{}) error {
	buf := bytes.NewBuffer(b)
	for _, value := range values {
		if err := binary.Read(buf, binary.LittleEndian, value); err != nil {
			return errors.Wrap(err, `binary.Read failed`)
		}
	}
	return nil
}

// LeDecodeToString converts a byte slice to a string.
func LeDecodeToString(b []byte) string { return string(b) }

// LeDecodeToInt converts a byte slice to an int.
func LeDecodeToInt(b []byte) int { return int(LeDecodeToInt64(b)) }

// LeDecodeToUint converts a byte slice to a uint.
func LeDecodeToUint(b []byte) uint { return uint(LeDecodeToUint64(b)) }

// LeDecodeToBool converts a byte slice to a bool.
func LeDecodeToBool(b []byte) bool { return len(b) > 0 && b[0] != 0 }

// LeDecodeToInt8 converts a byte slice to an int8.
func LeDecodeToInt8(b []byte) int8 { return int8(b[0]) }

// LeDecodeToUint8 converts a byte slice to a uint8.
func LeDecodeToUint8(b []byte) uint8 { return b[0] }

// LeDecodeToInt16 converts a byte slice to an int16.
func LeDecodeToInt16(b []byte) int16 { return int16(binary.LittleEndian.Uint16(LeFillUpSize(b, 2))) }

// LeDecodeToUint16 converts a byte slice to a uint16.
func LeDecodeToUint16(b []byte) uint16 { return binary.LittleEndian.Uint16(LeFillUpSize(b, 2)) }

// LeDecodeToInt32 converts a byte slice to an int32.
func LeDecodeToInt32(b []byte) int32 { return int32(binary.LittleEndian.Uint32(LeFillUpSize(b, 4))) }

// LeDecodeToUint32 converts a byte slice to a uint32.
func LeDecodeToUint32(b []byte) uint32 { return binary.LittleEndian.Uint32(LeFillUpSize(b, 4)) }

// LeDecodeToInt64 converts a byte slice to an int64.
func LeDecodeToInt64(b []byte) int64 { return int64(binary.LittleEndian.Uint64(LeFillUpSize(b, 8))) }

// LeDecodeToUint64 converts a byte slice to a uint64.
func LeDecodeToUint64(b []byte) uint64 { return binary.LittleEndian.Uint64(LeFillUpSize(b, 8)) }

// LeDecodeToFloat32 converts a byte slice to a float32.
func LeDecodeToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(LeFillUpSize(b, 4)))
}

// LeDecodeToFloat64 converts a byte slice to a float64.
func LeDecodeToFloat64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(LeFillUpSize(b, 8)))
}

// LeFillUpSize ensures a byte slice has the required length by padding it with zeros if necessary.
func LeFillUpSize(b []byte, l int) []byte {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c, b)
	return c
}
