// Copyright 2025 Focela – MIT Licensed. See LICENSE.
// Part of the Focela Open Initiative.
// Built to empower developers and partners.

// Package main is the entry point for the Fastkit API application.
package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"fastkit/internal/global"
)

// main initializes and starts the application.
func main() {
	// Get the initialization context from GoFrame
	ctx := gctx.GetInitCtx()

	// Initialize global application settings
	global.Init(ctx)
}
