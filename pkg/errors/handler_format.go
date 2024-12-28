// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package errors provides rich functionalities to manipulate errors.
package errors

import (
	"fmt"
	"io"
)

// Format provides custom formatting for the Error struct based on fmt.Formatter.
//
// Supported verbs:
// %v, %s   : Print the complete error string.
// %-v, %-s : Print only the current error message.
// %+s      : Print the full stack trace.
// %+v      : Print the error string along with the full stack trace.
//
// Usage Examples:
// fmt.Sprintf("%v", err)   -> Full error message
// fmt.Sprintf("%-v", err)  -> Current error message
// fmt.Sprintf("%+v", err)  -> Error + stack trace
func (err *Error) Format(state fmt.State, verb rune) {
	switch verb {
	case 's', 'v': // Handle %s and %v verbs.
		if state.Flag('-') { // %-s, %-v: Only current error message.
			if err.text != "" {
				_, _ = io.WriteString(state, err.text)
			} else {
				_, _ = io.WriteString(state, err.Error())
			}
			return
		}

		if state.Flag('+') { // %+s, %+v: Full stack trace.
			if verb == 's' {
				_, _ = io.WriteString(state, err.Stack())
			} else {
				_, _ = io.WriteString(state, fmt.Sprintf("%s\n%s", err.Error(), err.Stack()))
			}
			return
		}

		// Default: %s, %v without flags.
		_, _ = io.WriteString(state, err.Error())
	}
}
