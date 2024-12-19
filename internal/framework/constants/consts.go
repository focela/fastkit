// Copyright (c) 2024 Focela Technologies
// Internal Use Only - Unauthorized use prohibited
// Contact: legal@focela.com

package constants

// Configuration node names used for application configuration files.
const (
	// ConfigNodeNameDatabase represents the node name for database configuration.
	ConfigNodeNameDatabase = "database"

	// ConfigNodeNameLogger represents the node name for logger configuration.
	ConfigNodeNameLogger = "logger"

	// ConfigNodeNameRedis represents the node name for Redis configuration.
	ConfigNodeNameRedis = "redis"

	// ConfigNodeNameViewer represents the node name for viewer configuration.
	ConfigNodeNameViewer = "viewer"

	// ConfigNodeNameServer represents the node name for the primary server configuration.
	ConfigNodeNameServer = "server"

	// ConfigNodeNameServerSecondary represents the node name for the secondary HTTP server configuration.
	ConfigNodeNameServerSecondary = "httpserver"
)

// Framework-specific constants for internal processing.
const (
	// StackFilterKeyForAltura is a key used to filter stack traces in the application.
	StackFilterKeyForAltura = "github.com/focela/altura/"
)
