// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package charset provides APIs for character-set conversion functionality.
//
// Supported Character Set:
//
// - Chinese : GBK/GB18030/GB2312/Big5
// - Japanese: EUCJP/ISO2022JP/ShiftJIS
// - Korean  : EUCKR
// - Unicode : UTF-8/UTF-16/UTF-16BE/UTF-16LE
// - Other   : macintosh/IBM*/Windows*/ISO-*
package charset

import (
	"bytes"
	"context"
	"io"

	"github.com/focela/loom/internal/core"
	"github.com/focela/loom/pkg/errors"
	"github.com/focela/loom/pkg/errors/code"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
)

// charsetAlias maps alternative charset names to their standard representation.
var charsetAlias = map[string]string{
	"HZGB2312": "HZ-GB-2312",
	"hzgb2312": "HZ-GB-2312",
	"GB2312":   "HZ-GB-2312",
	"gb2312":   "HZ-GB-2312",
}

// Supported checks whether a given charset is supported.
func Supported(charset string) bool {
	return getEncoding(charset) != nil
}

// Convert converts a string from `srcCharset` to `dstCharset`.
// Returns the converted string and an error if the conversion fails.
func Convert(dstCharset string, srcCharset string, src string) (dst string, err error) {
	if dstCharset == srcCharset {
		return src, nil
	}
	dst = src

	// Convert `src` to UTF-8 if it's not already UTF-8.
	if srcCharset != "UTF-8" {
		src, err = decodeToUTF8(srcCharset, src)
		if err != nil {
			return "", err
		}
	}

	// Convert from UTF-8 to `dstCharset` if needed.
	if dstCharset != "UTF-8" {
		dst, err = encodeFromUTF8(dstCharset, src)
		if err != nil {
			return "", err
		}
	}

	return dst, nil
}

// ToUTF8 converts a string from `srcCharset` to UTF-8.
func ToUTF8(srcCharset string, src string) (string, error) {
	return Convert("UTF-8", srcCharset, src)
}

// UTF8To converts a string from UTF-8 to `dstCharset`.
func UTF8To(dstCharset string, src string) (string, error) {
	return Convert(dstCharset, "UTF-8", src)
}

// decodeToUTF8 decodes a string from `srcCharset` to UTF-8.
func decodeToUTF8(srcCharset string, src string) (string, error) {
	enc := getEncoding(srcCharset)
	if enc == nil {
		return "", errors.NewCodef(code.CodeInvalidParameter, "unsupported srcCharset '%s'", srcCharset)
	}
	tmp, err := io.ReadAll(
		transform.NewReader(bytes.NewReader([]byte(src)), enc.NewDecoder()),
	)
	if err != nil {
		return "", errors.Wrapf(err, "convert string '%s' to utf8 failed", srcCharset)
	}
	return string(tmp), nil
}

// encodeFromUTF8 encodes a string from UTF-8 to `dstCharset`.
func encodeFromUTF8(dstCharset string, src string) (string, error) {
	enc := getEncoding(dstCharset)
	if enc == nil {
		return "", errors.NewCodef(code.CodeInvalidParameter, "unsupported dstCharset '%s'", dstCharset)
	}
	tmp, err := io.ReadAll(
		transform.NewReader(bytes.NewReader([]byte(src)), enc.NewEncoder()),
	)
	if err != nil {
		return "", errors.Wrapf(err, "convert string from utf8 to '%s' failed", dstCharset)
	}
	return string(tmp), nil
}

// getEncoding retrieves the encoding.Encoding interface for a given charset.
// Returns nil if the charset is unsupported.
func getEncoding(charset string) encoding.Encoding {
	if c, ok := charsetAlias[charset]; ok {
		charset = c
	}
	enc, err := ianaindex.MIB.Encoding(charset)
	if err != nil {
		core.Errorf(context.TODO(), "%+v", err)
	}
	return enc
}
