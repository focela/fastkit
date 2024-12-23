/*
 * FOCELA TECHNOLOGIES INTERNAL USE ONLY LICENSE AGREEMENT
 *
 * Copyright (c) 2024 Focela Technologies. All rights reserved.
 *
 * Permission is hereby granted to employees or authorized personnel of Focela
 * Technologies (the "Company") to use this software solely for internal business
 * purposes within the Company.
 *
 * For inquiries or permissions, please contact: legal@focela.com
 */

// Package errors provides functionalities to manipulate errors for internal usage purposes.
package errors

import (
	"github.com/focela/altura/internal/cli"
)

// StackMode defines the mode for printing error stack information.
type StackMode string

// Command environment keys
const (
	// commandEnvKeyForBrief is the deprecated command environment key for brief error stack mode.
	// Deprecated: Use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "altura.error.brief"

	// commandEnvKeyForStackMode is the command environment key for configuring error stack mode.
	commandEnvKeyForStackMode = "altura.error.stack.mode"
)

// Stack modes
const (
	// StackModeBrief specifies brief error stack printing, excluding framework stack traces.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies detailed error stack printing, including framework stack traces.
	StackModeDetail StackMode = "detail"
)

// stackModeConfigured holds the current error stack mode configuration.
// Defaults to StackModeBrief.
var stackModeConfigured = StackModeBrief

// init initializes the error stack mode based on CLI arguments or environment variables.
func init() {
	// Check deprecated brief setting.
	briefSetting := cli.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Check current stack mode configuration.
	stackModeSetting := cli.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackModeSettingMode := StackMode(stackModeSetting)
		if stackModeSettingMode == StackModeBrief || stackModeSettingMode == StackModeDetail {
			stackModeConfigured = stackModeSettingMode
		}
	}
}

// IsStackModeBrief checks whether the current error stack mode is set to brief mode.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
