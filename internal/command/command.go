// Copyright (c) 2025 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package command provides functionality for handling command-line operations,
// such as parsing and retrieving command-line arguments and options.
// It supports both positional arguments and various option formats.
package command

import (
	"os"
	"regexp"
	"strings"
)

// Parser-related variables
var (
	// defaultParsedArgs stores positional arguments from command line
	defaultParsedArgs []string = make([]string, 0)

	// defaultParsedOptions stores option-value pairs from command line
	defaultParsedOptions map[string]string = make(map[string]string)

	// argumentOptionRegex matches command line options in formats:
	// -name, --name, -name=value, --name=value
	argumentOptionRegex = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=){0,1}(.*)$`)

	// initialized tracks if the default parser has been initialized
	initialized bool
)

// Init initializes the command parser with provided arguments.
// If no arguments are provided and parser is not initialized, it uses os.Args.
func Init(args ...string) {
	if len(args) == 0 {
		if initialized {
			return
		}
		args = os.Args
	} else {
		defaultParsedArgs = make([]string, 0)
		defaultParsedOptions = make(map[string]string)
	}

	// Parse arguments using default algorithm
	defaultParsedArgs, defaultParsedOptions = ParseUsingDefaultAlgorithm(args...)
	initialized = true
}

// ParseUsingDefaultAlgorithm parses command-line arguments using the default algorithm.
// It separates positional arguments from options and their values.
// Returns a slice of positional arguments and a map of option-value pairs.
func ParseUsingDefaultAlgorithm(args ...string) (parsedArgs []string, parsedOptions map[string]string) {
	parsedArgs = make([]string, 0)
	parsedOptions = make(map[string]string)

	for i := 0; i < len(args); {
		array := argumentOptionRegex.FindStringSubmatch(args[i])
		if len(array) > 2 {
			if array[2] == "=" {
				// Handle option with value in format: --name=value
				parsedOptions[array[1]] = array[3]
			} else if i < len(args)-1 {
				if len(args[i+1]) > 0 && args[i+1][0] == '-' {
					// Handle option without value followed by another option: --name --other
					parsedOptions[array[1]] = array[3]
				} else {
					// Handle option with value in format: --name value
					parsedOptions[array[1]] = args[i+1]
					i += 2
					continue
				}
			} else {
				// Handle option without value at the end: --name
				parsedOptions[array[1]] = array[3]
			}
		} else {
			// Handle positional argument
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}

	return
}

// ensureInitialized ensures the parser is initialized before use.
func ensureInitialized() {
	if !initialized {
		Init()
	}
}

// Option retrieval functions

// GetOpt returns the value of the option with the given name.
// If the option doesn't exist, it returns the first default value if provided, or an empty string.
func GetOpt(name string, def ...string) string {
	ensureInitialized()

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
	ensureInitialized()
	return defaultParsedOptions
}

// ContainsOpt checks if an option with the given name exists.
func ContainsOpt(name string) bool {
	ensureInitialized()
	_, ok := defaultParsedOptions[name]
	return ok
}

// Argument retrieval functions

// GetArg returns the positional argument at the given index.
// If the index is out of range, it returns the first default value if provided, or an empty string.
func GetArg(index int, def ...string) string {
	ensureInitialized()

	if index < len(defaultParsedArgs) {
		return defaultParsedArgs[index]
	}

	if len(def) > 0 {
		return def[0]
	}

	return ""
}

// GetArgAll returns all parsed positional arguments.
func GetArgAll() []string {
	ensureInitialized()
	return defaultParsedArgs
}

// Combined option and environment variable handling

// GetOptWithEnv returns the value from command line option or environment variable.
// It first checks for a command line option with the given key (converted to lowercase with dots).
// If not found, it checks for an environment variable (converted to uppercase with underscores).
// If neither exists, it returns the default value if provided, or an empty string.
//
// Fetching Rules:
// 1. Command line arguments are in lowercase format, eg: aegis.package.variable
// 2. Environment arguments are in uppercase format, eg: AEGIS_PACKAGE_VARIABLE
func GetOptWithEnv(key string, def ...string) string {
	ensureInitialized()

	// Try command line option
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))
	if ContainsOpt(cmdKey) {
		return GetOpt(cmdKey)
	}

	// Try environment variable
	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	if r, ok := os.LookupEnv(envKey); ok {
		return r
	}

	// Return default value if provided
	if len(def) > 0 {
		return def[0]
	}

	return ""
}
