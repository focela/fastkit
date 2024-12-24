// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package command provides console operations, like options/arguments reading.
package command

import (
	"os"
	"regexp"
	"strings"
)

// Regular expression for parsing command-line options.
var (
	defaultParsedArgs    = make([]string, 0)
	defaultParsedOptions = make(map[string]string)
	argumentOptionRegex  = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=){0,1}(.*)$`)
)

// Init initializes the default arguments and options.
// It parses the provided arguments or falls back to os.Args if none are given.
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

// ParseUsingDefaultAlgorithm parses command-line arguments into arguments and options.
func ParseUsingDefaultAlgorithm(args ...string) (parsedArgs []string, parsedOptions map[string]string) {
	parsedArgs = make([]string, 0)
	parsedOptions = make(map[string]string)

	for i := 0; i < len(args); {
		array := argumentOptionRegex.FindStringSubmatch(args[i])
		if len(array) > 2 {
			if array[2] == "=" {
				parsedOptions[array[1]] = array[3]
			} else if i < len(args)-1 && len(args[i+1]) > 0 && args[i+1][0] != '-' {
				parsedOptions[array[1]] = args[i+1]
				i += 2
				continue
			} else {
				parsedOptions[array[1]] = array[3]
			}
		} else {
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}
	return
}

// GetOpt retrieves the value of an option by its name.
// If the option does not exist, it returns the provided default value if specified.
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

// ContainsOpt checks if a specific option exists in the parsed options.
func ContainsOpt(name string) bool {
	Init()
	_, ok := defaultParsedOptions[name]
	return ok
}

// GetArg retrieves the argument at a specific index.
// If the index is out of range, it returns the provided default value if specified.
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

// GetOptWithEnv retrieves an option value by key.
// If the option doesn't exist, it tries to retrieve it from the environment variables.
// Command-line keys are in lowercase, and environment variable keys are in uppercase.
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
