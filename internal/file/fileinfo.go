// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package file provides utilities for working with file metadata and information.
package file

import (
	"os"
	"time"
)

// Info represents metadata about a file.
type Info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// New creates a new Info instance with the provided file metadata.
func New(name string, size int64, mode os.FileMode, modTime time.Time) *Info {
	return &Info{
		name:    name,
		size:    size,
		mode:    mode,
		modTime: modTime,
	}
}

// Name returns the name of the file.
func (i *Info) Name() string {
	return i.name
}

// Size returns the size of the file in bytes.
func (i *Info) Size() int64 {
	return i.size
}

// IsDir reports whether the file is a directory.
func (i *Info) IsDir() bool {
	return i.mode.IsDir()
}

// Mode returns the file mode bits.
func (i *Info) Mode() os.FileMode {
	return i.mode
}

// ModTime returns the last modification time of the file.
func (i *Info) ModTime() time.Time {
	return i.modTime
}

// Sys returns underlying data source (not implemented in this struct).
func (i *Info) Sys() interface{} {
	return nil
}
