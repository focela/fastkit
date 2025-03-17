// Copyright (c) 2024 Focela Technologies. All rights reserved.
// This source code is governed by an MIT License.
// See LICENSE file for full terms and conditions.

// Package config defines constants for configuration nodes and other settings used in the Aegis project.
package config

// Config node names used for accessing various configuration sections.
const (
	ConfigNodeNameDatabase        = "database"
	ConfigNodeNameLogger          = "logger"
	ConfigNodeNameRedis           = "redis"
	ConfigNodeNameViewer          = "viewer"
	ConfigNodeNameServer          = "server"
	ConfigNodeNameServerSecondary = "httpserver"
)

// Stack filtering key for Aegis.
const (
	StackFilterKeyForAegis = "github.com/focela/aegis/"
)
