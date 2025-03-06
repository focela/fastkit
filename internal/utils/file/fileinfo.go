// Copyright (c) 2025 Focela Technologies.
// This software is provided "as is", without any warranty.
// Licensed under the MIT License – see LICENSE file for details.

// Package file provides file system utility functions for the Aegis framework.
// This includes implementations of os.FileInfo for virtual file operations.
package file

import (
	"os"
	"time"
)

// Info implements os.FileInfo interface for virtual file operations.
// It allows creating file information objects without actual files on disk.
type Info struct {
	name    string      // Name of the file
	size    int64       // Size in bytes
	mode    os.FileMode // File mode and permissions
	modTime time.Time   // Modification time
}

// New creates a new Info instance with the specified file attributes.
// This constructor helps create virtual file information objects for testing or other purposes.
//
// Parameters:
//   - name: the name of the file (without path information)
//   - size: the file size in bytes
//   - mode: the file mode and permissions
//   - modTime: the file modification time
//
// Returns a pointer to the newly created Info struct.
func New(name string, size int64, mode os.FileMode, modTime time.Time) *Info {
	return &Info{
		name:    name,
		size:    size,
		mode:    mode,
		modTime: modTime,
	}
}

// Name returns the base name of the file.
// This implements part of the os.FileInfo interface.
func (i *Info) Name() string {
	return i.name
}

// Size returns the length in bytes of the file.
// This implements part of the os.FileInfo interface.
func (i *Info) Size() int64 {
	return i.size
}

// IsDir returns whether the file is a directory.
// This implements part of the os.FileInfo interface.
func (i *Info) IsDir() bool {
	return i.mode.IsDir()
}

// Mode returns the file mode bits.
// This implements part of the os.FileInfo interface.
func (i *Info) Mode() os.FileMode {
	return i.mode
}

// ModTime returns the modification time of the file.
// This implements part of the os.FileInfo interface.
func (i *Info) ModTime() time.Time {
	return i.modTime
}

// Sys returns the underlying data source (can return nil).
// This implements part of the os.FileInfo interface.
// This implementation always returns nil as there is no underlying system-specific data.
func (i *Info) Sys() interface{} {
	return nil
}
