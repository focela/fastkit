// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package kernel provides core utilities and foundational components for the Aegis framework.
package kernel

import (
	"github.com/focela/aegis/internal/command"
)

// StackMode represents the mode for printing error stack information.
type StackMode string

const (
	// Deprecated constants
	//
	// commandEnvKeyForBrief is the environment variable for switching to brief error stack mode.
	// Deprecated: Use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "aegis.error.brief"

	// Constants for stack mode configuration.
	//
	// commandEnvKeyForStackMode is the environment variable for configuring error stack mode.
	commandEnvKeyForStackMode = "aegis.error.stack.mode"
)

const (
	// StackModeBrief specifies that error stacks exclude framework-related errors.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies that error stacks include detailed framework-related errors.
	StackModeDetail StackMode = "detail"
)

// stackModeConfigured holds the currently configured stack mode.
// The default mode is StackModeBrief.
var stackModeConfigured = StackModeBrief

// init initializes the error stack mode configuration by reading command-line arguments or environment variables.
func init() {
	// Check deprecated brief mode configuration.
	briefSetting := command.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Configure stack mode using the new environment variable or command-line arguments.
	stackModeSetting := command.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackMode := StackMode(stackModeSetting)
		if stackMode == StackModeBrief || stackMode == StackModeDetail {
			stackModeConfigured = stackMode
		}
	}
}

// IsStackModeBrief checks whether the current error stack mode is set to brief mode.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
