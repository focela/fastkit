// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package command provides functionality for parsing and retrieving command-line arguments and options.
package command

import (
	"os"
	"regexp"
	"strings"
)

// Global variables for argument and option parsing.
var (
	// Stores parsed arguments from the command line.
	defaultParsedArgs = make([]string, 0)

	// Stores parsed options from the command line.
	defaultParsedOptions = make(map[string]string)

	// Regex for parsing command-line options.
	argumentOptionRegex = regexp.MustCompile(
		`^\-{1,2}([\w\?\.\-]+)(=){0,1}(.*)$`,
	)
)

// Init initializes the command-line arguments and options.
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

// GetOpt retrieves the value of a command-line option by name.
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

// GetOptAll returns all parsed command-line options.
func GetOptAll() map[string]string {
	Init()
	return defaultParsedOptions
}

// ContainsOpt checks whether an option exists in the parsed arguments.
func ContainsOpt(name string) bool {
	Init()
	_, ok := defaultParsedOptions[name]
	return ok
}

// GetOptWithEnv retrieves an option from command-line arguments or environment variables.
//
// Fetching Rules:
// 1. Command line arguments are in lowercase format, eg: aegis.package.variable.
// 2. Environment arguments are in uppercase format, eg: AEGIS_PACKAGE_VARIABLE.
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

// GetArg retrieves an argument by its index.
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

// GetArgAll returns all parsed command-line arguments.
func GetArgAll() []string {
	Init()
	return defaultParsedArgs
}
