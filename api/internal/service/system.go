// Copyright 2025 Focela – MIT Licensed. See LICENSE.
// Part of the Focela Open Initiative.
// Built to empower developers and partners.

// Package service provides system configuration interfaces.
package service

import (
	"context"

	"fastkit/internal/model"
)

type (
	// SystemConfig defines interfaces for system configuration.
	SystemConfig interface {
		// GetServerLogConfig returns the server logging configuration.
		GetServerLogConfig(ctx context.Context) (*model.ServeLogConfig, error)
	}
)
