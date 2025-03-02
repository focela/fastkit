// Copyright (c) 2025 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package command provides utilities for console operations,
// such as parsing and retrieving command-line options and arguments.
package command

import (
	"os"
	"regexp"
	"strings"
)

// Default storage for parsed command-line elements.
var (
	// defaultParsedArgs stores positional arguments from command line.
	defaultParsedArgs []string

	// defaultParsedOptions stores option flags with their values.
	defaultParsedOptions map[string]string

	// argumentOptionRegex defines the pattern for recognizing command-line options.
	// It matches both short (-a) and long (--option) formats, with optional values.
	argumentOptionRegex = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=){0,1}(.*)$`)
)

// Init initializes the command-line parser with given arguments.
// If no arguments are provided, it uses os.Args by default.
// It's idempotent if called multiple times without arguments.
func Init(args ...string) {
	// Return early if already initialized and no new args provided
	if len(args) == 0 {
		if len(defaultParsedArgs) > 0 || len(defaultParsedOptions) > 0 {
			return
		}
		args = os.Args
	} else {
		// Reset storage when new args are provided
		defaultParsedArgs = nil
		defaultParsedOptions = make(map[string]string)
	}

	// Parse arguments using the default algorithm
	defaultParsedArgs, defaultParsedOptions = ParseUsingDefaultAlgorithm(args...)
}

// ParseUsingDefaultAlgorithm parses command-line arguments using default algorithm.
// It separates positional arguments from option flags and their values.
// Returns two values:
// - A slice of positional arguments
// - A map of option names to their values
func ParseUsingDefaultAlgorithm(args ...string) ([]string, map[string]string) {
	parsedArgs := make([]string, 0, len(args))
	parsedOptions := make(map[string]string)

	for i := 0; i < len(args); {
		// Check if current argument matches option pattern
		matches := argumentOptionRegex.FindStringSubmatch(args[i])

		// If it's an option
		if len(matches) > 2 {
			optionName := matches[1]

			// Case 1: Option with value after equals sign (--option=value)
			if matches[2] == "=" {
				parsedOptions[optionName] = matches[3]
				i++
				continue
			}

			// Case 2: Option with value as next argument (--option value)
			if i < len(args)-1 && (len(args[i+1]) == 0 || args[i+1][0] != '-') {
				parsedOptions[optionName] = args[i+1]
				i += 2
				continue
			}

			// Case 3: Flag option without value (--flag)
			parsedOptions[optionName] = matches[3]
		} else {
			// It's a positional argument
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}

	return parsedArgs, parsedOptions
}

// GetOpt returns the value of option named 'name'.
// If the option doesn't exist, it returns the first value from 'def' or an empty string.
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

// GetOptAll returns all parsed options as a map.
func GetOptAll() map[string]string {
	Init()
	return defaultParsedOptions
}

// ContainsOpt checks whether an option named 'name' exists in the arguments.
func ContainsOpt(name string) bool {
	Init()
	_, ok := defaultParsedOptions[name]
	return ok
}

// GetArg returns the positional argument at the specified index.
// If the index is out of range, it returns the first value from 'def' or an empty string.
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

// GetArgAll returns all parsed positional arguments as a slice.
func GetArgAll() []string {
	Init()
	return defaultParsedArgs
}

// GetOptWithEnv returns the command line argument of the specified 'key'.
// If the argument does not exist, it returns the environment variable with the specified 'key'.
// It returns the default value 'def' if none of them exists.
//
// Fetching Rules:
// 1. Command line arguments are in lowercase format, e.g., aegis.package.variable
// 2. Environment arguments are in uppercase format, e.g., AEGIS_PACKAGE_VARIABLE
func GetOptWithEnv(key string, def ...string) string {
	// Try command line argument (lowercase with dots)
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))
	if ContainsOpt(cmdKey) {
		return GetOpt(cmdKey)
	}

	// Try environment variable (uppercase with underscores)
	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	if r, ok := os.LookupEnv(envKey); ok {
		return r
	}

	// Return default if provided
	if len(def) > 0 {
		return def[0]
	}

	return ""
}
