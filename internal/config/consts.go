// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package config provides configuration constants and related utilities.
package config

// Configuration node names used across the application.
const (
	// Database configuration node.
	ConfigNodeNameDatabase = "database"
	// Logger configuration node.
	ConfigNodeNameLogger = "logger"
	// Redis configuration node.
	ConfigNodeNameRedis = "redis"
	// Viewer configuration node.
	ConfigNodeNameViewer = "viewer"
	// Primary server configuration node.
	ConfigNodeNameServer = "server"
	// Secondary HTTP server configuration node.
	ConfigNodeNameServerSecondary = "httpserver"
)

// Special filter keys for stack tracing.
const (
	// Key used to filter stack traces for Loom package.
	StackFilterKeyForLoom = "github.com/focela/loom/"
)
