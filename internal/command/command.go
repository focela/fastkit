// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: opensource@focela.com

// Package command provides console operations, like options/arguments reading.
package command

import (
	"os"
	"regexp"
	"strings"
)

var (
	defaultParsedArgs    []string
	defaultParsedOptions map[string]string
	argumentOptionRegex  = regexp.MustCompile(`^-{1,2}([\w?.\-]+)(=)?(.*)$`)
)

func init() {
	defaultParsedArgs = make([]string, 0)
	defaultParsedOptions = make(map[string]string)
}

// Init initializes argument parsing. If no arguments are provided,
// it will use os.Args as the input.
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
	defaultParsedArgs, defaultParsedOptions = ParseUsingDefaultAlgorithm(args...)
}

// ParseUsingDefaultAlgorithm parses command-line arguments using default parsing rules.
func ParseUsingDefaultAlgorithm(args ...string) (parsedArgs []string, parsedOptions map[string]string) {
	parsedArgs = make([]string, 0)
	parsedOptions = make(map[string]string)

	for i := 0; i < len(args); {
		match := argumentOptionRegex.FindStringSubmatch(args[i])
		if len(match) > 2 {
			if match[2] == "=" {
				parsedOptions[match[1]] = match[3]
			} else if i < len(args)-1 && len(args[i+1]) > 0 && args[i+1][0] != '-' {
				parsedOptions[match[1]] = args[i+1]
				i += 2
				continue
			} else {
				parsedOptions[match[1]] = match[3]
			}
		} else {
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}
	return
}

// GetOpt retrieves the value of an option by its name.
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

// ContainsOpt checks whether a specific option exists in the parsed arguments.
func ContainsOpt(name string) bool {
	Init()
	_, ok := defaultParsedOptions[name]
	return ok
}

// GetArg retrieves an argument by its index.
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

// GetArgAll returns all parsed arguments as a slice.
func GetArgAll() []string {
	Init()
	return defaultParsedArgs
}

// GetOptWithEnv retrieves a command-line option or its corresponding environment variable.
// If neither exists, it returns the provided default value.
//
// Fetching Rules:
// 1. Command-line arguments are in lowercase format, e.g., aegis.package.variable.
// 2. Environment variables are in uppercase format, e.g., AEGIS_PACKAGE_VARIABLE.
func GetOptWithEnv(key string, def ...string) string {
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))
	if ContainsOpt(cmdKey) {
		return GetOpt(cmdKey)
	}

	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	if value, ok := os.LookupEnv(envKey); ok {
		return value
	}

	if len(def) > 0 {
		return def[0]
	}
	return ""
}
