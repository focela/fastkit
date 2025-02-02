// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package consts defines constants that are shared among all packages of the framework.
package consts

// Configuration node names.
const (
	ConfigNodeDatabase        = "database"
	ConfigNodeLogger          = "logger"
	ConfigNodeRedis           = "redis"
	ConfigNodeViewer          = "viewer"
	ConfigNodeServer          = "server"
	ConfigNodeServerSecondary = "httpserver"
)

// StackFilterKeyAegis is the stack filter key used for filtering Aegis-specific stack traces.
const StackFilterKeyAegis = "github.com/focela/aegis/"
