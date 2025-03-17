// Copyright (c) 2024 Focela Technologies. All rights reserved.
// This source code is governed by an MIT License.
// See LICENSE file for full terms and conditions.

// Package errors configures error stack display modes for the Aegis project.
package errors

import (
	"github.com/focela/aegis/internal/cmd"
)

// StackMode defines the mode for printing error stack information.
type StackMode string

const (
	// Deprecated: use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "aegis.error.brief"

	// commandEnvKeyForStackMode is the environment variable key for error stack mode.
	commandEnvKeyForStackMode = "aegis.error.stack.mode"
)

const (
	// StackModeBrief specifies a brief error stack, excluding framework details.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies a detailed error stack, including framework details.
	StackModeDetail StackMode = "detail"
)

// stackModeConfigured holds the configured error stack mode; it defaults to brief.
var stackModeConfigured = StackModeBrief

func init() {
	// Deprecated: support legacy brief setting.
	if brief := cmd.GetOptWithEnv(commandEnvKeyForBrief); brief == "1" || brief == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Configure error stack mode from command-line or environment.
	if modeSetting := cmd.GetOptWithEnv(commandEnvKeyForStackMode); modeSetting != "" {
		if mode := StackMode(modeSetting); mode == StackModeBrief || mode == StackModeDetail {
			stackModeConfigured = mode
		}
	}
}

// IsStackModeBrief returns true if the current error stack mode is brief.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
