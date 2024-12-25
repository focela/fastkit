// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package core provides essential utilities and foundational tools for the application.
package core

import (
	"github.com/focela/loom/internal/command"
)

// StackMode represents the mode for printing stack information.
type StackMode string

// Environment variable keys for stack mode configuration.
const (
	// commandEnvKeyForBrief is the command environment name for switch key for brief error stack.
	// Deprecated: use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "loom.error.brief"

	// commandEnvKeyForStackMode is the command environment name for switch key for error stack mode.
	commandEnvKeyForStackMode = "loom.error.stack.mode"
)

// Stack mode constants.
const (
	// StackModeBrief specifies printing error stacks without framework error stacks.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies printing detailed error stacks, including framework stacks.
	StackModeDetail StackMode = "detail"
)

// stackModeConfigured holds the currently configured error stack mode.
// Defaults to StackModeBrief.
var stackModeConfigured = StackModeBrief

// init initializes the error stack mode configuration.
// It reads settings from command line arguments or environment variables.
func init() {
	// Handle deprecated brief stack mode setting.
	briefSetting := command.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Handle stack mode setting from environment or arguments.
	stackModeSetting := command.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackModeSettingMode := StackMode(stackModeSetting)
		switch stackModeSettingMode {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = stackModeSettingMode
		}
	}
}

// IsStackModeBrief returns whether the current error stack mode is set to brief.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
