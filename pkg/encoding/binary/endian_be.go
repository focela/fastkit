// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package binary provides APIs for handling binary/bytes data with BigEndian encoding.
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

// BeEncode encodes multiple values into big-endian binary format.
func BeEncode(values ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for _, value := range values {
		if value == nil {
			return buf.Bytes()
		}
		switch v := value.(type) {
		case int:
			buf.Write(BeEncodeInt(v))
		case int8:
			buf.Write(BeEncodeInt8(v))
		case int16:
			buf.Write(BeEncodeInt16(v))
		case int32:
			buf.Write(BeEncodeInt32(v))
		case int64:
			buf.Write(BeEncodeInt64(v))
		case uint:
			buf.Write(BeEncodeUint(v))
		case uint8:
			buf.Write(BeEncodeUint8(v))
		case uint16:
			buf.Write(BeEncodeUint16(v))
		case uint32:
			buf.Write(BeEncodeUint32(v))
		case uint64:
			buf.Write(BeEncodeUint64(v))
		case bool:
			buf.Write(BeEncodeBool(v))
		case string:
			buf.Write(BeEncodeString(v))
		case []byte:
			buf.Write(v)
		case float32:
			buf.Write(BeEncodeFloat32(v))
		case float64:
			buf.Write(BeEncodeFloat64(v))
		default:
			if err := binary.Write(buf, binary.BigEndian, v); err != nil {
				core.Errorf(context.TODO(), `%+v`, err)
				buf.Write(BeEncodeString(fmt.Sprintf("%v", v)))
			}
		}
	}
	return buf.Bytes()
}

// BeEncodeByLength encodes values into a fixed-length byte slice.
func BeEncodeByLength(length int, values ...interface{}) []byte {
	b := BeEncode(values...)
	if len(b) < length {
		b = append(b, make([]byte, length-len(b))...)
	} else if len(b) > length {
		b = b[:length]
	}
	return b
}

// BeEncodeString converts a string to a byte slice.
func BeEncodeString(s string) []byte { return []byte(s) }

// BeEncodeBool encodes a boolean value into one byte.
func BeEncodeBool(b bool) []byte {
	if b {
		return []byte{1}
	}
	return []byte{0}
}

// BeEncodeInt encodes an integer into big-endian binary format.
func BeEncodeInt(i int) []byte {
	switch {
	case i <= math.MaxInt8:
		return BeEncodeInt8(int8(i))
	case i <= math.MaxInt16:
		return BeEncodeInt16(int16(i))
	case i <= math.MaxInt32:
		return BeEncodeInt32(int32(i))
	default:
		return BeEncodeInt64(int64(i))
	}
}

// BeEncodeUint encodes an unsigned integer into big-endian binary format.
func BeEncodeUint(i uint) []byte {
	switch {
	case i <= math.MaxUint8:
		return BeEncodeUint8(uint8(i))
	case i <= math.MaxUint16:
		return BeEncodeUint16(uint16(i))
	case i <= math.MaxUint32:
		return BeEncodeUint32(uint32(i))
	default:
		return BeEncodeUint64(uint64(i))
	}
}

// BeEncodeInt8 encodes an int8 into one byte.
func BeEncodeInt8(i int8) []byte { return []byte{byte(i)} }

// BeEncodeUint8 encodes a uint8 into one byte.
func BeEncodeUint8(i uint8) []byte { return []byte{i} }

// BeEncodeInt16 encodes an int16 into two bytes.
func BeEncodeInt16(i int16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(i))
	return b
}

// BeEncodeUint16 encodes a uint16 into two bytes.
func BeEncodeUint16(i uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return b
}

// BeEncodeInt32 encodes an int32 into four bytes.
func BeEncodeInt32(i int32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}

// BeEncodeUint32 encodes a uint32 into four bytes.
func BeEncodeUint32(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

// BeEncodeInt64 encodes an int64 into eight bytes.
func BeEncodeInt64(i int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

// BeEncodeUint64 encodes a uint64 into eight bytes.
func BeEncodeUint64(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

// BeEncodeFloat32 encodes a float32 into four bytes.
func BeEncodeFloat32(f float32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, math.Float32bits(f))
	return b
}

// BeEncodeFloat64 encodes a float64 into eight bytes.
func BeEncodeFloat64(f float64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, math.Float64bits(f))
	return b
}

// BeDecode decodes binary data into provided variables.
func BeDecode(b []byte, values ...interface{}) error {
	buf := bytes.NewBuffer(b)
	for _, value := range values {
		if err := binary.Read(buf, binary.BigEndian, value); err != nil {
			return errors.Wrap(err, `binary.Read failed`)
		}
	}
	return nil
}

// BeDecodeToInt8 decodes binary data into an int8.
func BeDecodeToInt8(b []byte) int8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return int8(b[0])
}

// BeDecodeToUint8 decodes binary data into a uint8.
func BeDecodeToUint8(b []byte) uint8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return b[0]
}

// BeDecodeToInt16 decodes binary data into an int16.
func BeDecodeToInt16(b []byte) int16 { return int16(binary.BigEndian.Uint16(BeFillUpSize(b, 2))) }

// BeDecodeToUint16 decodes binary data into a uint16.
func BeDecodeToUint16(b []byte) uint16 { return binary.BigEndian.Uint16(BeFillUpSize(b, 2)) }

// BeDecodeToFloat32 decodes binary data into a float32.
func BeDecodeToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(BeFillUpSize(b, 4)))
}

// BeDecodeToFloat64 decodes binary data into a float64.
func BeDecodeToFloat64(b []byte) float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(BeFillUpSize(b, 8)))
}

// BeFillUpSize ensures a byte slice has the required length by padding it with zeros if necessary.
func BeFillUpSize(b []byte, l int) []byte {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c[l-len(b):], b)
	return c
}
