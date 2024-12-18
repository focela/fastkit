// Copyright (c) 2024 Focela Technologies
// Internal Use Only - Unauthorized use prohibited
// Contact: legal@focela.com

// Package constants defines constants that are shared across all framework packages.
package constants

// ConfigNodeNames define keys for configuration nodes.
const (
	ConfigNodeNameDatabase        = "database"   // Key for database configuration settings.
	ConfigNodeNameLogger          = "logger"     // Key for logger configuration settings.
	ConfigNodeNameRedis           = "redis"      // Key for Redis cache settings.
	ConfigNodeNameViewer          = "viewer"     // Key for viewer-specific settings.
	ConfigNodeNameServer          = "server"     // Key for general server settings.
	ConfigNodeNameServerSecondary = "httpserver" // Key for secondary server settings (introduced in v2).
)

// StackFilterKeys define keys for stack trace filtering.
const (
	StackFilterKeyForAltura = "github.com/focela/altura/" // Key to identify stack traces belonging to Altura framework.
)
