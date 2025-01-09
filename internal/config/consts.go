// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package config provides configuration constants for the Aegis application.
package config

// Configuration node names used in the application.
const (
	// ConfigNodeNameDatabase represents the configuration node for database settings.
	ConfigNodeNameDatabase = "database"

	// ConfigNodeNameLogger represents the configuration node for logger settings.
	ConfigNodeNameLogger = "logger"

	// ConfigNodeNameRedis represents the configuration node for Redis settings.
	ConfigNodeNameRedis = "redis"

	// ConfigNodeNameViewer represents the configuration node for viewer settings.
	ConfigNodeNameViewer = "viewer"

	// ConfigNodeNameServer represents the configuration node for primary server settings.
	ConfigNodeNameServer = "server"

	// ConfigNodeNameServerSecondary represents the configuration node for secondary HTTP server settings.
	ConfigNodeNameServerSecondary = "httpserver"
)

// Filter keys used for stack tracing or filtering.
const (
	// StackFilterKeyForAegis represents the key used to filter stack traces specific to the Aegis application.
	StackFilterKeyForAegis = "github.com/focela/aegis"
)
