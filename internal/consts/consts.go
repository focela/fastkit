// Copyright (c) 2025 Focela Technologies.
// This software is provided "as is", without any warranty.
// Licensed under the MIT License – see LICENSE file for details.

// Package consts defines constants that are shared all among packages of framework.
package consts

const (
	ConfigNodeNameDatabase        = "database"
	ConfigNodeNameLogger          = "logger"
	ConfigNodeNameRedis           = "redis"
	ConfigNodeNameViewer          = "viewer"
	ConfigNodeNameServer          = "server"     // General version configuration item name.
	ConfigNodeNameServerSecondary = "httpserver" // New version configuration item name support from v2.
	StackFilterKeyForAegis        = "github.com/focela/aegis/"
)
