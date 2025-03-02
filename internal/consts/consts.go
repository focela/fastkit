// Copyright (c) 2025 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package consts defines constants that are shared among all packages of the framework.
package consts

// Configuration node names for different components of the framework.
// These are used to access specific sections in configuration files.
const (
	// Database configuration node name.
	ConfigNodeNameDatabase = "database"

	// Logger configuration node name.
	ConfigNodeNameLogger = "logger"

	// Redis configuration node name.
	ConfigNodeNameRedis = "redis"

	// Viewer configuration node name for template rendering.
	ConfigNodeNameViewer = "viewer"
)

// Server configuration node names.
// The framework supports two different naming conventions for server configuration.
const (
	// ConfigNodeNameServer is the general version configuration node name.
	ConfigNodeNameServer = "server"

	// ConfigNodeNameServerSecondary is the alternative configuration node name
	// supported from v2 onwards.
	ConfigNodeNameServerSecondary = "httpserver"
)

// Framework-specific constants for internal operations.
const (
	// StackFilterKeyForAegis is the stack filtering key for all Aegis module paths.
	// Used for error stack filtering to improve error readability.
	StackFilterKeyForAegis = "github.com/focela/aegis/"
)
