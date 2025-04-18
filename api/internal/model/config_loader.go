// Copyright 2025 Focela – MIT Licensed. See LICENSE.
// Part of the Focela Open Initiative.
// Built to empower developers and partners.

// Package model provides data structures and operations for the application's data layer.
package model

// ServeLogConfig defines service logging configuration.
type ServeLogConfig struct {
	// Switch enables/disables logging
	Switch bool `json:"switch"`

	// Queue enables queue-based processing
	Queue bool `json:"queue"`

	// LevelFormat defines format for different log levels
	LevelFormat []string `json:"levelFormat"`
}
