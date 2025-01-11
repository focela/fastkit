// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"io"
)

// ReadCloser implements the io.ReadCloser interface.
// It allows for reading content repeatedly if marked as repeatable.
type ReadCloser struct {
	index      int    // Current position in the content.
	content    []byte // Content to be read.
	repeatable bool   // If true, the content can be read repeatedly.
}

// NewReadCloser creates and returns a ReadCloser object.
//
// Parameters:
// - content: The byte slice to be read.
// - repeatable: A boolean indicating whether the content can be read multiple times.
//
// Returns:
// - An implementation of io.ReadCloser.
func NewReadCloser(content []byte, repeatable bool) io.ReadCloser {
	return &ReadCloser{
		content:    content,
		repeatable: repeatable,
	}
}

// Read implements the io.Reader interface. It reads data into p and updates the read index.
//
// If repeatable is true, the read position resets to the beginning when the end of the content is reached.
func (b *ReadCloser) Read(p []byte) (n int, err error) {
	// Reset index if repeatable is enabled and at the end of the content.
	if b.index >= len(b.content) {
		if b.repeatable {
			b.index = 0
		} else {
			return 0, io.EOF
		}
	}

	// Copy content to the buffer and update the index.
	n = copy(p, b.content[b.index:])
	b.index += n

	// Return EOF if the end of the content is reached.
	if b.index >= len(b.content) {
		return n, io.EOF
	}
	return n, nil
}

// Close implements the io.Closer interface. It is a no-op since no resources are allocated.
//
// Returns:
// - Always returns nil.
func (b *ReadCloser) Close() error {
	return nil
}
