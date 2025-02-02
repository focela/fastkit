// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package errors provides functionalities to manipulate errors for internal usage.
package errors

import (
	"github.com/focela/aegis/internal/command"
)

// StackMode defines the level of detail for printing error stack traces.
type StackMode string

const (
	// commandEnvKeyForBrief is the environment key for brief error stack mode.
	// Deprecated: Use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "aegis.error.brief"

	// commandEnvKeyForStackMode is the environment key for stack mode selection.
	commandEnvKeyForStackMode = "aegis.error.stack.mode"
)

const (
	// StackModeBrief prints error stacks excluding framework-related errors.
	StackModeBrief StackMode = "brief"

	// StackModeDetail prints detailed error stacks including framework-related errors.
	StackModeDetail StackMode = "detail"
)

// stackModeConfigured stores the current configured error stack mode.
// Default value is StackModeBrief.
var stackModeConfigured = StackModeBrief

func init() {
	// Check for deprecated brief mode setting.
	briefSetting := command.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Check and apply stack mode configuration from environment or command-line arguments.
	stackModeSetting := command.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		mode := StackMode(stackModeSetting)
		if mode == StackModeBrief || mode == StackModeDetail {
			stackModeConfigured = mode
		}
	}
}

// IsStackModeBrief returns true if the current error stack mode is set to brief.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
