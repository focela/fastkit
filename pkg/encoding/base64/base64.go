// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package b64 provides useful API for BASE64 encoding/decoding algorithm.
package b64

import (
	"encoding/base64"
	"os"

	"github.com/focela/loom/pkg/errors"
)

// Encode encodes bytes using the BASE64 algorithm.
func Encode(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

// EncodeToString encodes bytes into a BASE64 string.
func EncodeToString(src []byte) string {
	return string(Encode(src))
}

// EncodeString encodes a string using the BASE64 algorithm.
func EncodeString(src string) string {
	return EncodeToString([]byte(src))
}

// EncodeFile encodes the content of a file at `path` using BASE64.
func EncodeFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, `failed to read file "%s"`, path)
	}
	return Encode(content), nil
}

// MustEncodeFile encodes the content of a file and panics if an error occurs.
func MustEncodeFile(path string) []byte {
	result, err := EncodeFile(path)
	if err != nil {
		panic(err)
	}
	return result
}

// EncodeFileToString encodes the content of a file into a BASE64 string.
func EncodeFileToString(path string) (string, error) {
	content, err := EncodeFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// MustEncodeFileToString encodes the content of a file into a BASE64 string and panics if an error occurs.
func MustEncodeFileToString(path string) string {
	result, err := EncodeFileToString(path)
	if err != nil {
		panic(err)
	}
	return result
}

// Decode decodes bytes using the BASE64 algorithm.
func Decode(data []byte) ([]byte, error) {
	src := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(src, data)
	if err != nil {
		return nil, errors.Wrap(err, `failed to decode BASE64 bytes`)
	}
	return src[:n], nil
}

// MustDecode decodes bytes and panics if an error occurs.
func MustDecode(data []byte) []byte {
	result, err := Decode(data)
	if err != nil {
		panic(err)
	}
	return result
}

// DecodeString decodes a BASE64 string into bytes.
func DecodeString(data string) ([]byte, error) {
	return Decode([]byte(data))
}

// MustDecodeString decodes a BASE64 string into bytes and panics if an error occurs.
func MustDecodeString(data string) []byte {
	result, err := DecodeString(data)
	if err != nil {
		panic(err)
	}
	return result
}

// DecodeToString decodes a BASE64 string into a regular string.
func DecodeToString(data string) (string, error) {
	b, err := DecodeString(data)
	return string(b), err
}

// MustDecodeToString decodes a BASE64 string into a regular string and panics if an error occurs.
func MustDecodeToString(data string) string {
	result, err := DecodeToString(data)
	if err != nil {
		panic(err)
	}
	return result
}
