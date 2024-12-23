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

// Package errors provides functionalities to manipulate errors for internal usage purpose.
package errors

import (
	"github.com/focela/altura/internal/core/command"
)

// StackMode defines the mode for printing error stack information.
type StackMode string

const (
	// commandEnvKeyForBrief specifies the environment variable key for configuring brief error stack mode.
	// Deprecated: Use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "altura.error.brief"

	// commandEnvKeyForStackMode specifies the environment variable key for configuring the stack mode.
	commandEnvKeyForStackMode = "altura.error.stack.mode"
)

// StackMode options.
const (
	// StackModeBrief specifies minimal error stack information.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies detailed error stack information, including framework stacks.
	StackModeDetail StackMode = "detail"
)

// stackModeConfigured stores the configured error stack mode.
// Default value is StackModeBrief.
var stackModeConfigured = StackModeBrief

// configureStackMode initializes the error stack mode based on environment variables or command-line arguments.
func configureStackMode() {
	// Support for legacy brief setting.
	briefSetting := command.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Retrieve stack mode from environment variables or command line arguments.
	stackModeSetting := command.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackMode := StackMode(stackModeSetting)
		switch stackMode {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = stackMode
		}
	}
}

func init() {
	// Initialize stack mode configuration.
	configureStackMode()
}

// IsStackModeBrief checks whether the current stack mode is in brief mode.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
