// Copyright 2025 Focela – MIT Licensed. See LICENSE.
// Part of the Focela Open Initiative.
// Built to empower developers and partners.

// Package global provides application-wide initialization and configuration functionality.
package global

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmode"

	"fastkit/utility/validate"
)

// Init initializes global application settings.
func Init(ctx context.Context) {
	// Configure GoFrame runtime mode
	SetGFMode(ctx)
}

// SetGFMode configures the GoFrame runtime mode based on application configuration.
// It reads the mode from configuration and sets it if valid.
func SetGFMode(ctx context.Context) {
	mode := g.Cfg().MustGet(ctx, "system.mode").String()
	if len(mode) == 0 {
		mode = gmode.NOT_SET
	}

	// Valid runtime modes
	validModes := []string{
		gmode.DEVELOP,
		gmode.TESTING,
		gmode.STAGING,
		gmode.PRODUCT,
	}

	// Set the mode if it's valid
	if validate.InSlice(validModes, mode) {
		gmode.Set(mode)
	}
}
