// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"github.com/focela/loom/internal/command"
)

// Constants for debug configuration.
const (
	// commandEnvKeyForDebugKey is used to check if debug mode is enabled.
	commandEnvKeyForDebugKey = "loom.debug"
)

// isDebugEnabled indicates whether Loom debug mode is enabled.
var isDebugEnabled = false

// init initializes the debug mode configuration based on environment variables or command-line arguments.
func init() {
	value := command.GetOptWithEnv(commandEnvKeyForDebugKey)
	isDebugEnabled = !(value == "" || value == "0" || value == "false")
}

// IsDebugEnabled checks if debug mode is enabled.
// Debug mode is activated via the "loom.debug" command argument or environment variable.
func IsDebugEnabled() bool {
	return isDebugEnabled
}

// SetDebugEnabled enables or disables the debug mode.
func SetDebugEnabled(enabled bool) {
	isDebugEnabled = enabled
}
