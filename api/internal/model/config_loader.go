// Copyright 2025 Focela – MIT Licensed. See LICENSE.
// Part of the Focela Open Initiative.
// Built to empower developers and partners.

// Package model provides data structures and operations for the application's data layer.
package model

// LogConfig defines general logging configuration settings.
type LogConfig struct {
	// Switch enables or disables the logging functionality.
	Switch bool `json:"switch"`

	// Queue enables queue-based log processing for improved performance.
	Queue bool `json:"queue"`

	// Module specifies which modules should be logged.
	Module []string `json:"module"`

	// SkipCode defines error codes that should be ignored during logging.
	SkipCode []string `json:"skipCode"`
}

// ServeLogConfig defines service-specific logging configuration.
type ServeLogConfig struct {
	// Switch enables or disables service logging.
	Switch bool `json:"switch"`

	// Queue enables queue-based processing for service logs.
	Queue bool `json:"queue"`

	// LevelFormat defines format specifications for different log levels.
	LevelFormat []string `json:"levelFormat"`
}
