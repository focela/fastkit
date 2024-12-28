// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"io"
)

// ReadCloser implements io.ReadCloser for reading repeatable content.
type ReadCloser struct {
	index      int    // Current read position.
	content    []byte // Content to be read.
	repeatable bool   // Indicates if content can be read repeatedly.
}

// NewReadCloser creates a new ReadCloser instance.
//
// Parameters:
// - content: Byte slice representing the content to be read.
// - repeatable: If true, allows the content to be read repeatedly.
//
// Returns:
// - An io.ReadCloser instance.
func NewReadCloser(content []byte, repeatable bool) io.ReadCloser {
	return &ReadCloser{
		content:    content,
		repeatable: repeatable,
	}
}

// Read reads data into p from the content.
//
// Implements io.ReadCloser interface.
//
// Parameters:
// - p: A byte slice to store the read data.
//
// Returns:
// - n: Number of bytes read.
// - err: EOF if end of content is reached.
func (b *ReadCloser) Read(p []byte) (n int, err error) {
	if b.index >= len(b.content) && b.repeatable {
		b.index = 0 // Reset index if repeatable is enabled.
	}

	n = copy(p, b.content[b.index:])
	b.index += n

	if b.index >= len(b.content) {
		return n, io.EOF
	}
	return n, nil
}

// Close is a no-op close method.
//
// Implements io.ReadCloser interface.
//
// Returns:
// - Always returns nil.
func (b *ReadCloser) Close() error {
	return nil
}
