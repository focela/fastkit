// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package utils provides utility functions for internal usage.
package utils

import (
	"github.com/focela/aegis/internal/command"
)

const (
	// debugKeyEnv is the key for enabling debug mode via command line or environment variable.
	debugKeyEnv = "aegis.debug"
)

var (
	// isDebugEnabled indicates whether the debug mode is enabled.
	isDebugEnabled = false
)

func init() {
	// Configure debug mode based on command line arguments or environment variables.
	value := command.GetOptWithEnv(debugKeyEnv)
	isDebugEnabled = !(value == "" || value == "0" || value == "false")
}

// IsDebugEnabled checks whether debug mode is enabled.
//
// Debug mode is enabled when the command line argument "aegis.debug" or the environment variable "AEGIS_DEBUG" is set.
func IsDebugEnabled() bool {
	return isDebugEnabled
}

// SetDebugEnabled allows enabling or disabling debug mode programmatically.
//
// Parameters:
// - enabled: A boolean value to set the debug mode.
func SetDebugEnabled(enabled bool) {
	isDebugEnabled = enabled
}
