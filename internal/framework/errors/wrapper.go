// Copyright (c) 2024 Focela Technologies
// Internal Use Only - Unauthorized use prohibited
// Contact: legal@focela.com

package errors

import (
	"github.com/focela/altura/internal/framework/command"
)

// StackMode defines the mode for printing stack information.
type StackMode string

const (
	// Deprecated: use CommandEnvKeyForStackMode instead.
	CommandEnvKeyForBrief = "altura.error.brief"

	// CommandEnvKeyForStackMode is the key for configuring the stack mode via environment variables.
	CommandEnvKeyForStackMode = "altura.error.stack.mode"
)

const (
	// StackModeBrief indicates that error stacks will exclude framework-specific details.
	StackModeBrief StackMode = "brief"

	// StackModeDetail indicates that error stacks will include detailed framework-specific information.
	StackModeDetail StackMode = "detail"
)

// stackModeConfigured holds the current error stack mode.
// By default, it is set to StackModeBrief.
var stackModeConfigured = StackModeBrief

func init() {
	// Deprecated: Support for the legacy brief mode configuration.
	briefSetting := command.GetOptWithEnv(CommandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Configure the error stack mode using command-line arguments or environment variables.
	stackModeSetting := command.GetOptWithEnv(CommandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackModeSettingMode := StackMode(stackModeSetting)
		switch stackModeSettingMode {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = stackModeSettingMode
		}
	}
}

// IsStackModeBrief returns whether the current error stack mode is in brief mode.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
