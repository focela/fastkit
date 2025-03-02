// Copyright (c) 2025 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package errors provides functionalities to manipulate and control error stack traces
// for better debugging and user experience.
package errors

import (
	"github.com/focela/aegis/internal/command"
)

// StackMode defines how error stack traces are displayed: brief or detailed.
type StackMode string

// Stack mode constants define the available display options for error stack traces.
const (
	// StackModeBrief hides framework internal stack frames for cleaner error output.
	StackModeBrief StackMode = "brief"

	// StackModeDetail shows complete stack traces including framework internals.
	StackModeDetail StackMode = "detail"
)

// Environment variable keys used to configure error stack behavior.
const (
	// commandEnvKeyForBrief is the legacy environment variable for enabling brief stack mode.
	// Deprecated: use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "aegis.error.brief"

	// commandEnvKeyForStackMode is the current environment variable for setting stack mode.
	commandEnvKeyForStackMode = "aegis.error.stack.mode"
)

// stackModeConfigured stores the current stack mode configuration.
// Default is brief mode for cleaner output.
var stackModeConfigured = StackModeBrief

func init() {
	// First check for legacy brief setting (deprecated)
	briefSetting := command.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Then check for the newer stack mode setting which takes precedence
	stackModeSetting := command.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackModeSettingMode := StackMode(stackModeSetting)

		// Only accept valid modes
		switch stackModeSettingMode {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = stackModeSettingMode
		}
	}
}

// IsStackModeBrief returns whether the current error stack mode is set to brief.
// This is used by error handling code to determine how much stack information to include.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
