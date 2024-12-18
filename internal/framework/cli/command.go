// Copyright (c) 2024 Focela Technologies
// Internal Use Only - Unauthorized use prohibited
// Contact: legal@focela.com

// Package cli provides console operations, like options/arguments reading.
package cli

import (
	"os"
	"regexp"
	"strings"
)

// Global variables for storing parsed arguments and options.
var (
	defaultParsedArgs    = make([]string, 0)
	defaultParsedOptions = make(map[string]string)
	argumentOptionRegex  = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=)?(.*)$`)
)

// Init initializes the CLI parser with the provided arguments.
// If no arguments are provided, it defaults to using `os.Args`.
func Init(args ...string) {
	if len(args) == 0 {
		if len(defaultParsedArgs) == 0 && len(defaultParsedOptions) == 0 {
			args = os.Args
		} else {
			return
		}
	}
	defaultParsedArgs = []string{}
	defaultParsedOptions = map[string]string{}
	defaultParsedArgs, defaultParsedOptions = ParseUsingDefaultAlgorithm(args...)
}

// ParseUsingDefaultAlgorithm parses arguments using a predefined algorithm.
// It separates positional arguments and key-value options.
func ParseUsingDefaultAlgorithm(args ...string) ([]string, map[string]string) {
	parsedArgs := make([]string, 0, len(args))
	parsedOptions := make(map[string]string)

	for i := 0; i < len(args); i++ {
		array := argumentOptionRegex.FindStringSubmatch(args[i])
		if len(array) > 2 {
			if array[2] == "=" {
				parsedOptions[array[1]] = array[3]
			} else if i < len(args)-1 && len(args[i+1]) > 0 && args[i+1][0] != '-' {
				parsedOptions[array[1]] = args[i+1]
				i++ // Skip next argument as it's part of the current option
			} else {
				parsedOptions[array[1]] = array[3]
			}
		} else {
			parsedArgs = append(parsedArgs, args[i])
		}
	}
	return parsedArgs, parsedOptions
}

// GetOpt retrieves the value of a specified option by name.
// Returns the default value if the option is not found.
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

// ContainsOpt checks whether an option exists in the parsed options.
func ContainsOpt(name string) bool {
	Init()
	_, ok := defaultParsedOptions[name]
	return ok
}

// GetArg retrieves the argument at the specified index.
// Returns the default value if the index is out of range.
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

// GetOptWithEnv retrieves the value of a specified key from command-line options or environment variables.
// Returns the default value if neither exists.
func GetOptWithEnv(key string, def ...string) string {
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))
	if ContainsOpt(cmdKey) {
		return GetOpt(cmdKey)
	}
	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	if envVal, ok := os.LookupEnv(envKey); ok {
		return envVal
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}
