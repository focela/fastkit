// Copyright (c) 2024 Focela Technologies
// Internal Use Only - Unauthorized use prohibited
// Contact: legal@focela.com

// Package errors provides functionalities to manipulate errors for internal usage purpose.
package errors

import (
	"github.com/focela/altura/internal/framework/cli"
)

// StackMode defines the mode for printing stack information.
type StackMode string

// Constants for environment keys.
const (
	// EnvKeyForBrief is the environment variable key for switching to brief error stack.
	// Deprecated: use EnvKeyForStackMode instead.
	EnvKeyForBrief = "altura.error.brief"

	// EnvKeyForStackMode is the environment variable key for configuring error stack mode.
	EnvKeyForStackMode = "altura.error.stack.mode"
)

// Constants for stack modes.
const (
	// StackModeBrief specifies printing error stacks without framework details.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies printing detailed error stacks including framework details.
	StackModeDetail StackMode = "detail"
)

// stackModeConfigured is the currently configured error stack mode.
// Defaults to StackModeBrief.
var stackModeConfigured = StackModeBrief

// init initializes the stack mode configuration based on environment variables or command-line arguments.
func init() {
	// Check deprecated brief setting.
	briefSetting := cli.GetOptWithEnv(EnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Configure stack mode using the new environment key.
	stackModeSetting := cli.GetOptWithEnv(EnvKeyForStackMode)
	if stackModeSetting != "" {
		stackMode := StackMode(stackModeSetting)
		switch stackMode {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = stackMode
		}
	}
}

// IsStackModeBrief checks whether the current error stack mode is set to brief.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
