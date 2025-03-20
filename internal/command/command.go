// Copyright (c) 2025 Focela Technologies.
// This software is provided "as is", without any warranty.
// Licensed under the MIT License – see LICENSE file for details.

// Package command provides utilities for parsing command-line arguments and options.
package command

import (
	"os"
	"regexp"
	"strings"
)

// Matches options in formats: -flag, --flag, -flag=value, --flag=value
var (
	defaultParsedArgs    []string
	defaultParsedOptions map[string]string
	argumentOptionRegex  = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=){0,1}(.*)$`)
)

// Init initializes the package with provided arguments or os.Args if none given.
func Init(args ...string) {
	if len(args) == 0 {
		if len(defaultParsedArgs) > 0 || len(defaultParsedOptions) > 0 {
			return
		}
		args = os.Args
	} else {
		defaultParsedArgs = nil
		defaultParsedOptions = make(map[string]string)
	}

	defaultParsedArgs, defaultParsedOptions = ParseUsingDefaultAlgorithm(args...)
}

// ParseUsingDefaultAlgorithm separates command arguments into args and options.
// Handles formats: -name=value, --name=value, -name value, --name value, or -flag, --flag
func ParseUsingDefaultAlgorithm(args ...string) ([]string, map[string]string) {
	parsedArgs := make([]string, 0, len(args))
	parsedOptions := make(map[string]string)

	for i := 0; i < len(args); {
		matches := argumentOptionRegex.FindStringSubmatch(args[i])

		if len(matches) > 2 {
			optName := matches[1]
			hasEquals := matches[2] == "="
			optValue := matches[3]

			if hasEquals {
				// -name=value format
				parsedOptions[optName] = optValue
			} else if i < len(args)-1 && (len(args[i+1]) == 0 || args[i+1][0] != '-') {
				// -name value format
				parsedOptions[optName] = args[i+1]
				i += 2
				continue
			} else {
				// -flag (without value)
				parsedOptions[optName] = optValue
			}
		} else {
			// Regular argument
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}

	return parsedArgs, parsedOptions
}

// GetOpt returns the value of option 'name' or default/empty if not found.
func GetOpt(name string, def ...string) string {
	Init()

	if v, ok := defaultParsedOptions[name]; ok {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return ""
}

// GetOptAll returns all parsed options.
func GetOptAll() map[string]string {
	Init()
	return defaultParsedOptions
}

// ContainsOpt checks if option 'name' exists.
func ContainsOpt(name string) bool {
	Init()
	_, ok := defaultParsedOptions[name]
	return ok
}

// GetArg returns the argument at 'index' or default/empty if not found.
func GetArg(index int, def ...string) string {
	Init()

	if index < len(defaultParsedArgs) {
		return defaultParsedArgs[index]
	}

	if len(def) > 0 {
		return def[0]
	}

	return ""
}

// GetArgAll returns all parsed arguments.
func GetArgAll() []string {
	Init()
	return defaultParsedArgs
}

// GetOptWithEnv returns option value or environment variable or default.
// Options: lowercase with dots (app.setting.name)
// Env vars: uppercase with underscores (APP_SETTING_NAME)
func GetOptWithEnv(key string, def ...string) string {
	// Command line format (lowercase with dots)
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))

	if ContainsOpt(cmdKey) {
		return GetOpt(cmdKey)
	}

	// Environment variable format (uppercase with underscores)
	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))

	if value, exists := os.LookupEnv(envKey); exists {
		return value
	}

	if len(def) > 0 {
		return def[0]
	}

	return ""
}
