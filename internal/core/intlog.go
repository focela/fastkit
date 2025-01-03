// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package core provides essential utilities and foundational tools for the application.
package core

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/focela/loom/internal/debug"
	"github.com/focela/loom/internal/utils"
)

const (
	stackFilterKey = "/internal/core/intlog"
)

// Print logs messages in debug mode with newline.
func Print(ctx context.Context, v ...interface{}) {
	doPrint(ctx, fmt.Sprint(v...), false)
}

// Printf logs formatted messages in debug mode.
func Printf(ctx context.Context, format string, v ...interface{}) {
	doPrint(ctx, fmt.Sprintf(format, v...), false)
}

// PrintFunc executes and logs output from a function in debug mode.
func PrintFunc(ctx context.Context, f func() string) {
	if s := f(); s != "" {
		doPrint(ctx, s, false)
	}
}

// Error logs error messages in debug mode with newline.
func Error(ctx context.Context, v ...interface{}) {
	doPrint(ctx, fmt.Sprint(v...), true)
}

// Errorf logs formatted error messages in debug mode.
func Errorf(ctx context.Context, format string, v ...interface{}) {
	doPrint(ctx, fmt.Sprintf(format, v...), true)
}

// ErrorFunc executes and logs error output from a function in debug mode.
func ErrorFunc(ctx context.Context, f func() string) {
	if s := f(); s != "" {
		doPrint(ctx, s, true)
	}
}

// doPrint handles the actual printing of log messages.
func doPrint(ctx context.Context, content string, stack bool) {
	if !utils.IsDebugEnabled() {
		return
	}

	buffer := bytes.NewBuffer(nil)
	buffer.WriteString(time.Now().Format("2006-01-02 15:04:05.000"))
	buffer.WriteString(" [INTE] ")
	buffer.WriteString(getCallerFileInfo())
	buffer.WriteString(" ")

	if traceID := getTraceID(ctx); traceID != "" {
		buffer.WriteString(traceID + " ")
	}

	buffer.WriteString(content)
	buffer.WriteString("\n")

	if stack {
		buffer.WriteString("Caller Stack:\n")
		buffer.WriteString(debug.StackWithFilter([]string{stackFilterKey}))
	}

	fmt.Print(buffer.String())
}

// getTraceID retrieves the trace ID from the context.
func getTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	spanCtx := trace.SpanContextFromContext(ctx)
	if traceID := spanCtx.TraceID(); traceID.IsValid() {
		return "{" + traceID.String() + "}"
	}
	return ""
}

// getCallerFileInfo returns the caller's file name and line number.
func getCallerFileInfo() string {
	_, path, line := debug.CallerWithFilter([]string{stackFilterKey})
	return fmt.Sprintf("%s:%d", filepath.Base(path), line)
}
