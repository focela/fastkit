// Copyright (c) 2025 Focela Technologies.
// This software is provided "as is", without any warranty.
// Licensed under the MIT License – see LICENSE file for details.

// Package command provides console operations, like options/arguments reading.
package command

import (
	"os"
	"regexp"
	"strings"
)

// argumentOptionRegex defines pattern to match command line options.
// Matches formats like: -option, --option, -option=value, --option=value
var argumentOptionRegex = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=){0,1}(.*)$`)

// Default storage for parsed arguments and options
var (
	defaultParsedArgs    = make([]string, 0)
	defaultParsedOptions = make(map[string]string)
)

// Init initializes the command package by parsing command-line arguments.
// If args is provided, it uses those instead of os.Args.
// If the function has already been called and args is empty, it does nothing.
func Init(args ...string) {
	if len(args) == 0 {
		if len(defaultParsedArgs) == 0 && len(defaultParsedOptions) == 0 {
			args = os.Args
		} else {
			return
		}
	} else {
		defaultParsedArgs = make([]string, 0)
		defaultParsedOptions = make(map[string]string)
	}

	// Parse arguments using default algorithm
	defaultParsedArgs, defaultParsedOptions = ParseUsingDefaultAlgorithm(args...)
}

// ParseUsingDefaultAlgorithm parses command-line arguments using the default algorithm.
// It separates arguments from options and returns them as separate collections.
// Returns:
//   - parsedArgs: slice of regular arguments
//   - parsedOptions: map of option names to their values
func ParseUsingDefaultAlgorithm(args ...string) (parsedArgs []string, parsedOptions map[string]string) {
	parsedArgs = make([]string, 0)
	parsedOptions = make(map[string]string)

	for i := 0; i < len(args); {
		array := argumentOptionRegex.FindStringSubmatch(args[i])
		if len(array) > 2 {
			if array[2] == "=" {
				// Format: -option=value or --option=value
				parsedOptions[array[1]] = array[3]
			} else if i < len(args)-1 {
				if len(args[i+1]) > 0 && args[i+1][0] == '-' {
					// Example: aegis gen -d -n 1
					parsedOptions[array[1]] = array[3]
				} else {
					// Example: aegis gen -n 2
					parsedOptions[array[1]] = args[i+1]
					i += 2
					continue
				}
			} else {
				// Example: aegis gen -h
				parsedOptions[array[1]] = array[3]
			}
		} else {
			// Regular argument (not an option)
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}
	return
}

// GetOpt returns the value of the option with the specified name.
// If the option doesn't exist, it returns the default value (if provided) or an empty string.
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

// GetOptAll returns all parsed options as a map of name-value pairs.
func GetOptAll() map[string]string {
	Init()
	return defaultParsedOptions
}

// ContainsOpt checks whether an option with the specified name exists in the parsed options.
func ContainsOpt(name string) bool {
	Init()
	_, ok := defaultParsedOptions[name]
	return ok
}

// GetArg returns the argument at the specified index.
// If the index is out of range, it returns the default value (if provided) or an empty string.
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

// GetArgAll returns all parsed arguments as a slice.
func GetArgAll() []string {
	Init()
	return defaultParsedArgs
}

// GetOptWithEnv returns the value from either command line or environment variable.
// Fetching Rules:
// 1. Command line arguments are in lowercase format, e.g., aegis.package.variable
// 2. Environment variables are in uppercase format, e.g., AEGIS_PACKAGE_VARIABLE
func GetOptWithEnv(key string, def ...string) string {
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))
	if ContainsOpt(cmdKey) {
		return GetOpt(cmdKey)
	}

	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	if r, ok := os.LookupEnv(envKey); ok {
		return r
	}

	if len(def) > 0 {
		return def[0]
	}
	return ""
}
