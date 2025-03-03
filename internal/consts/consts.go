// Copyright (c) 2025 Focela Technologies.
// This software is provided "as is", without any warranty.
// Licensed under the MIT License – see LICENSE file for details.

// Package consts defines constants that are shared among all packages of the framework.
package consts

// Framework constants define configuration node names and identifiers used throughout the Aegis framework.
const (
	// ConfigNodeNameDatabase is the configuration node name for database settings.
	ConfigNodeNameDatabase = "database"

	// ConfigNodeNameLogger is the configuration node name for logger settings.
	ConfigNodeNameLogger = "logger"

	// ConfigNodeNameRedis is the configuration node name for Redis settings.
	ConfigNodeNameRedis = "redis"

	// ConfigNodeNameViewer is the configuration node name for viewer settings.
	ConfigNodeNameViewer = "viewer"

	// ConfigNodeNameServer is the configuration node name for primary server settings.
	ConfigNodeNameServer = "server"

	// ConfigNodeNameServerSecondary is the configuration node name for secondary HTTP server settings.
	ConfigNodeNameServerSecondary = "httpserver"

	// StackFilterKeyForAegis is the identifier used for stack trace filtering.
	StackFilterKeyForAegis = "github.com/focela/aegis/"
)
