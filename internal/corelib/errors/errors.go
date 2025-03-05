// Copyright (c) 2025 Focela Technologies.
// This software is provided "as is", without any warranty.
// Licensed under the MIT License – see LICENSE file for details.

// Package errors provides enhanced error handling functionalities for the Aegis framework.
// It includes error stack management with configurable detail levels.
package errors

import (
	"github.com/focela/aegis/internal/command"
)

// StackMode represents the level of detail for error stack traces.
// It controls whether framework-internal stack frames are included in error reporting.
type StackMode string

// Stack mode constants define the available detail levels for error stack traces.
const (
	// StackModeBrief displays error stacks without framework-internal details.
	// This is the default mode and is more suitable for application-level error reporting.
	StackModeBrief StackMode = "brief"

	// StackModeDetail displays complete error stacks including all framework-internal details.
	// This is useful for debugging framework-level issues.
	StackModeDetail StackMode = "detail"
)

// Command environment key constants for configuring error stack behavior.
const (
	// commandEnvKeyForBrief is the legacy environment variable name for brief error stack setting.
	// Deprecated: use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "aegis.error.brief"

	// commandEnvKeyForStackMode is the environment variable name for configuring error stack mode.
	commandEnvKeyForStackMode = "aegis.error.stack.mode"
)

// stackModeConfigured holds the currently active stack trace mode.
// By default, it uses brief mode to hide framework internal details.
var stackModeConfigured = StackModeBrief

func init() {
	// First check legacy brief setting (deprecated)
	briefSetting := command.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Then check the newer stack mode setting which takes precedence
	stackModeSetting := command.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		requestedMode := StackMode(stackModeSetting)
		switch requestedMode {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = requestedMode
		}
	}
}

// IsStackModeBrief returns true if the current error stack mode is set to brief mode.
// Brief mode omits framework-internal stack frames from error stack traces.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
